// Copyright (c) 2019 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wallet

import (
	"context"

	"github.com/valhallacoin/vhcd/chaincfg/chainhash"
	"github.com/valhallacoin/vhcd/vhcutil"
	"github.com/valhallacoin/vhcd/gcs"
	"github.com/valhallacoin/vhcd/wire"
)

// mockNetwork implements all methods of NetworkBackend, returning zero values
// without error.  It may be embedded in a struct to create another
// NetworkBackend which dispatches to particular implementations of the methods.
type mockNetwork struct{}

func (mockNetwork) GetBlocks(ctx context.Context, blockHashes []*chainhash.Hash) ([]*wire.MsgBlock, error) {
	return nil, nil
}
func (mockNetwork) GetCFilters(ctx context.Context, blockHashes []*chainhash.Hash) ([]*gcs.Filter, error) {
	return nil, nil
}
func (mockNetwork) GetHeaders(ctx context.Context, blockLocators []*chainhash.Hash, hashStop *chainhash.Hash) ([]*wire.BlockHeader, error) {
	return nil, nil
}
func (mockNetwork) PublishTransactions(ctx context.Context, txs ...*wire.MsgTx) error { return nil }
func (mockNetwork) LoadTxFilter(ctx context.Context, reload bool, addrs []vhcutil.Address, outpoints []wire.OutPoint) error {
	return nil
}
func (mockNetwork) Rescan(ctx context.Context, blocks []chainhash.Hash, r RescanSaver) error {
	return nil
}
func (mockNetwork) StakeDifficulty(ctx context.Context) (vhcutil.Amount, error) { return 0, nil }
