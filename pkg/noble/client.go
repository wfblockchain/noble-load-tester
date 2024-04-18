package noble

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/gogo/protobuf/proto"
	"github.com/informalsystems/tm-load-test/pkg/loadtest"
	"github.com/wfblockchain/noblechain/v5/app"
	"github.com/wfblockchain/noblechain/v5/cmd"
	tftypes "github.com/wfblockchain/noblechain/v5/x/tokenfactory/types"
	"google.golang.org/grpc"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
)

const (
	chainID          = "noble-1"
	minterAddr       = "noble1trtrz3frm525u3f3p7ljg95eck46nnamucjtkn" // minter-address
	aliceAddr        = "noble1tweszwejavrg03rzr0rz58zud9g32jst979tud"
	minterPrivKey    = "d76e7da1dac0301e0491c560bd5923b94df2f25f06fc804ebad87e7478cf52b2" // minter-privkey
	alicePrivKey     = "cd5d41f79c7ef00e44f747a7b404cbaa9172416ded9fe98760b0ebbf82dddaed" // alice-privkey
	mintAmt          = 100000000
	burnAmt          = 1
	transferAmt      = 10
	transferCntLimit = 100
	tfDenom          = "utoken"
)

type NobleClientFactory struct{}

type NobleClient struct{}

var (
	_            loadtest.ClientFactory = (*NobleClientFactory)(nil)
	_            loadtest.Client        = (*NobleClient)(nil)
	minterAccSeq uint64                 = uint64(722)
	minterAccNum uint64                 = uint64(8)
	aliceAccSeq  uint64                 = uint64(59)
	aliceAccNum  uint64                 = uint64(12)
	mint         bool                   = true
	transfer     bool                   = false
	burn         bool                   = false
	transferCnt  int                    = 0
)

func init() {
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("noble", "noblepub")
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
	return &NobleClient{}, nil
}

func (c *NobleClient) GenerateTx() ([]byte, error) {
	TxConfig := cmd.MakeEncodingConfig(app.ModuleBasics).TxConfig
	TxBuilder := TxConfig.NewTxBuilder()
	TxBuilder.SetGasLimit(500000)

	// accNum, accSeq := getUserInfo(grpcConn)

	err := createMsgsAndSign(TxBuilder, TxConfig)
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

func createMsgsAndSign(TxBuilder client.TxBuilder, TxConfig client.TxConfig) error {
	var msg sdk.Msg
	switch {
	case mint:
		msg = tftypes.NewMsgMint(minterAddr, aliceAddr, sdk.NewInt64Coin(tfDenom, mintAmt))
		err := TxBuilder.SetMsgs(msg)
		if err != nil {
			panic(err)
		}
		minterAccNum, minterAccSeq := getUserInfo(minterAddr)
		err = signTX(TxBuilder, TxConfig, minterAccNum, minterAccSeq, minterPrivKey)
		if err != nil {
			panic(err)
		}
		// minterAccSeq++
		mint = false
		transfer = true
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
		aliceAccNum, aliceAccSeq := getUserInfo(aliceAddr)
		err = signTX(TxBuilder, TxConfig, aliceAccNum, aliceAccSeq, alicePrivKey)
		if err != nil {
			panic(err)
		}
		// aliceAccSeq++
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
		minterAccNum, minterAccSeq := getUserInfo(minterAddr)
		err = signTX(TxBuilder, TxConfig, minterAccNum, minterAccSeq, minterPrivKey)
		// minterAccSeq++
		if err != nil {
			panic(err)
		}
	}
	err := TxBuilder.SetMsgs(msg)
	if err != nil {
		fmt.Println(err)
		panic("msg setting failed")
	}
	return nil
}

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
		"127.0.0.1:9090",    // Or your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)
	if err2 != nil {
		fmt.Println(err2)
	}
	defer conn.Close()
	// keyBase := keys.ExportKeyCommand()
	clientCtx := client.Context{}
	// clientCtx.Offline = true
	clientCtx.WithChainID(chainID)                  // set the chain ID
	clientCtx.WithNodeURI("http://localhost:26657") // set the node URL
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
