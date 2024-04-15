package web3

import (
	"encoding/hex"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
)

type Wallet struct {
	privateKey string
	publicKey  string
	address    string
}

func CreateWalletP2TR(net *chaincfg.Params) *Wallet {
	key, _ := btcec.NewPrivateKey()
	pk := key.PubKey().SerializeCompressed()
	wif, _ := btcutil.NewWIF(key, &chaincfg.TestNet3Params, true)
	privateKey := wif.String()
	address, _ := btcutil.NewAddressPubKeyHash(btcutil.Hash160(pk), net)
	return &Wallet{privateKey: privateKey, address: address.EncodeAddress(), publicKey: hex.EncodeToString(pk)}
}

func ImportPrivateKeyHexToLegacy(hex []byte, net *chaincfg.Params) string {
	key, _ := btcec.PrivKeyFromBytes(hex)
	pk := key.PubKey().SerializeCompressed()
	address, _ := btcutil.NewAddressPubKeyHash(btcutil.Hash160(pk), net)
	return address.EncodeAddress()
}

func ImportPrivateKeyHexToP2TR(hex []byte, net *chaincfg.Params) string {
	key, _ := btcec.PrivKeyFromBytes(hex)
	tapKey := txscript.ComputeTaprootKeyNoScript(key.PubKey())
	p2tr, _ := btcutil.NewAddressTaproot(schnorr.SerializePubKey(tapKey), net)
	return p2tr.EncodeAddress()
}

func ImportPrivateKeyToLegacy(privateKey string, net *chaincfg.Params) string {
	wif, err := btcutil.DecodeWIF(privateKey)
	if err != nil {
		panic("privateKey error")
	}
	privKey := wif.PrivKey.Serialize()
	return ImportPrivateKeyHexToLegacy(privKey, net)
}

func ImportPrivateKeyToP2TR(privateKey string, net *chaincfg.Params) string {
	wif, err := btcutil.DecodeWIF(privateKey)
	if err != nil {
		panic("privateKey error")
	}
	privKey := wif.PrivKey.Serialize()
	return ImportPrivateKeyHexToP2TR(privKey, net)
}

func GetAddressTransactions(addr string, url string) {

}
