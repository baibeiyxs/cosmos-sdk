package legacy

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

// Cdc defines a global generic sealed Amino codec to be used throughout sdk. It
// has all Tendermint crypto and evidence types registered.
//
// TODO: Deprecated - remove this global.
var Cdc *codec.LegacyAmino

func init() {
	Cdc = codec.NewLegacyAmino()
	cryptocodec.RegisterCrypto(Cdc)
	codec.RegisterEvidences(Cdc)
}

// PrivKeyFromBytes unmarshals private key bytes and returns a PrivKey
func PrivKeyFromBytes(privKeyBytes []byte) (privKey cryptotypes.PrivKey, err error) {
	err = Cdc.UnmarshalBinaryBare(privKeyBytes, &privKey)
	return
}

// PubKeyFromBytes unmarshals public key bytes and returns a PubKey
func PubKeyFromBytes(pubKeyBytes []byte) (pubKey cryptotypes.PubKey, err error) {
	err = Cdc.UnmarshalBinaryBare(pubKeyBytes, &pubKey)
	return
}

// AminoPubKeyFromBytes unmarshals public key bytes and returns a PubKey
func AminoPubKeyFromBytes(pubKeyBytes []byte) (pubKey multisig.LegacyAminoPubKey, err error) {
	err = Cdc.UnmarshalBinaryBare(pubKeyBytes, &pubKey)
	return
}
