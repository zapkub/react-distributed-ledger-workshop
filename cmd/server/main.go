package main

import (
	"io/ioutil"

	"github.com/stellar/go/keypair"
	pkg "github.com/zapkub/react-distributed-ledger-workshop/pkg/services"
)

func main() {
	if result, err := ioutil.ReadFile("./key"); err != nil {
		panic(err)
	} else {

		// Read Masterkey
		var b [32]byte
		copy(b[:], result)
		kp, _ := keypair.FromRawSeed([32]byte(b))
		pkg.NewApplicationServer(pkg.ApplicationConfig{
			MasterKey: kp,
		})
	}

}
