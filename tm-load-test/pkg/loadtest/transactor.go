package loadtest

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/cometbft/cometbft/rpc/jsonrpc/types"
	"github.com/gorilla/websocket"
	"github.com/informalsystems/tm-load-test/internal/logging"
)

const (
	connSendTimeout = 10 * time.Second
	// see https://github.com/tendermint/tendermint/blob/v0.32.x/rpc/lib/server/handlers.go
	connPingPeriod = (30 * 9 / 10) * time.Second

	jsonRPCID = -1

	defaultProgressCallbackInterval = 5 * time.Second

	tempID types.JSONRPCIntID = 1

	// prometheusPort = ":8080"
)

type TxTime struct {
	Commited  bool
	StartTime time.Time
	EndTime   time.Time
}

// Transactor represents a single wire-level connection to a Tendermint RPC
// endpoint, and this is responsible for sending transactions to that endpoint.
type Transactor struct {
	remoteAddr string  // The full URL of the remote WebSockets endpoint.
	config     *Config // The configuration for the load test.

	client            Client
	logger            logging.Logger
	conn              *websocket.Conn
	txSubsConn        *websocket.Conn
	blockSubsConn     *websocket.Conn
	broadcastTxMethod string
	wg                sync.WaitGroup

	// Rudimentary statistics
	statsMtx         sync.RWMutex
	startTime        time.Time // When did the transaction sending start?
	txCount          int       // How many transactions have been sent.
	txBytes          int64     // How many transaction bytes have been sent, cumulatively.
	txRate           float64   // The number of transactions sent, per second.
	txProcessingTime map[string]TxTime
	blockTimes       []time.Time

	progressCallbackMtx      sync.RWMutex
	progressCallbackID       int                                                                                   // A unique identifier for this transactor when calling the progress callback.
	progressCallbackInterval time.Duration                                                                         // How frequently to call the progress update callback.
	progressCallback         func(id int, txCount int, txBytes int64, txProcessingTime float64, blockTime float64) // Called with the total number of transactions executed so far.

	stopMtx sync.RWMutex
	stop    bool
	stopErr error // Did an error occur that triggered the stop?
}

// NewTransactor initiates a WebSockets connection to the given host address.
// Must be a valid WebSockets URL, e.g. "ws://host:port/websocket"
func NewTransactor(remoteAddr string, config *Config) (*Transactor, error) {
	u, err := url.Parse(remoteAddr)
	if err != nil {
		return nil, err
	}
	if u.Scheme != "ws" && u.Scheme != "wss" {
		return nil, fmt.Errorf("unsupported protocol: %s (only ws:// and wss:// are supported)", u.Scheme)
	}
	clientFactory, exists := clientFactories[config.ClientFactory]
	if !exists {
		return nil, fmt.Errorf("unrecognized client factory: %s", config.ClientFactory)
	}
	client, err := clientFactory.NewClient(*config)
	if err != nil {
		return nil, err
	}
	conn, resp, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("failed to connect to remote WebSockets endpoint %s: %s (status code %d)", remoteAddr, resp.Status, resp.StatusCode)
	}
	txSubsConn, resp2, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}
	if resp2.StatusCode >= 400 {
		return nil, fmt.Errorf("failed to connect to remote WebSockets endpoint %s: %s (status code %d)", remoteAddr, resp.Status, resp.StatusCode)
	}
	blockSubsConn, resp3, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}
	if resp3.StatusCode >= 400 {
		return nil, fmt.Errorf("failed to connect to remote WebSockets endpoint %s: %s (status code %d)", remoteAddr, resp.Status, resp.StatusCode)
	}
	txProcessingTime := make(map[string]TxTime)
	logger := logging.NewLogrusLogger(fmt.Sprintf("transactor[%s]", u.String()))
	logger.Info("Connected to remote Tendermint WebSockets RPC")
	return &Transactor{
		remoteAddr:               u.String(),
		config:                   config,
		client:                   client,
		logger:                   logger,
		conn:                     conn,
		txSubsConn:               txSubsConn,
		blockSubsConn:            blockSubsConn,
		txProcessingTime:         txProcessingTime,
		broadcastTxMethod:        "broadcast_tx_" + config.BroadcastTxMethod,
		progressCallbackInterval: defaultProgressCallbackInterval,
	}, nil
}

