// Copyright (c) 2016 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wallet

import (
	"github.com/valhallacoin/vhcd/vhcutil"
	"github.com/valhallacoin/vhcwallet/errors"
	"github.com/valhallacoin/vhcwallet/wallet/walletdb"
	"github.com/valhallacoin/vhcwallet/wallet/udb"
)

// StakePoolUserInfo returns the stake pool user information for a user
// identified by their P2SH voting address.
func (w *Wallet) StakePoolUserInfo(userAddress vhcutil.Address) (*udb.StakePoolUser, error) {
	const op errors.Op = "wallet.StakePoolUserInfo"

	switch userAddress.(type) {
	case *vhcutil.AddressPubKeyHash: // ok
	case *vhcutil.AddressScriptHash: // ok
	default:
		return nil, errors.E(op, errors.Invalid, "address must be P2PKH or P2SH")
	}

	var user *udb.StakePoolUser
	err := walletdb.View(w.db, func(tx walletdb.ReadTx) error {
		stakemgrNs := tx.ReadBucket(wstakemgrNamespaceKey)
		var err error
		user, err = w.StakeMgr.StakePoolUserInfo(stakemgrNs, userAddress)
		return err
	})
	if err != nil {
		return nil, errors.E(op, err)
	}
	return user, nil
}
