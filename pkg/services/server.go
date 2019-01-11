package services

import (
	"github.com/gin-gonic/gin/json"
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
		switch r.Method {
		case "GET":
			{
				// Create new keypair and asking for XLM from Steller dev network
				w.WriteHeader(http.StatusOK)
				configJSON, _ := json.Marshal(ClientConfiguration{
					IssuerAddress: configuration.IssuerAddress,
					Candidates:    configuration.Candidates,
					AssetName:     configuration.AssetName,
				})
				if _, err := w.Write(configJSON); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
				}

			}
		default:
			{
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
		}

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
