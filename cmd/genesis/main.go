package main

import (
	"bufio"
	"os"

	"github.com/stellar/go/keypair"
)

// Genesis cmd
// This excutable will create
// - new Master keypair as distributor
// - create new Issuer keypair and send Asset to distributor

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Create new keypair
	kp, err := keypair.Random()
	check(err)
	f, _ := os.Create("./key")
	w := bufio.NewWriter(f)

	// Create new file and write seed key to file
	{
		_, err := w.WriteString(kp.Seed())
		check(err)
		check(w.Flush())
	}

	// Ask XLM from stellar dev network
	{

	}

}
