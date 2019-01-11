package main

import (
	pkg "github.com/zapkub/react-distributed-ledger-workshop/pkg/services"
)

func main() {
	configuration := pkg.ReadConfiguration()
	pkg.NewApplicationServer(configuration)
}
