package main

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
	// Create new Issuer keypair
	// Ask XLM from stellar dev network
	{

	}

	// Checking Issuer balance
	{

	}

	// Create Distributor account
	{

		// Create new account Tx, using Issuer account
	}

	//var v4bpAsset build.Asset
	{

		// Create Issuer Custom Asset
		// Distributor change to Trust Issuer line

	}

	// Issuer transfer custom asset to distributor
	// and then invalidate issuer
	// and create candidate account for 4 members
	// - lisa
	// - jennie
	// - jisoo
	// - rose
	{

	}

	// Create new file and write seed key to file
	{
		//fmt.Printf("%s \n", issuerKeypair.Address())
		//
		//config := services.Configuration{
		//	DistributorAddress: distributorKeypair.Address(),
		//	DistributorSecret:  distributorKeypair.Seed(),
		//	AssetName:          v4bpAsset.Code,
		//	IssuerAddress:      issuerKeypair.Address(),
		//	Candidates: []services.Candidate{
		//		{
		//			Address: lisaKeypair.Address(),
		//			Name:    "Lisa",
		//		},
		//		{
		//			Address: jisooKeypair.Address(),
		//			Name:    "Jisoo",
		//		},
		//		{
		//			Address: roseKeypair.Address(),
		//			Name:    "Rosé",
		//		},
		//		{
		//			Address: jennieKeypair.Address(),
		//			Name:    "Jennie",
		//		},
		//	},
		//}

		//configJSON, err := json.Marshal(config)
		//
		//f, _ := os.Create("./config.json")
		//w := bufio.NewWriter(f)
		//_, err = w.WriteString(string(configJSON))
		//check(err)
		//check(w.Flush())
	}

}
