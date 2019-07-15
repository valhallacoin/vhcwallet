rpcclient
=========

[![Build Status](http://img.shields.io/travis/valhallacoin/vhcd.svg)](https://travis-ci.org/valhallacoin/vhcd)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/valhallacoin/vhcd/rpcclient)

rpcclient implements a Websocket-enabled Valhalla JSON-RPC client package written
in [Go](http://golang.org/).  It provides a robust and easy to use client for
interfacing with a Valhalla RPC server that uses a vhcd compatible Valhalla
JSON-RPC API.

## Status

This package is currently under active development.  It is already stable and
the infrastructure is complete.  However, there are still several RPCs left to
implement and the API is not stable yet.

## Documentation

* [API Reference](http://godoc.org/github.com/valhallacoin/vhcd/rpcclient)
* [vhcd Websockets Example](https://github.com/valhallacoin/vhcd/tree/master/rpcclient/examples/vhcdwebsockets)
  Connects to a vhcd RPC server using TLS-secured websockets, registers for
  block connected and block disconnected notifications, and gets the current
  block count
* [vhcwallet Websockets Example](https://github.com/valhallacoin/vhcd/tree/master/rpcclient/examples/vhcwalletwebsockets)  
  Connects to a vhcwallet RPC server using TLS-secured websockets, registers for
  notifications about changes to account balances, and gets a list of unspent
  transaction outputs (utxos) the wallet can sign

## Major Features

* Supports Websockets (vhcd/vhcwallet) and HTTP POST mode (bitcoin core-like)
* Provides callback and registration functions for vhcd/vhcwallet notifications
* Supports vhcd extensions
* Translates to and from higher-level and easier to use Go types
* Offers a synchronous (blocking) and asynchronous API
* When running in Websockets mode (the default):
  * Automatic reconnect handling (can be disabled)
  * Outstanding commands are automatically reissued
  * Registered notifications are automatically reregistered
  * Back-off support on reconnect attempts

## Installation

```bash
$ go get -u github.com/valhallacoin/vhcd/rpcclient
```

## License

Package rpcclient is licensed under the [copyfree](http://copyfree.org) ISC
License.
