// Copyright (c) 2014 The btcsuite developers
// Copyright (c) 2015-2018 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

// BlockOneLedgerMainNet is the block one output ledger for the main
// network.
var BlockOneLedgerMainNet = []*TokenPayout{}

// BlockOneLedgerTestNet is the block one output ledger for testnet version 3.
var BlockOneLedgerTestNet = []*TokenPayout{
	{"TsnHhYH48JY4Li5nC25Y3pcvcvgFehRzjuT", 100000 * 1e8},
}

// BlockOneLedgerSimNet is the block one output ledger for the simulation
// network.  See "Valhalla organization related parameters" in simnetparams.go for
// information on how to spend these outputs.
var BlockOneLedgerSimNet = []*TokenPayout{
	{"Sshw6S86G2bV6W32cbc7EhtFy8f93rU6pae", 100000 * 1e8},
	{"SsjXRK6Xz6CFuBt6PugBvrkdAa4xGbcZ18w", 100000 * 1e8},
	{"SsfXiYkYkCoo31CuVQw428N6wWKus2ZEw5X", 100000 * 1e8},
}

// BlockOneLedgerRegNet is the block one output ledger for the regression test
// network.  See "Valhalla organization related parameters" in regnetparams.go for
// information on how to spend these outputs.
var BlockOneLedgerRegNet = []*TokenPayout{
	{"RsKrWb7Vny1jnzL1sDLgKTAteh9RZcRr5g6", 100000 * 1e8},
	{"Rs8ca5cDALtsMVD4PV3xvFTC7dmuU1juvLv", 100000 * 1e8},
	{"RsHzbGt6YajuHpurtpqXXHz57LmYZK8w9tX", 100000 * 1e8},
}
