package main

import (
	"encoding/hex"
	"fmt"
	"encoding/json"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/core/types"

)

func main() {
	var txSub *types.Transaction

	raxTx := "0xf87c108504a817c800833d0900940000000000000000000000000000000000000000880de0b6b3a76400008c48656c6c6f20576f726c64218303c4a4a0d956a764f81866d9b16d744cd0d513daeac0a5831215a2d34ed9a0163b292462a05d78bfeaed709951333506f5043fdcb3b75ebecb4d35d648416149da37273f84"
	rawTxBytes, err := hex.DecodeString(raxTx[2:])
	if err != nil {
		fmt.Println(err)
	}

	err = rlp.DecodeBytes(rawTxBytes, &txSub)
	if err != nil {
		fmt.Println(err)
	}

	txbyte, err := json.Marshal(txSub)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(txbyte))

	fmt.Println("Submitted transaction", "fullhash", txSub.Hash().Hex(), "recipient", txSub.To())
}

