package pkg

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/stellar/go/keypair"
)

type ApplicationConfig struct {
	MasterKey *keypair.Full
}
type ApplicationRepositoryOptions interface {
	Config() *ApplicationConfig
}

func walletHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		{
			// Create new keypair and asking for XLM from Steller dev network
		}
	default:
		{
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}

}
func NewApplicationServer(config ApplicationConfig) {

	http.HandleFunc("/wallet", walletHandler)

	target := "http://localhost:3001"
	url, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(url)
	http.HandleFunc("/", proxy.ServeHTTP)
	http.ListenAndServe(":3000", nil)
}
