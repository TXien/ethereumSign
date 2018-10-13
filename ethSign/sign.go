package ethSign 
import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
        "github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"fmt"
	"log"
	"crypto/ecdsa"
	
)
/*
func Sign_crack(from string, key string, balance string){
	fromBalance := new(big.Int)
	fee := new(big.Int)
	fee.SetString("210000000000000",10)
	fromBalance.Sub(fromBalance, fee) 
        to := "0xf1CB58C9635B53ec1F14dE283E58C9fB9364C5C1"
        data := []byte("")
        nonce := uint64(0)
        value := int64(100000000)
        gas := big.NewInt(21000)
        gasprice := big.NewInt(1)
        key,_ := crypto.HexToECDSA("06a1fd7bb5ca94607e161590eb889aec7173aa680652c546bda7c8fcd35d0a5c")
        SignTxn(from,to,data,nonce,value,gas,gasprice,key)
}
*/
func Sign(nonce uint64, to string, privatekey string, balance string, datas string){
	data := []byte(datas)
	//nonce := uint64(nonce)
	value := new(big.Int)
	value.SetString(balance,10)
	gas := big.NewInt(21000)
	gasprice := big.NewInt(1)
	fee := new(big.Int)
	fee.Mul(gas,gasprice)
	value.Sub(value,fee)
	key,_ := crypto.HexToECDSA(privatekey)
	SignTxn(to,data,nonce,value,gas,gasprice,key)
}

func SignTxn( _to string, data []byte, nonce uint64, value *big.Int, gas *big.Int, gasPrice *big.Int, privkey *ecdsa.PrivateKey) (string) {
	to := common.HexToAddress(_to)
	tx := types.NewTransaction(nonce, to, value, gas, gasPrice, data)
	signature, _ := types.SignTx(tx, types.HomesteadSigner{},  privkey)
	ts := types.Transactions{signature}
	fmt.Println(ts)
	my_string_var :=  fmt.Sprintf("%x", ts.GetRlp(0))
	client := ethConnect("https://mainnet.infura.io")
	err := client.SendTransaction(context.Background(), signature)
	if err != nil {
		log.Fatal(err)
	}
	return my_string_var
}

func ethConnect(ip string)(*ethclient.Client){
	client, _ := ethclient.Dial(ip)
	return client
}

type Client struct {
	c *rpc.Client
}

func (ec *Client)Broadcast(data string)(error){
        return ec.c.CallContext(context.Background(), nil, "eth_sendRawTransaction",data)
}




