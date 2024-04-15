package web3

import (
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/gophero/goal/httpx"
	"io"
	"net/http"
	"testing"
)

func TestImportWallet(t *testing.T) {
	legacy := ImportPrivateKeyToLegacy("cRUYvPYvF8RJiPKxx5tahcBBhGJ1eqxSUN5UKizwmVqSRcchQUv4", &chaincfg.TestNet3Params)
	p2tr := ImportPrivateKeyToP2TR("cRUYvPYvF8RJiPKxx5tahcBBhGJ1eqxSUN5UKizwmVqSRcchQUv4", &chaincfg.TestNet3Params)
	fmt.Printf("legacy addr: %s \n", legacy)
	fmt.Printf("P2TR addr: %s \n", p2tr)
}

func TestCreateWalletP2TR(t *testing.T) {
	legacyWallet := CreateWalletP2TR(&chaincfg.TestNet3Params)
	fmt.Println(legacyWallet)
	p2tr := ImportPrivateKeyToP2TR(legacyWallet.privateKey, &chaincfg.TestNet3Params)
	fmt.Println(p2tr)
}

func TestGetAddressTransactions(t *testing.T) {
	httpx.Get("https://mempool.space/testnet/api/address/tb1px60kt67jexk8ns7h0lmxalz9zz3xst72vmufd3gk6vvsma2ezrzs87nfja/txs", func(resp *http.Response) {
		bs, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bs))
	}, func(err error) {
		if err != nil {
			panic(err)
		}
	})
}
