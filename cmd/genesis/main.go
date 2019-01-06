package main

import (
	"bufio"
	"fmt"
	"github.com/stellar/go/clients/horizon"
	"net/http"
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
	issuerKeypair, err := keypair.Random()
	check(err)
	f, _ := os.Create("./key")
	w := bufio.NewWriter(f)

	// Create new file and write seed key to file
	{
		_, err := w.WriteString(issuerKeypair.Seed())
		check(err)
		check(w.Flush())
	}

	// Ask XLM from stellar dev network
	{
		url := fmt.Sprintf("https://friendbot.stellar.org?addr=%s", issuerKeypair.Address())
		resp, err := http.Get(url)
		check(err)

		if resp.StatusCode == 200 {
			fmt.Println("Success")
		}
	}

	// Checking balance
	{
		testNetClient := horizon.DefaultTestNetClient
		account, err := testNetClient.LoadAccount(issuerKeypair.Address())
		check(err)

		for _, balance := range account.Balances {
			fmt.Printf("%s: %s \n", balance.Code, balance.Balance)
		}

	}





}