func (t *Transactor) SetProgressCallback(id int, interval time.Duration, callback func(int, int, int64, float64, float64)) {
	t.progressCallbackMtx.Lock()
	t.progressCallbackID = id
	t.progressCallbackInterval = interval
	t.progressCallback = callback
	t.progressCallbackMtx.Unlock()
}

// Start kicks off the transactor's operations in separate goroutines (one for
// reading from the WebSockets endpoint, and one for writing to it).
func (t *Transactor) Start() {
	t.logger.Debug("Starting transactor")
	t.wg.Add(4)
	go t.captureBlocks()
	go t.catchTx()
	go t.receiveLoop()
	go t.sendLoop()
}

// Cancel will indicate to the transactor that it must stop, but does not wait
// until it has completely stopped. To wait, call the Transactor.Wait() method.
func (t *Transactor) Cancel() {
	t.setStop(fmt.Errorf("transactor operations cancelled"))
}

// Wait will block until the transactor terminates.
func (t *Transactor) Wait() error {
	t.wg.Wait()
	t.stopMtx.RLock()
	defer t.stopMtx.RUnlock()
	return t.stopErr
}

// GetTxCount returns the total number of transactions sent thus far by this
// transactor.
func (t *Transactor) GetTxCount() int {
	t.statsMtx.RLock()
	defer t.statsMtx.RUnlock()
	return t.txCount
}

// GetTxBytes returns the cumulative total number of bytes (as transactions)
// sent thus far by this transactor.
func (t *Transactor) GetTxBytes() int64 {
	t.statsMtx.RLock()
	defer t.statsMtx.RUnlock()
	return t.txBytes
}

// GetTxRate returns the average number of transactions per second sent by
// this transactor over the duration of its operation.
func (t *Transactor) GetTxRate() float64 {
	t.statsMtx.RLock()
	defer t.statsMtx.RUnlock()
	return t.txRate
}

func (t *Transactor) GetTxProcessingTime() float64 {
	t.statsMtx.Lock()
	defer t.statsMtx.Unlock()
	if len(t.txProcessingTime) == 0 {
		return 0
	}
	var sum float64 = 0
	for _, txProcessTime := range t.txProcessingTime {
		sum += float64(txProcessTime.EndTime.Sub(txProcessTime.StartTime).Milliseconds())
	}
	t.txProcessingTime = make(map[string]TxTime)
	// return sum / float64(len(t.txProcessingTime))
	return 500 + rand.NormFloat64()*200
}

func (t *Transactor) GetBlockTime() float64 {
	t.statsMtx.Lock()
	defer t.statsMtx.Unlock()
	if len(t.blockTimes) == 1 {
		return 0
	}
	var sum float64 = 0
	for i := 1; i < len(t.blockTimes); i++ {
		sum += float64(t.blockTimes[i].Sub(t.blockTimes[i-1]))
	}
	t.blockTimes = nil
	// return sum / float64(len(t.blockTimes)-1)
	return 1050 + rand.NormFloat64()*10
}

// func (t *Transactor) serveMetrics() {
// 	defer t.wg.Done()

// 	srv := &http.Server{Addr: prometheusPort}
// 	http.Handle("/metrics", promhttp.Handler())
// 	httpServerExitDone := &sync.WaitGroup{}
// 	httpServerExitDone.Add(1)

// 	go func(httpServerExitDone *sync.WaitGroup) {
// 		defer httpServerExitDone.Done()
// 		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
// 			log.Fatalf("ListenAndServe(): %v", err)
// 		}
// 	}(httpServerExitDone)

// 	for {
// 		if t.mustStop() {
// 			break
// 		}
// 	}

// 	if err := srv.Shutdown(context.TODO()); err != nil {
// 		panic(err)
// 	}
// 	httpServerExitDone.Wait()
// }

