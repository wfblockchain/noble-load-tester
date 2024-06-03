package noble

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"

	// "github.com/CosmWasm/wasmd/x/wasm/ioutils"
	// "github.com/CosmWasm/wasmd/x/wasm/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/informalsystems/tm-load-test/pkg/loadtest"

	// "go.starlark.net/lib/proto"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/gogo/protobuf/proto"
	"google.golang.org/grpc"

	// "github.com/wfblockchain/distributed_finance/chains/difi/app"

	"github.com/wfblockchain/noblechain/v5/app"
	"github.com/wfblockchain/noblechain/v5/cmd"
	tftypes "github.com/wfblockchain/noblechain/v5/x/tokenfactory/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
)

const (
	chainID = "noble-1"

	mintAmt          = 1
	burnAmt          = 1
	transferAmt      = 10
	transferCntLimit = 100
	tfDenom          = "utoken"
	wasmBinaryPath   = "/home/leo10/distributed_finance/cw-contracts/target/wasm32-unknown-unknown/release"
)

type NobleClientFactory struct{}

type NobleClient struct {
	// minterAddr    string
	// aliceAddr     string
	// minterPrivKey string
	// alicePrivKey  string
	// minterAccSeq  uint64
	// minterAccNum  uint64
	// aliceAccSeq   uint64
	// aliceAccNum   uint64
}

var (
	// _            wasmtypes.MsgStoreCode
	_             loadtest.ClientFactory = (*NobleClientFactory)(nil)
	_             loadtest.Client        = (*NobleClient)(nil)
	minterAccSeq  uint64                 = 1715
	minterAccNum  uint64                 = 8
	aliceAccSeq   uint64                 = 0
	aliceAccNum   uint64                 = 11
	store         bool                   = false
	mint          bool                   = true
	transfer      bool                   = false
	burn          bool                   = false
	transferCnt   int                    = 0
	minterAddr    string                 = "noble1ea9xey2ujyrm8xzadvrykkfg59tqyjr4md2p8y"
	aliceAddr     string                 = "noble16aq2nvjac83x4yykmpdd86p099x40shd7sad2w"
	minterPrivKey string                 = "929dc0dbdba0f90c837fdf4db88fdd5cd71ffed09f3cea66bcb680a13b2dc58c"
	alicePrivKey  string                 = "d788bfafb3464cc006bd1e2f29832d85bebfb291342143a6f1e56635a69034fe"
)

func init() {
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("noble", "noblepub")
	minterAddr = os.Getenv("MINTER_ADDR")
	aliceAddr = os.Getenv("ALICE_ADDR")
	minterPrivKey = os.Getenv("MINTER_PRIV")
	alicePrivKey = os.Getenv("ALICE_PRIV")
	minterAccNumInt, err := strconv.Atoi(os.Getenv("MINTER_ACC_NUM"))
	if err != nil {
		panic(err)
	}
	minterAccNum = uint64(minterAccNumInt)
	aliceAccNumInt, err := strconv.Atoi(os.Getenv("ALICE_ACC_NUM"))
	if err != nil {
		panic(err)
	}
	aliceAccNum = uint64(aliceAccNumInt)
	minterAccSeqInt, err := strconv.Atoi(os.Getenv("MINTER_ACC_SEQ"))
	if err != nil {
		panic(err)
	}
	minterAccSeq = uint64(minterAccSeqInt)
	aliceAccSeqInt, err := strconv.Atoi(os.Getenv("ALICE_ACC_SEQ"))
	if err != nil {
		panic(err)
	}
	aliceAccSeq = uint64(aliceAccSeqInt)
	// minterAccNum, minterAccSeq = getUserInfo(minterAddr)
	// aliceAccNum, aliceAccSeq = getUserInfo(aliceAddr)
}

func NewNobleClientFactory() *NobleClientFactory {
	return &NobleClientFactory{}
}

func (f *NobleClientFactory) ValidateConfig(cfg loadtest.Config) error {
	return nil
}

