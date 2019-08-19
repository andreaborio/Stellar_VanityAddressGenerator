package main

import (
	"github.com/stellar/go/keypair"
	"log"
	"os"
	"strings"
)

func main() {
	var pair, _ = keypair.Random()
	var len = len(os.Args) - 1
	for len > 0 {
		for pair, _ = keypair.Random(); !(strings.Contains(pair.Address(), strings.ToUpper(os.Args[len]))); pair, _ = keypair.Random() {
		}
		log.Println("\n\r Hey! Pair found!! \n\r", pair.Address(), pair.Seed())
		len--
	}
}
