package main

import (
	"flag"
	"fmt"

	"github.com/FactomProject/factomd/common/primitives"
)

func main() {
	var (
		simple = flag.Bool("s", false, "No indicator on which is public, and which is private")
	)

	flag.Parse()

	p := primitives.RandomPrivateKey()

	if *simple {
		fmt.Printf("%s\n", p.PrivateKeyString())
		fmt.Printf("%s\n", p.PublicKeyString())
	} else {
		fmt.Printf("Priv: %s\n", p.PrivateKeyString())
		fmt.Printf("Pub : %s\n", p.PublicKeyString())
	}
}

