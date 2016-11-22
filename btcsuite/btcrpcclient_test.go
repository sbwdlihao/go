package btcsuite

import (
	"testing"
	"fmt"
	rpc "github.com/btcrpcclient"
	"github.com/btcsuite/btcd/btcjson"
	"os"
)

var client *rpc.Client

func TestMain(m *testing.M) {
	config := &rpc.ConnConfig{
		Host: "127.0.0.1:19001",
		User: "admin1",
		Pass: "123",
		HTTPPostMode: true,
		DisableTLS: true,
	}
	var err error
	client, err = rpc.New(config, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	os.Exit(m.Run())
}

func TestBlockTemplate0(t *testing.T) {
	request := &btcjson.TemplateRequest{

	}
	result, err := client.GetBlockTemplate(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", result)
}

// 当前的bitcoind尚未实现serverlist，coinbasetxn
func TestBlockTemplate1(t *testing.T) {
	request := &btcjson.TemplateRequest{
		Capabilities: []string{"serverlist", "coinbasetxn"},
	}
	result, err := client.GetBlockTemplate(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", result)
}

