package noble

import (
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/informalsystems/tm-load-test/pkg/loadtest"
	"github.com/wfblockchain/noblechain/v5/app"
	"github.com/wfblockchain/noblechain/v5/cmd"
	tf "github.com/wfblockchain/noblechain/v5/x/tokenfactory/types"

	"github.com/gogo/protobuf/proto"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
)

const (
	KVStoreClientIDLen int = 5 // Allows for 6,471,002 random client IDs (62C5)
	kvstoreMinValueLen int = 1 // We at least need 1 character in a key/value pair's value.
	chainID                = "noble-1"
	pvtKey                 = "1e888b53ef1278956c590255b77259b4a2572fa1938bc61304011b3aa549cada" // pvtkey of owner
	ownerAddr              = "noble1hsjfews729je9h0v5tdd94xcgqr4phkuptakxp"                     // owner-address
	masterMinterAddr       = "noble1a3j0z8pq960apqvhwp4nu5gv80mk4hvhwg3mu0"                     // masterMinter-address
	ownerPubKey            = "Avor+7fYYiDhfHhmM9RQSEFHK+IDsw32G+K/PD+Z3F1f"                     // owner-address
	masterMinterPubKey     = "AmVDyRnk7I/VYh5jpt1EJTgcbSAsZxO5UdZrIK0iuBvY"                     // masterMinter-address
	masterMinterPvtKey     = "7c27b0844932f078c941960c70bf4d0bec51f77cf6349b84617f19350a09a8d0" // masterMinter-address
)

type NobleClientFactory struct{}

type NobleClient struct{}

var (
	_ loadtest.ClientFactory = (*NobleClientFactory)(nil)
	_ loadtest.Client        = (*NobleClient)(nil)
)

func init() {
	// if err := loadtest.RegisterClientFactory("noble", NewNobleClientFactory()); err != nil {
	// 	fmt.Println("yoooo")
	// 	panic(err)
	// }
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
	accNum, accSeq := uint64(0), uint64(0)

	msg, err := createMsgs()
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

func createMsgs() (*tf.MsgUpdateMasterMinter, error) {
	// amt, ok := math.NewIntFromString("1")
	// if !ok {
	// 	panic("amnt wrong")
	// }
	// mmPub, _ := hex.DecodeString(masterMinterPubKey)
	// masterMinterPub := secp256k1.PubKey{Key: mmPub}
	// masterMinter := sdk.AccAddress(masterMinterPub.Address())
	// fmt.Printf(masterMinter.String())
	// // if err != nil {
	// // 	fmt.Printf("err: %s\n", err)
	// // 	panic("master address not correct")
	// // }
	// oPub, _ := hex.DecodeString(ownerPubKey)
	// ownerPub := secp256k1.PubKey{Key: oPub}
	// owner := sdk.AccAddress(ownerPub.Address())
	// // owner, err := sdk.AccAddressFromBech32(ownerAddr)
	// // if err != nil {
	// // 	panic("owner address not correct")
	// // }
	// msg := bank.NewMsgSend(owner, masterMinter, sdk.NewCoins(sdk.NewCoin("ustake", amt)))
	privKey := secp256k1.GenPrivKey()
	err := proto.Unmarshal([]byte(masterMinterPvtKey), privKey)
	pbkk := privKey.PubKey()
	fmt.Println(pbkk)

	if err != nil {
		fmt.Println("Error while unmarshaling privKey: ", err)
		panic(pbkk)
		// return nil, err
	}
	msg := tf.NewMsgUpdateMasterMinter(ownerAddr, masterMinterAddr)
	if err := msg.ValidateBasic(); err != nil {
		panic(err)
	}

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
