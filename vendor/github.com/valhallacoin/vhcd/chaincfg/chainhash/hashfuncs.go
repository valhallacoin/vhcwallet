// Copyright (c) 2015-2016 The Decred developers
// Copyright (c) 2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chainhash

import (
	"github.com/dchest/blake256"
)

// HashFunc calculates the hash of the supplied bytes.
// TODO(jcv) Should modify blake256 so it has the same interface as blake2
// and sha256 so these function can look more like btcsuite.  Then should
// try to get it to the upstream blake256 repo
func HashFunc(b []byte) [blake256.Size]byte {
	var outB [blake256.Size]byte
	copy(outB[:], HashB(b))

	return outB
}

// HashB calculates hash(b) and returns the resulting bytes.
func HashB(b []byte) []byte {
	a := blake256.New()
	a.Write(b)
	out := a.Sum(nil)
	return out
}

// HashH calculates hash(b) and returns the resulting bytes as a Hash.
func HashH(b []byte) Hash {
	return Hash(HashFunc(b))
}

// XORBytes modifies b1 by perfoming a bitwise XOR with the respective
// elements in b2.
func XORBytes(b1 []byte, b2 []byte) {
	if len(b1) != len(b2) {
		panic("chainhash: XORBytes(): lengths don't match")
	}
	for i := 0; i < len(b1); i++ {
		b1[i] = b1[i] ^ b2[i]
	}
}

// PoWHashB calculates (hash(b) ^ MagicBytes) and returns the resulting bytes.
func PoWHashB(b []byte) []byte {
	b1 := HashB(b)
	XORBytes(b1, MagicBytes)
	return b1
}

// PoWHashH calculates (hash(b) ^ MagicBytes) and returns the resulting bytes
// as a Hash.
func PoWHashH(b []byte) Hash {
	var array [HashSize]byte
	copy(array[:], PoWHashB(b))
	return Hash(array)
}

// HashBlockSize is the block size of the hash algorithm in bytes.
const HashBlockSize = blake256.BlockSize

// MagicBytes is the magic 256-bit random number used for Proof-of-Work in Valhalla.
var MagicBytes = []byte{
	0xfa, 0xce, 0x3f, 0xd3, 0xfc, 0x4a, 0x53, 0xe6,
	0xe3, 0x0d, 0x6d, 0x0d, 0xef, 0x6d, 0xb7, 0x3e,
	0xd1, 0xed, 0xac, 0x0d, 0x06, 0x14, 0x4f, 0xfa,
	0x4a, 0x22, 0xa1, 0x73, 0xa1, 0xf6, 0x6c, 0x20,
}
