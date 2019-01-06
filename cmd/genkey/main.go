package main

import (
	"bufio"
	"github.com/stellar/go/keypair"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}


func main() {
	if kp, err := keypair.Random();err == nil {
		f, _ := os.Create("./key")
		w := bufio.NewWriter(f)
		{
			_, err := w.WriteString(kp.Seed())
			check(err)
		}
		{
			check( w.Flush() )
		}
	} else {
		panic(err)
	}
}
