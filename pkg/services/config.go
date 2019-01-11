package services

import (
	"encoding/json"
	"io/ioutil"
)

type Candidate struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}
type Configuration struct {
	DistributorAddress string `json:"distributorAddress"`
	DistributorSecret  string `json:"distributorSecret"`

	AssetName     string `json:"assetName"`
	IssuerAddress string `json:"issuerAddress"`

	Candidates []Candidate `json:"candidates"`
}

func ReadConfiguration() Configuration {

	if result, err := ioutil.ReadFile("./config.json"); err != nil {
		panic(err)
	} else {
		var configuration Configuration
		err := json.Unmarshal(result, &configuration)
		if err != nil {
			panic(err)
		}
		return configuration
	}
}
