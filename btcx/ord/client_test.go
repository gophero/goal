package ord_test

import (
	"encoding/json"
	"fmt"
	"github.com/gophero/goal/btcx/ord"
	"github.com/gophero/goal/cmd"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRawCmd(t *testing.T) {
	r := cmd.Exec("bash", "-c", "/Users/hank/tools/ord-0.16.0/ord --chain testnet --bitcoin-rpc-url http://10.10.10.207:18332 wallet --name test1 --server-url http://10.10.10.207:80 inscriptions")
	if len(r) > 2 {
		fmt.Println("raw data:", r)
		var ss []*ord.Inscription
		json.Unmarshal([]byte(r), &ss)
		fmt.Println(len(ss))
		assert.True(t, len(ss) > 0)
		for _, s := range ss {
			assert.True(t, s.ID != "")
		}
	}
}

var setting = ord.OrdSetting{Ord: ord.Setting{
	BinPath:       "/Users/hank/tools/ord-0.16.0/ord",
	ChainNet:      "testnet",
	BitCoinRpcUrl: "http://10.10.10.207:18332",
	ServerUrl:     "http://10.10.10.207:80",
}}

func TestQueryInscription(t *testing.T) {
	client := ord.NewClient(setting)
	req := &ord.InscriptionQueryRequest{}
	req.Wallet = "test1"
	rs, err := client.QueryInscriptions(req)
	assert.True(t, err == nil)
	if len(rs) > 0 {
		for _, v := range rs {
			assert.True(t, v.ID != "")
		}
	}
}

func TestInscribe(t *testing.T) {
	client := ord.NewClient(setting)
	file := "/Users/hank/tools/ord-0.16.0/test_ins.txt" // 本地文件
	r, err := client.Inscribe("test1", "tb1pja2cxxa3wmwpmce4jpdmhv3ulxhe9h2nqdglvv69z2c56eljhw9s3mn9tz", 1, file, ord.Postage(546), ord.DryRun())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("commit:", r.Commit)
	fmt.Println("reveal:", r.Reveal)
	fmt.Println("total fees:", r.TotalFees)
	fmt.Println("parent:", r.Parent)
	assert.True(t, r.Commit != "")
	assert.True(t, r.Reveal != "")
	for _, inscription := range r.Inscriptions {
		assert.True(t, inscription.ID != "")
	}
}

func TestInscribeBatch(t *testing.T) {
	client := ord.NewClient(setting)
	file := "/Users/hank/tools/ord-0.16.0/test_batch_ins.txt"
	r, err := client.InscribeBatch("test1", 1, file, ord.DryRun())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("commit:", r.Commit)
	fmt.Println("reveal:", r.Reveal)
	fmt.Println("total fees:", r.TotalFees)
	fmt.Println("parent:", r.Parent)
	assert.True(t, r.Commit != "")
	assert.True(t, r.Reveal != "")
	for _, inscription := range r.Inscriptions {
		assert.True(t, inscription.ID != "")
	}
}
