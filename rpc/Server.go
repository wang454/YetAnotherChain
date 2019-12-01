package main

import (
	"YetAnotherChain/core"
	"encoding/json"
	"io"
	"net/http"
)

var blockchain *core.Blockchain

func run() {
	http.HandleFunc("/blockchain/get", blockchainGetHandler)

	// http://localhost:9999/blockchain/write?data=send%201%20btc%20to%20chao
	http.HandleFunc("/blockchain/write", blockchainWriterHandler)

	http.ListenAndServe("localhost:9999", nil)
}

func blockchainGetHandler(w http.ResponseWriter, r *http.Request){
	bytes, error := json.Marshal(blockchain)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}

	io.WriteString(w, string(bytes))
}

func blockchainWriterHandler(w http.ResponseWriter, r *http.Request){
	blockData := r.URL.Query().Get("data")
	blockchain.SendData(blockData)
	blockchainGetHandler(w, r)
}

func main() {
	blockchain = core.NewBlockchain()
	run()
}