func (f *NobleClientFactory) NewClient(cfg loadtest.Config) (loadtest.Client, error) {
	// minterAddr := os.Getenv("MINTER_ADDR")
	// aliceAddr := os.Getenv("ALICE_ADDR")
	// minterPrivKey := os.Getenv("MINTER_PRIV")
	// alicePrivKey := os.Getenv("ALICE_PRIV")
	// minterAccNum, minterAccSeq := getUserInfo(minterAddr)
	// aliceAccNum, aliceAccSeq := getUserInfo(aliceAddr)
	// minterAccNumInt, err := strconv.Atoi(os.Getenv("MINTER_ACC_NUM"))
	// if err != nil {
	// 	panic(err)
	// }
	// minterAccNum := uint64(minterAccNumInt)
	// aliceAccNumInt, err := strconv.Atoi(os.Getenv("ALICE_ACC_NUM"))
	// if err != nil {
	// 	panic(err)
	// }
	// aliceAccNum := uint64(aliceAccNumInt)
	// minterAccSeqInt, err := strconv.Atoi(os.Getenv("MINTER_ACC_SEQ"))
	// if err != nil {
	// 	panic(err)
	// }
	// minterAccSeq := uint64(minterAccSeqInt)
	// aliceAccSeqInt, err := strconv.Atoi(os.Getenv("ALICE_ACC_SEQ"))
	// if err != nil {
	// 	panic(err)
	// }
	// aliceAccSeq := uint64(aliceAccSeqInt)
	// minterAccNum, minterAccSeq = getUserInfo(minterAddr)
	// aliceAccNum, aliceAccSeq = getUserInfo(aliceAddr)
	// return &NobleClient{
	// 	minterAddr:    minterAddr,
	// 	aliceAddr:     aliceAddr,
	// 	minterPrivKey: minterPrivKey,
	// 	alicePrivKey:  alicePrivKey,
	// 	minterAccNum:  minterAccNum,
	// 	minterAccSeq:  minterAccSeq,
	// 	aliceAccNum:   aliceAccNum,
	// 	aliceAccSeq:   aliceAccSeq,
	// }, nil
	return &NobleClient{}, nil
}

func (c *NobleClient) GenerateTx() ([]byte, error) {
	TxConfig := cmd.MakeEncodingConfig(app.ModuleBasics).TxConfig
	TxBuilder := TxConfig.NewTxBuilder()
	TxBuilder.SetGasLimit(500000)

	err := c.createMsgsAndSign(TxBuilder, TxConfig)
	if err != nil {
		fmt.Println(err)
		panic("msg creation failed")
	}

	txBytes, err := TxConfig.TxEncoder()(TxBuilder.GetTx())
	if err != nil {
		fmt.Println(err)
		panic("encoder failed")
	}

	return txBytes, nil
}

func (c *NobleClient) createMsgsAndSign(TxBuilder client.TxBuilder, TxConfig client.TxConfig) error {
	var msg sdk.Msg
	switch {
	case mint:
		msg = tftypes.NewMsgMint(minterAddr, aliceAddr, sdk.NewInt64Coin(tfDenom, mintAmt))
		err := TxBuilder.SetMsgs(msg)
		if err != nil {
			panic(err)
		}
		// minterAccNum, minterAccSeq := getUserInfo(minterAddr)
		err = signTX(TxBuilder, TxConfig, minterAccNum, minterAccSeq, minterPrivKey)
		if err != nil {
			panic(err)
		}
		minterAccSeq++
		// mint = false
		// transfer = true
		fmt.Println(minterAccSeq)
	case transfer:
		alice, err := sdk.AccAddressFromBech32(aliceAddr)
		if err != nil {
			panic(err)
		}
		mint, err := sdk.AccAddressFromBech32(minterAddr)
		if err != nil {
			panic(err)
		}
		msg = banktypes.NewMsgSend(alice, mint, sdk.NewCoins(sdk.NewInt64Coin(tfDenom, transferAmt)))
		err = TxBuilder.SetMsgs(msg)
		if err != nil {
			panic(err)
		}
		// aliceAccNum, aliceAccSeq := getUserInfo(aliceAddr)
		err = signTX(TxBuilder, TxConfig, aliceAccNum, aliceAccSeq, alicePrivKey)
		if err != nil {
			panic(err)
		}
		aliceAccSeq++
		transferCnt++
		if transferCnt > transferCntLimit {
			transfer = false
			burn = true
		}
	case burn:
		msg = tftypes.NewMsgBurn(minterAddr, sdk.NewInt64Coin(tfDenom, burnAmt))
		err := TxBuilder.SetMsgs(msg)
		if err != nil {
			panic(err)
		}
		// minterAccNum, minterAccSeq := getUserInfo(minterAddr)
		err = signTX(TxBuilder, TxConfig, minterAccNum, minterAccSeq, minterPrivKey)
		minterAccSeq++
		if err != nil {
			panic(err)
		}
		// case store:
		// 	sender := aliceAddr
		// 	msg, err := prepareStoreCodeMsg(wasmBinaryPath, sender)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	err = TxBuilder.SetMsgs(&msg)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	err = signTX(TxBuilder, TxConfig, aliceAccNum, aliceAccSeq, alicePrivKey)
		// 	aliceAccSeq++
		// 	if err != nil {
		// 		panic(err)
		// 	}
	}

	return nil
}

