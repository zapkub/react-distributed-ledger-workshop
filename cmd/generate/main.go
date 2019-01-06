package main

import (
	"fmt"
	"github.com/stellar/go/keypair"
)


func main() {


	var seed string
	var addr string
	var prefix string

	for prefix != "8FOX" {
		kp, _ := keypair.Random()
		addr =kp.Address()
		seed = kp.Seed()
		prefix = string([]rune(addr)[len(addr) - 4:len(addr)])
		fmt.Println(prefix)
	}

	fmt.Println(addr)
	fmt.Println(seed)

}
