package main
import "github.com/ethereum/go-ethereum/core/types"
import "github.com/ethereum/go-ethereum/common"
import "github.com/ethereum/go-ethereum/crypto"
import "math/big"
import "fmt"
//import "os"
//import "encoding/hex"
import "crypto/ecdsa"
//import "encoding/json"
//import "github.com/ethereum/go-ethereum/cmd/utils"
type GethTxn struct {
  To   string     `json:"to"`
  From string     `json:"from"`
  Gas string      `json:"gas"`
  GasPrice string `json:"gasPrice"`
  Value string    `json:"value"`
  Data string     `json:"input"`
}

func main(){
	from := "0xf1CB58C9635B53ec1F14dE283E58C9fB9364C5C1"
	to := "0xf1CB58C9635B53ec1F14dE283E58C9fB9364C5C1"
	data := []byte("")
	nonce := uint64(1203)
	value := int64(100000000)
	gas := big.NewInt(21000)
	gasprice := big.NewInt(1)
	key,_ := crypto.HexToECDSA("06a1fd7bb5ca94607e161590eb889aec7173aa680652c546bda7c8fcd35d0a5c")
	SignTxn(from,to,data,nonce,value,gas,gasprice,key)
}

func SignTxn(from string, _to string, data []byte, nonce uint64, value int64, gas *big.Int, gasPrice *big.Int, privkey *ecdsa.PrivateKey)/* (*GethTxn, error)*/ {

//  var parsed_tx = new(GethTxn)
  var amount = big.NewInt(value)
  
  to := common.HexToAddress(_to)

//signer := types.NewEIP155Signer(nil)
  tx := types.NewTransaction(nonce, to, amount, gas, gasPrice, data)
  //signature, _ := crypto.Sign(tx.SigHash(signer).Bytes(), privkey)
  signature, _ := types.SignTx(tx, types.HomesteadSigner{}/*types.NewEIP155Signer(big.NewInt(1))*/,  privkey)
//  fmt.Println(signature)

  ts := types.Transactions{signature}
  my_string_var :=  fmt.Sprintf("%x", ts.GetRlp(0))
  fmt.Println(my_string_var)
/*
  signed_tx, _ := tx.WithSignature(signer, signature)

  json_tx, _ := signed_tx.MarshalJSON()
  _ = json.Unmarshal(json_tx, parsed_tx)
  parsed_tx.From = from
  fmt.Println("data", parsed_tx.Data)
  return parsed_tx, nil
*/
}