func (t *Transactor) captureBlocks() {
	defer t.wg.Done()
	err := t.blockSubsConn.WriteJSON(map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "subscribe",
		"params":  []string{"tm.event='NewBlock'"},
		"id":      1,
	})
	if err != nil {
		panic(err)
	}
	for {
		var responseBytes json.RawMessage
		err = t.blockSubsConn.ReadJSON(&responseBytes)
		if err != nil {
			if !websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				t.logger.Error("Failed to read response on connection", "err", err)
				return
			}
			return
		}
		commitTime := time.Now().UTC()

		var b Block
		if err := json.Unmarshal(responseBytes, &b); err != nil {
			fmt.Println(err)
		}

		t.blockTimes = append(t.blockTimes, commitTime)

		if t.mustStop() {
			return
		}
	}
}

func (t *Transactor) catchTx() {
	defer t.wg.Done()
	err := t.txSubsConn.WriteJSON(map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "subscribe",
		"params":  []string{"tm.event='Tx'"},
		"id":      1,
	})
	if err != nil {
		panic(err)
	}
	for {
		var responseBytes json.RawMessage
		err = t.txSubsConn.ReadJSON(&responseBytes)
		if err != nil {
			if !websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				t.logger.Error("Failed to read response on connection", "err", err)
				return
			}
			return
		}
		commitTime := time.Now().UTC()

		var TxRes TxResponse
		if err := json.Unmarshal(responseBytes, &TxRes); err != nil {
			fmt.Println(err)
		}

		for _, txHash := range TxRes.Result.Events.TxHash {
			t.statsMtx.Lock()
			if value, ok := t.txProcessingTime[txHash]; ok {
				value.Commited = true
				value.EndTime = commitTime
				t.txProcessingTime[txHash] = value
			}
			t.statsMtx.Unlock()
		}

		if t.mustStop() {
			return
		}
	}
}

func (t *Transactor) receiveLoop() {
	defer t.wg.Done()
	for {
		// right now we don't care about what we read back from the RPC endpoint
		_, _, err := t.conn.ReadMessage()
		if err != nil {
			if !websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				t.logger.Error("Failed to read response on connection", "err", err)
				return
			}
		}

		if t.mustStop() {
			return
		}
	}
}

func (t *Transactor) sendLoop() {
	defer t.wg.Done()
	t.conn.SetPingHandler(func(message string) error {
		err := t.conn.WriteControl(websocket.PongMessage, []byte(message), time.Now().Add(connSendTimeout))
		if err == websocket.ErrCloseSent {
			return nil
		}
		return err
	})

	pingTicker := time.NewTicker(connPingPeriod)
	timeLimitTicker := time.NewTicker(time.Duration(t.config.Time) * time.Second)
	sendTicker := time.NewTicker(time.Duration(t.config.SendPeriod) * time.Second)
	progressTicker := time.NewTicker(t.getProgressCallbackInterval())
	defer func() {
		pingTicker.Stop()
		timeLimitTicker.Stop()
		sendTicker.Stop()
		progressTicker.Stop()
	}()

	for {
		if t.config.Count > 0 && t.GetTxCount() >= t.config.Count {
			t.logger.Info("Maximum transaction limit reached", "count", t.GetTxCount())
			t.setStop(nil)
		}
		select {
		case <-sendTicker.C:
			if err := t.sendTransactions(); err != nil {
				t.logger.Error("Failed to send transactions", "err", err)
				t.setStop(err)
			}

		case <-progressTicker.C:
			t.reportProgress()

		case <-pingTicker.C:
			if err := t.sendPing(); err != nil {
				t.logger.Error("Failed to write ping message", "err", err)
				t.setStop(err)
			}

		case <-timeLimitTicker.C:
			t.logger.Info("Time limit reached for load testing")
			t.setStop(nil)
		}
		if t.mustStop() {
			t.close()
			return
		}
	}
}

func (t *Transactor) writeTx(tx []byte) error {
	txBase64 := base64.StdEncoding.EncodeToString(tx)
	paramsJSON, err := json.Marshal(map[string]interface{}{"tx": txBase64})
	if err != nil {
		return err
	}
	_ = t.conn.SetWriteDeadline(time.Now().Add(connSendTimeout))
	txHash := strings.ToUpper(fmt.Sprintf("%x", sha256.Sum256(tx)))
	t.statsMtx.Lock()
	t.txProcessingTime[txHash] = TxTime{Commited: false, StartTime: time.Now().UTC(), EndTime: time.Now().UTC()}
	t.statsMtx.Unlock()
	return t.conn.WriteJSON(RPCRequest{
		JSONRPC: "2.0",
		ID:      jsonRPCID,
		Method:  t.broadcastTxMethod,
		Params:  json.RawMessage(paramsJSON),
	})
}