// func prepareStoreCodeMsg(file string, sender string) (wasmtypes.MsgStoreCode, error) {
// 	wasm, err := os.ReadFile(file)
// 	if err != nil {
// 		return wasmtypes.MsgStoreCode{}, err
// 	}

// 	// gzip the wasm file
// 	if ioutils.IsWasm(wasm) {
// 		wasm, err = ioutils.GzipIt(wasm)

// 		if err != nil {
// 			return wasmtypes.MsgStoreCode{}, err
// 		}
// 	} else if !ioutils.IsGzip(wasm) {
// 		return wasmtypes.MsgStoreCode{}, fmt.Errorf("invalid input file. Use wasm binary or gzip")
// 	}

// 	msg := types.MsgStoreCode{
// 		Sender:                sender,
// 		WASMByteCode:          wasm,
// 		InstantiatePermission: nil,
// 	}
// 	return msg, msg.ValidateBasic()
// }

func signTX(TxBuilder client.TxBuilder, TxConfig client.TxConfig, accNum uint64, accSeq uint64, privateKey string) error {
	privB, _ := hex.DecodeString(privateKey)
	priv1 := secp256k1.PrivKey{Key: privB}
	/*************************************************************************************************/
	privs := []cryptotypes.PrivKey{&priv1}
	accNums := []uint64{accNum} // The accounts' account numbers
	accSeqs := []uint64{accSeq} // The accounts' sequence numbers

	// First round: we gather all the signer infos. We use the "set empty
	// signature" hack to do that.
	var sigsV2 []signing.SignatureV2
	for i, priv := range privs {
		sigV2 := signing.SignatureV2{
			PubKey: priv.PubKey(),
			Data: &signing.SingleSignatureData{
				SignMode:  TxConfig.SignModeHandler().DefaultMode(),
				Signature: nil,
			},
			Sequence: accSeqs[i],
		}

		sigsV2 = append(sigsV2, sigV2)
	}
	err := TxBuilder.SetSignatures(sigsV2...)
	if err != nil {
		return err
	}

	// Second round: all signer infos are set, so each signer can sign.
	sigsV2 = []signing.SignatureV2{}
	for i, priv := range privs {
		signerData := xauthsigning.SignerData{
			ChainID:       chainID,
			AccountNumber: accNums[i],
			Sequence:      accSeqs[i],
		}
		sigV2, err := tx.SignWithPrivKey(
			TxConfig.SignModeHandler().DefaultMode(),
			signerData,
			TxBuilder,
			priv,
			TxConfig,
			accSeqs[i],
		)
		if err != nil {
			return err
		}

		sigsV2 = append(sigsV2, sigV2)
	}
	err = TxBuilder.SetSignatures(sigsV2...)
	if err != nil {
		return err
	}
	return nil
}

func getUserInfo(addressStr string) (uint64, uint64) {
	conn, err2 := grpc.Dial(
		"172.17.0.1:9090",   // Or your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)
	if err2 != nil {
		fmt.Println(err2)
	}
	defer conn.Close()
	// keyBase := keys.ExportKeyCommand()
	clientCtx := client.Context{}
	// clientCtx.Offline = true
	clientCtx.WithChainID(chainID)                   // set the chain ID
	clientCtx.WithNodeURI("http://172.17.0.1:26657") // set the node URL
	queryClient := authTypes.NewQueryClient(conn)
	accountAddr, err1 := sdk.AccAddressFromBech32(addressStr)
	if err1 != nil {
		panic(err1)
	}
	accountResp, err := queryClient.Account(context.Background(), &authTypes.QueryAccountRequest{Address: accountAddr.String()})
	if err != nil {
		panic(err)
	}
	account := accountResp.GetAccount()

	var acc authTypes.BaseAccount
	err = proto.Unmarshal(account.Value, &acc)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(acc)
	return acc.AccountNumber, acc.Sequence
}
