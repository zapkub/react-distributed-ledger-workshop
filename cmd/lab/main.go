package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/keypair"
	"net/http"
	"os"
)

func panicErr(err error) {
	if err != nil {
		if e, ok := err.(*horizon.Error); ok {
			b, _ := e.Problem.Extras["result_codes"].MarshalJSON()

			fmt.Println(string(b))
		} else {
			panic(err)
		}
	}
}

func main() {
	fmt.Println("Hello world")

	type Candidate struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	}

	type ElectionConfig struct {
		IssuerAddress string      `json:"issuerAddress"`
		AssetName     string      `json:"assetName"`
		Candidates    []Candidate `json:"candidates"`
	}

	electConfig := ElectionConfig{}

	masterAccountKeypair, _ := keypair.Random()

	electConfig.IssuerAddress = masterAccountKeypair.Address()

	distributorAccountKeypair, _ := keypair.Random()
	v4bpAsset := build.CreditAsset("V4BP", masterAccountKeypair.Address())

	electConfig.AssetName = v4bpAsset.Code

	{

		url := fmt.Sprintf("https://friendbot.stellar.org?addr=%s", masterAccountKeypair.Address())
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		if resp.StatusCode == http.StatusOK {
			fmt.Printf("Account has been created ! %s\n", masterAccountKeypair.Address())
		}

	}

	candidateNames := []string{
		"lisa", "jennie", "jisoo", "rose",
	}

	// สร้าง Operation
	operations := []build.TransactionMutator{
		build.SourceAccount{
			AddressOrSeed: masterAccountKeypair.Address(),
		},
		build.TestNetwork,
		build.AutoSequence{
			SequenceProvider: horizon.DefaultTestNetClient,
		},

		build.CreateAccount(
			build.Destination{
				AddressOrSeed: distributorAccountKeypair.Address(),
			},
			build.NativeAmount{
				Amount: "50",
			},
		),

		build.Trust(
			v4bpAsset.Code,
			v4bpAsset.Issuer,
			build.SourceAccount{
				AddressOrSeed: distributorAccountKeypair.Address(),
			},
		),

		build.Payment(
			build.Destination{
				AddressOrSeed: distributorAccountKeypair.Address(),
			},

			build.CreditAmount{
				Issuer: v4bpAsset.Issuer,
				Code:   v4bpAsset.Code,
				Amount: "300000",
			},
			build.SourceAccount{
				AddressOrSeed: masterAccountKeypair.Address(),
			},
		),
	}

	candidateSecrets := []string{}

	for _, name := range candidateNames {

		candidateKeypair, _ := keypair.Random()

		fmt.Printf("Generate wallet for %s: %s\n", name, candidateKeypair.Address())

		electConfig.Candidates = append(electConfig.Candidates, Candidate{
			Address: candidateKeypair.Address(),
			Name:    name,
		})

		candidateSecrets = append(candidateSecrets, candidateKeypair.Seed())

		operations = append(operations, build.CreateAccount(
			build.Destination{
				AddressOrSeed: candidateKeypair.Address(),
			},
			build.NativeAmount{
				Amount: "50",
			},
		))
		operations = append(operations, build.Trust(
			v4bpAsset.Code,
			v4bpAsset.Issuer,
			build.SourceAccount{
				AddressOrSeed: candidateKeypair.Address(),
			},
		))
	}

	// สร้าง Transaction
	tx, err := build.Transaction(operations...)
	if err != nil {
		panic(err)
	}

	// Sign transaction
	candidateSecrets = append(candidateSecrets, masterAccountKeypair.Seed())
	candidateSecrets = append(candidateSecrets, distributorAccountKeypair.Seed())
	txe, err := tx.Sign(candidateSecrets...)
	if err != nil {
		panic(err)
	}
	txe64, err := txe.Base64()

	// Submit transaction
	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txe64)

	if err != nil {
		panicErr(err)
	}
	fmt.Printf("Issuer addr: %s\n", masterAccountKeypair.Address())
	fmt.Printf("Distributor addr: %s\nsecret: %s\n", distributorAccountKeypair.Address(), distributorAccountKeypair.Seed())
	fmt.Printf("Tx hash: %s\n", resp.Hash)

	configJSON, err := json.Marshal(&electConfig)
	if err != nil {
		panic(err)
	}

	f, _ := os.Create("./config.json")
	w := bufio.NewWriter(f)
	_, err = w.WriteString(string(configJSON))
	if err != nil {
		panic(err)
	}
	err = w.Flush()
	if err != nil {
		panic(err)
	}

}