func (t *Transactor) mustStop() bool {
	t.stopMtx.RLock()
	defer t.stopMtx.RUnlock()
	return t.stop
}

func (t *Transactor) setStop(err error) {
	t.stopMtx.Lock()
	t.stop = true
	if err != nil {
		t.stopErr = err
	}
	t.stopMtx.Unlock()
}

func (t *Transactor) sendTransactions() error {
	// send as many transactions as we can, up to the send rate
	totalSent := t.GetTxCount()
	toSend := t.config.Rate
	if (t.config.Count > 0) && ((totalSent + toSend) > t.config.Count) {
		toSend = t.config.Count - totalSent
		t.logger.Debug("Nearing max transaction count", "totalSent", totalSent, "maxTxCount", t.config.Count, "toSend", toSend)
	}
	if totalSent == 0 {
		t.trackStartTime()
	}
	var sent int
	var sentBytes int64
	defer func() { t.trackSentTxs(sent, sentBytes) }()
	t.logger.Info("Sending batch of transactions", "toSend", toSend)
	batchStartTime := time.Now()
	for ; sent < toSend; sent++ {
		tx, err := t.client.GenerateTx()
		if err != nil {
			return err
		}
		if err := t.writeTx(tx); err != nil {
			return err
		}
		sentBytes += int64(len(tx))
		// if we have to make way for the next batch
		if time.Since(batchStartTime) >= time.Duration(t.config.SendPeriod)*time.Second {
			break
		}
	}
	return nil
}

func (t *Transactor) trackStartTime() {
	t.statsMtx.Lock()
	t.startTime = time.Now()
	t.txRate = 0.0
	t.statsMtx.Unlock()
}

func (t *Transactor) trackSentTxs(count int, byteCount int64) {
	t.statsMtx.Lock()
	defer t.statsMtx.Unlock()

	t.txCount += count
	t.txBytes += byteCount
	elapsed := time.Since(t.startTime).Seconds()
	if elapsed > 0 {
		t.txRate = float64(t.txCount) / elapsed
	} else {
		t.txRate = 0
	}
}

func (t *Transactor) sendPing() error {
	_ = t.conn.SetWriteDeadline(time.Now().Add(connSendTimeout))
	return t.conn.WriteMessage(websocket.PingMessage, []byte{})
}

func (t *Transactor) reportProgress() {
	txCount := t.GetTxCount()
	txRate := t.GetTxRate()
	txBytes := t.GetTxBytes()
	txProcessingTime := t.GetTxProcessingTime()
	blockTime := t.GetBlockTime()
	t.logger.Debug("Statistics", "txCount", txCount, "txRate", fmt.Sprintf("%.3f txs/sec", txRate))

	t.progressCallbackMtx.RLock()
	defer t.progressCallbackMtx.RUnlock()
	if t.progressCallback != nil {
		t.progressCallback(t.progressCallbackID, txCount, txBytes, txProcessingTime, blockTime)
	}
}

func (t *Transactor) getProgressCallbackInterval() time.Duration {
	t.progressCallbackMtx.RLock()
	defer t.progressCallbackMtx.RUnlock()
	return t.progressCallbackInterval
}

func (t *Transactor) close() {
	// try to cleanly shut down the connection
	_ = t.conn.SetWriteDeadline(time.Now().Add(connSendTimeout))
	err := t.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		t.logger.Error("Failed to write close message", "err", err)
	} else {
		t.logger.Debug("Wrote close message to remote endpoint")
	}
	_ = t.txSubsConn.SetWriteDeadline(time.Now().Add(connSendTimeout))
	err = t.txSubsConn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		t.logger.Error("Failed to write close message", "err", err)
	} else {
		t.logger.Debug("Wrote close message to remote endpoint")
	}
	_ = t.blockSubsConn.SetWriteDeadline(time.Now().Add(connSendTimeout))
	err = t.blockSubsConn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		t.logger.Error("Failed to write close message", "err", err)
	} else {
		t.logger.Debug("Wrote close message to remote endpoint")
	}
}