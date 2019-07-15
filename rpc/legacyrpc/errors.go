// Copyright (c) 2013-2015 The btcsuite developers
// Copyright (c) 2016-2018 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package legacyrpc

import (
	"fmt"

	"github.com/valhallacoin/vhcd/vhcjson"
	"github.com/valhallacoin/vhcwallet/errors"
)

func convertError(err error) *vhcjson.RPCError {
	if err, ok := err.(*vhcjson.RPCError); ok {
		return err
	}

	code := vhcjson.ErrRPCWallet
	if err, ok := err.(*errors.Error); ok {
		switch err.Kind {
		case errors.Bug:
			code = vhcjson.ErrRPCInternal.Code
		case errors.Encoding:
			code = vhcjson.ErrRPCInvalidParameter
		case errors.Locked:
			code = vhcjson.ErrRPCWalletUnlockNeeded
		case errors.Passphrase:
			code = vhcjson.ErrRPCWalletPassphraseIncorrect
		case errors.NoPeers:
			code = vhcjson.ErrRPCClientNotConnected
		case errors.InsufficientBalance:
			code = vhcjson.ErrRPCWalletInsufficientFunds
		}
	}
	return &vhcjson.RPCError{
		Code:    code,
		Message: err.Error(),
	}
}

func rpcError(code vhcjson.RPCErrorCode, err error) *vhcjson.RPCError {
	return &vhcjson.RPCError{
		Code:    code,
		Message: err.Error(),
	}
}

func rpcErrorf(code vhcjson.RPCErrorCode, format string, args ...interface{}) *vhcjson.RPCError {
	return &vhcjson.RPCError{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}

// Errors variables that are defined once here to avoid duplication.
var (
	errUnloadedWallet = &vhcjson.RPCError{
		Code:    vhcjson.ErrRPCWallet,
		Message: "request requires a wallet but wallet has not loaded yet",
	}

	errRPCClientNotConnected = &vhcjson.RPCError{
		Code:    vhcjson.ErrRPCClientNotConnected,
		Message: "disconnected from consensus RPC",
	}

	errNoNetwork = &vhcjson.RPCError{
		Code:    vhcjson.ErrRPCClientNotConnected,
		Message: "disconnected from network",
	}

	errAccountNotFound = &vhcjson.RPCError{
		Code:    vhcjson.ErrRPCWalletInvalidAccountName,
		Message: "account not found",
	}

	errAddressNotInWallet = &vhcjson.RPCError{
		Code:    vhcjson.ErrRPCWallet,
		Message: "address not found in wallet",
	}

	errNotImportedAccount = &vhcjson.RPCError{
		Code:    vhcjson.ErrRPCWallet,
		Message: "imported addresses must belong to the imported account",
	}

	errNeedPositiveAmount = &vhcjson.RPCError{
		Code:    vhcjson.ErrRPCInvalidParameter,
		Message: "amount must be positive",
	}

	errWalletUnlockNeeded = &vhcjson.RPCError{
		Code:    vhcjson.ErrRPCWalletUnlockNeeded,
		Message: "enter the wallet passphrase with walletpassphrase first",
	}

	errReservedAccountName = &vhcjson.RPCError{
		Code:    vhcjson.ErrRPCInvalidParameter,
		Message: "account name is reserved by RPC server",
	}
)
