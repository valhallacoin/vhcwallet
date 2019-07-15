chaincfg
========

[![Build Status](http://img.shields.io/travis/valhallacoin/vhcd.svg)](https://travis-ci.org/valhallacoin/vhcd)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/valhallacoin/vhcd/chaincfg)

Package chaincfg defines chain configuration parameters for the three standard
Valhalla networks and provides the ability for callers to define their own custom
Valhalla networks.

Although this package was primarily written for vhcd, it has intentionally been
designed so it can be used as a standalone package for any projects needing to
use parameters for the standard Valhalla networks or for projects needing to
define their own network.

## Sample Use

```Go
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/valhallacoin/vhcd/vhcutil"
	"github.com/valhallacoin/vhcd/chaincfg"
)

var testnet = flag.Bool("testnet", false, "operate on the testnet Valhalla network")

// By default (without -testnet), use mainnet.
var chainParams = &chaincfg.MainNetParams

func main() {
	flag.Parse()

	// Modify active network parameters if operating on testnet.
	if *testnet {
		chainParams = &chaincfg.TestNetParams
	}

	// later...

	// Create and print new payment address, specific to the active network.
	pubKeyHash := make([]byte, 20)
	addr, err := btcutil.NewAddressPubKeyHash(pubKeyHash, chainParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(addr)
}
```

## Installation and Updating

```bash
$ go get -u github.com/valhallacoin/vhcd/chaincfg
```

## License

Package chaincfg is licensed under the [copyfree](http://copyfree.org) ISC
License.
