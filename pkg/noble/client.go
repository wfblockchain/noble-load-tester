package noble

import (
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/informalsystems/tm-load-test/pkg/loadtest"
	"github.com/wfblockchain/noblechain/v5/app"
	"github.com/wfblockchain/noblechain/v5/cmd"
	tftypes "github.com/wfblockchain/noblechain/v5/x/tokenfactory/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
)

const (
	chainID             = "noble-1"
	pvtKey              = "3e3d956c26ef52304a07d91950fc4a6a8da6153ba2d589abf2f7cb9afedbe33b" // pvtkey of owner
	ownerAddr           = "noble16kdnj7qvpysku6lzsns6zqxckldtd73thw7xc8"                     // owner-address
	masterMinterAddr    = "noble167rtuqztu3hu0rgg0sflhvqa9k342n4p067j6f"                     // masterMinter-address
	ownerPrivKey        = pvtKey                                                             // owner-privkey
	masterMinterPrivKey = "4c0f886131459c4890e734f75f96f50fef872059ec77493e8cb0631e2bfad480" // masterMinter-privkey
)

type NobleClientFactory struct{}

type NobleClient struct{}

var (
	_      loadtest.ClientFactory = (*NobleClientFactory)(nil)
	_      loadtest.Client        = (*NobleClient)(nil)
	accSeq uint64                 = uint64(181)
	accNum uint64                 = uint64(0)
)

func init() {
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("noble", "noblepub")
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
	TxBuilder.SetGasLimit(500000000)

	// accNum, accSeq := getUserInfo(grpcConn)

	msg, err := createMsgs1()
	if err != nil {
		fmt.Println(err)
		panic("msg creation failed")
	}

	err1 := TxBuilder.SetMsgs(msg)
	if err1 != nil {
		fmt.Println(err1)
		panic("msg setting failed")
	}

	//signing the msg
	err = signTX(TxBuilder, TxConfig, accNum, accSeq)
	accSeq++
	if err != nil {
		fmt.Println(err)
		panic("sign failed")
	}
	txBytes, err := TxConfig.TxEncoder()(TxBuilder.GetTx())
	if err != nil {
		fmt.Println(err)
		panic("encoder failed")
	}

	return txBytes, nil
}

func createMsgs() (*banktypes.MsgSend, error) {
	owner, err := sdk.AccAddressFromBech32(ownerAddr)
	if err != nil {
		panic(err)
	}
	masterMinter, err := sdk.AccAddressFromBech32(masterMinterAddr)
	if err != nil {
		panic(err)
	}

	msg := banktypes.NewMsgSend(owner, masterMinter, sdk.NewCoins(sdk.NewInt64Coin("stake", 10000)))

	return msg, nil
}

func createMsgs1() (*tftypes.MsgUpdateMasterMinter, error) {
	// owner, err := sdk.AccAddressFromBech32(ownerAddr)
	// if err != nil {
	// 	panic(err)
	// }
	// masterMinter, err := sdk.AccAddressFromBech32(masterMinterAddr)
	// if err != nil {
	// 	panic(err)
	// }

	msg := tftypes.NewMsgUpdateMasterMinter(ownerAddr, masterMinterAddr)

	return msg, nil
}

func signTX(TxBuilder client.TxBuilder, TxConfig client.TxConfig, accNum uint64, accSeq uint64) error {
	privB, _ := hex.DecodeString(pvtKey)
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
