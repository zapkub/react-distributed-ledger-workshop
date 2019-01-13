package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/keypair"
	"github.com/zapkub/react-distributed-ledger-workshop/pkg/utils"
	"net/http"
	"os"
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
	{
		testNetClient := horizon.DefaultTestNetClient

		// Create new account Tx, using Issuer account
		tx, err := build.Transaction(
			build.SourceAccount{AddressOrSeed: issuerKeypair.Address()},
			build.TestNetwork,
			build.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
			//build.SetOptions(
			//	build.SetAuthRequired(),
			//),
			build.CreateAccount(
				build.Destination{
					AddressOrSeed: distributorKeypair.Address(),
				},
				// This mean XLM
				build.NativeAmount{
					Amount: "5000",
				},
			),
			build.MemoText{
				Value: "Distributor account",
			},
		)
		check(err)
		txe, err := tx.Sign(issuerKeypair.Seed())
		check(err)
		txeB64, err := txe.Base64()
		check(err)
		resp, err := testNetClient.SubmitTransaction(txeB64)
		check(err)

		fmt.Printf("Create distributor account tx hash: %s\n", resp.Hash)
	}

	v4bpAsset := build.CreditAsset("V4BP", issuerKeypair.Address())
	{

		// Create Custom Asset

		// Trust Issuer line
		tx, err := build.Transaction(
			build.SourceAccount{
				AddressOrSeed: distributorKeypair.Address(),
			},
			build.AutoSequence{horizon.DefaultTestNetClient},
			build.TestNetwork,
			build.Trust(v4bpAsset.Code, v4bpAsset.Issuer, build.Limit("300000")),
		)

		check(err)
		txe, err := tx.Sign(distributorKeypair.Seed())
		check(err)
		txe64, err := txe.Base64()
		check(err)
		resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txe64)
		check(err)

		fmt.Printf("Change distributor trust tx hash: %s \n", resp.Hash)

	}

	// Issuer transfer custom asset to distributor
	// and then invalidate issuer
	// and create candidate account for 4 members
	// - lisa
	// - jennie
	// - jisoo
	// - rose

	lisaKeypair, _ := keypair.Random()
	jennieKeypair, _ := keypair.Random()
	jisooKeypair, _ := keypair.Random()
	roseKeypair, _ := keypair.Random()
	{

		tx, err := build.Transaction(
			build.SourceAccount{
				AddressOrSeed: issuerKeypair.Seed(),
			},
			build.AutoSequence{horizon.DefaultTestNetClient},
			build.TestNetwork,
			build.Payment(

				// Transfer to who
				build.Destination{
					AddressOrSeed: distributorKeypair.Address(),
				},

				// Asset to transfer
				build.CreditAmount{
					Code:   v4bpAsset.Code,
					Issuer: issuerKeypair.Address(),
					Amount: "300000",
				},
			),

			build.CreateAccount(
				build.Destination{
					AddressOrSeed: lisaKeypair.Address(),
				},
				build.NativeAmount{
					Amount: "5",
				},
			),

			build.CreateAccount(
				build.Destination{
					AddressOrSeed: jennieKeypair.Address(),
				},
				build.NativeAmount{
					Amount: "5",
				},
			),

			build.CreateAccount(
				build.Destination{
					AddressOrSeed: jisooKeypair.Address(),
				},
				build.NativeAmount{
					Amount: "5",
				},
			),

			build.CreateAccount(
				build.Destination{
					AddressOrSeed: roseKeypair.Address(),
				},
				build.NativeAmount{
					Amount: "5",
				},
			),

			build.SetOptions(
				build.MasterWeight(0),
			),

			build.Trust(v4bpAsset.Code, v4bpAsset.Issuer, build.Limit("300000"),
				build.SourceAccount{
					AddressOrSeed: lisaKeypair.Address(),
				},
			),

			build.Trust(v4bpAsset.Code, v4bpAsset.Issuer, build.Limit("300000"),
				build.SourceAccount{
					AddressOrSeed: jisooKeypair.Address(),
				},
			),

			build.Trust(v4bpAsset.Code, v4bpAsset.Issuer, build.Limit("300000"),
				build.SourceAccount{
					AddressOrSeed: jennieKeypair.Address(),
				},
			),

			build.Trust(v4bpAsset.Code, v4bpAsset.Issuer, build.Limit("300000"),
				build.SourceAccount{
					AddressOrSeed: roseKeypair.Address(),
				},
			),
		)

		check(err)
		txe, err := tx.Sign(
			issuerKeypair.Seed(),
			lisaKeypair.Seed(),
			jisooKeypair.Seed(),
			jennieKeypair.Seed(),
			roseKeypair.Seed(),
		)

		txe64, err := txe.Base64()
		check(err)
		resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txe64)
		check(err)
		fmt.Printf("Send asset to distributor, create 4 candidate, and invalidate issuer account tx hash :%s\n", resp.Hash)

	}

	fmt.Printf(
		"Distributor public: %s\nDistributor secret: %s\n",
		distributorKeypair.Address(),
		distributorKeypair.Seed(),
	)

	// Create new file and write seed key to file
	{
		fmt.Printf("%s \n", issuerKeypair.Address())

		config := utils.Configuration{
			DistributorAddress: distributorKeypair.Address(),
			DistributorSecret:  distributorKeypair.Seed(),
			AssetName:          v4bpAsset.Code,
			IssuerAddress:      issuerKeypair.Address(),
			Candidates: []utils.Candidate{
				{
					Address: lisaKeypair.Address(),
					Name:    "Lisa",
				},
				{
					Address: jisooKeypair.Address(),
					Name:    "Jisoo",
				},
				{
					Address: roseKeypair.Address(),
					Name:    "Rosé",
				},
				{
					Address: jennieKeypair.Address(),
					Name:    "Jennie",
				},
			},
		}

		// Write config file
		{
			configJSON, err := json.Marshal(config)

			f, _ := os.Create("./config.distributor.json")
			w := bufio.NewWriter(f)
			_, err = w.WriteString(string(configJSON))
			check(err)
			check(w.Flush())
		}

		// Write config for client (No distributor information)
		{
			clientConfig := utils.Configuration{
				Candidates:    config.Candidates,
				AssetName:     config.AssetName,
				IssuerAddress: config.IssuerAddress,
			}

			configJSON, err := json.Marshal(clientConfig)
			f, _ := os.Create("./config.client.json")
			w := bufio.NewWriter(f)
			_, err = w.WriteString(string(configJSON))
			check(err)
			check(w.Flush())
		}

	}

}
