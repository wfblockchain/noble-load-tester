package noble

import (
	"fmt"

	"github.com/noble-assets/noble/v5/app"

	"github.com/informalsystems/tm-load-test/pkg/loadtest"

	"encoding/hex"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
)

const (
	KVStoreClientIDLen int = 5 // Allows for 6,471,002 random client IDs (62C5)
	kvstoreMinValueLen int = 1 // We at least need 1 character in a key/value pair's value.
	chainID                = "noble-1"
	pvtKey                 = "key"
)

type NobleClientFactory struct{}

// KVStoreClient generates arbitrary transactions (random key=value pairs) to
// be sent to the kvstore ABCI application. The keys are structured as follows:
//
// `[client_id][tx_id]=[tx_id]`
//
// where each value (`client_id` and `tx_id`) is padded with 0s to meet the
// transaction size requirement.
type NobleClient struct {
}

var (
	_ loadtest.ClientFactory = (*NobleClientFactory)(nil)
	_ loadtest.Client        = (*NobleClient)(nil)
)

func init() {
	if err := loadtest.RegisterClientFactory("noble", NewNobleClientFactory()); err != nil {
		panic(err)
	}
}

func NewNobleClientFactory() *NobleClientFactory {
	return &NobleClientFactory{}
}

func (f *NobleClientFactory) ValidateConfig(cfg loadtest.Config) error {
	// maxTxsPerEndpoint := cfg.MaxTxsPerEndpoint()
	// if maxTxsPerEndpoint < 1 {
	// 	return fmt.Errorf("cannot calculate an appropriate maximum number of transactions per endpoint (got %d)", maxTxsPerEndpoint)
	// }
	// minKeySuffixLen, err := requiredKVStoreSuffixLen(maxTxsPerEndpoint)
	// if err != nil {
	// 	return err
	// }
	// // "[client_id][random_suffix]=[value]"
	// minTxSize := KVStoreClientIDLen + minKeySuffixLen + 1 + kvstoreMinValueLen
	// if cfg.Size < minTxSize {
	// 	return fmt.Errorf("transaction size %d is too small for given parameters (should be at least %d bytes)", cfg.Size, minTxSize)
	// }
	return nil
}

func (f *NobleClientFactory) NewClient(cfg loadtest.Config) (loadtest.Client, error) {
	// keyPrefix := []byte(randStr(KVStoreClientIDLen))
	// keySuffixLen, err := requiredKVStoreSuffixLen(cfg.MaxTxsPerEndpoint())
	// if err != nil {
	// 	return nil, err
	// }
	// keyLen := len(keyPrefix) + keySuffixLen
	// // value length = key length - 1 (to cater for "=" symbol)
	// valueLen := cfg.Size - keyLen - 1
	return &NobleClient{}, nil
}

// func requiredKVStoreSuffixLen(maxTxCount uint64) (int, error) {
// 	// for l, maxTxs := range kvstoreMaxTxsByKeySuffixLen {
// 	// 	if maxTxCount < maxTxs {
// 	// 		if l+1 > len(kvstoreMaxTxsByKeySuffixLen) {
// 	// 			return -1, fmt.Errorf("cannot cater for maximum tx count of %d (too many unique transactions, suffix length %d)", maxTxCount, l+1)
// 	// 		}
// 	// 		// we use l+1 to minimize collision probability
// 	// 		return l + 1, nil
// 	// 	}
// 	// }
// 	// return -1, fmt.Errorf("cannot cater for maximum tx count of %d (too many unique transactions)", maxTxCount)
// }

func (c *NobleClient) GenerateTx() ([]byte, error) {
	TxConfig := app.MakeEncodingConfig().TxConfig

	TxBuilder := TxConfig.NewTxBuilder()
	TxBuilder.SetGasLimit(500000)

	// accNum, accSeq := getUserInfo(grpcConn)
	accNum, accSeq := uint64(0), uint64(0)

	msg, err := createMsgs()
	if err != nil {
		fmt.Println(err)
	}

	err1 := TxBuilder.SetMsgs(msg)
	if err1 != nil {
		fmt.Println(err1)
	}

	//signing the msg
	error := signTX(TxBuilder, TxConfig, accNum, accSeq)
	if error != nil {
		fmt.Println(error)
	}
	return []byte{}, nil
}

func createMsgs() (error, error) {

	// did := types.NewChainDID(chainID, id)

	// vmID := did.NewVerificationMethodID(signer)

	// auth := types.NewVerification(
	// 	types.NewVerificationMethod(
	// 		vmID,
	// 		did,
	// 		types.NewPublicKeyMultibase([]byte(pubKey), vmType),
	// 	),
	// 	[]string{types.Authentication},
	// 	nil,
	// )

	// msg := types.NewMsgCreateDidDocument(
	// 	did.String(),
	// 	types.Verifications{auth},
	// 	types.Services{},
	// 	signer,
	// )
	return nil, nil
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
