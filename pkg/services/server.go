package services

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type ClientConfiguration struct {
	AssetName     string      `json:"assetName"`
	IssuerAddress string      `json:"issuerAddress"`
	Candidates    []Candidate `json:"candidates"`
}

func configurationHandler(configuration Configuration) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		panic("Implement me")
	}
}
func NewApplicationServer(config Configuration) {

	http.HandleFunc("/configuration", configurationHandler(config))
	target := "http://localhost:3001"
	viewURL, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(viewURL)
	proxy.FlushInterval = 100 * time.Millisecond
	http.HandleFunc("/", proxy.ServeHTTP)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}
