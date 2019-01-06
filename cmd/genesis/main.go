package main

import (
	"bufio"
	"fmt"
	"github.com/stellar/go/build"
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
		fmt.Printf("%s \n", issuerKeypair.Address())
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

	// Create Distributor account
	distributorKeypair, err := keypair.Random()
	fmt.Println(distributorKeypair.Address())
	{
		testNetClient := horizon.DefaultTestNetClient

		// Create new account Tx, using Issuer account
		tx, err := build.Transaction(
			build.SourceAccount{AddressOrSeed: issuerKeypair.Address()},
			build.TestNetwork,
			build.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
			build.CreateAccount(
				build.Destination{
					AddressOrSeed: distributorKeypair.Address(),
				},
				// This mean XLM
				build.NativeAmount{
					Amount: "100",
				},
			),
		)
		check(err)
		txe, err := tx.Sign(issuerKeypair.Seed())
		check(err)
		txeB64, err := txe.Base64()
		check(err)
		resp, err := testNetClient.SubmitTransaction(txeB64)
		check(err)

		fmt.Println(resp.Ledger)
	}

}
