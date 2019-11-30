package main

import "YetAnotherChain/core"

func main() {
	bc := core.NewBlockchain()
	bc.SendData("Send 1 BTC to Osa")
	bc.SendData("Send 1 ETH to Chao")

	bc.Print()
}
