package main

import (
	"log"
)

func main() {
	btc, err := bitcoin.getBitcoin()
	if err != nil {
		log.Fatal(err)
	}
}
