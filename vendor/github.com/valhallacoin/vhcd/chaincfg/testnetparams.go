// Copyright (c) 2014-2016 The btcsuite developers
// Copyright (c) 2015-2019 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"time"

	"github.com/valhallacoin/vhcd/wire"
)

// TestNetParams defines the network parameters for the test currency network.
// This network is sometimes simply called "testnet".
var TestNetParams = Params{
	Name:        "testnet",
	Net:         wire.TestNet,
	DefaultPort: "19208",
	DNSSeeds: []DNSSeed{
		{"testnet-seed.valhallacoin.org", false},
		{"testnet-seed.valhallacoin.net", false},
		{"testnet-seed.valhalla.cash", false},
		{"d4snyonmvnxyubznqq4e2d3jzge2dphr7n4ikc253yosq7ebilvdzsid.onion", false},
		{"qsell64hbqzldcc2bznxjdx2avsbs7h2xrpy63fnpihlfgoapjyeftyd.onion", false},
	},

	// Chain parameters
	GenesisBlock:             &testNetGenesisBlock,
	GenesisHash:              &testNetGenesisHash,
	PowLimit:                 testNetPowLimit,
	PowLimitBits:             0x1e00ffff,
	ReduceMinDifficulty:      true,
	MinDiffReductionTime:     time.Minute * 10, // ~99.3% chance to be mined before reduction
	GenerateSupported:        true,
	MaximumBlockSizes:        []int{1310720},
	MaxTxSize:                1000000,
	TargetTimePerBlock:       time.Minute * 2,
	WorkDiffAlpha:            1,
	WorkDiffWindowSize:       144,
	WorkDiffWindows:          20,
	TargetTimespan:           time.Minute * 2 * 144, // TimePerBlock * WindowSize
	RetargetAdjustmentFactor: 4,

	// Subsidy parameters.
	BaseSubsidy:              2500000000, // 25 Coin
	MulSubsidy:               100,
	DivSubsidy:               101,
	SubsidyReductionInterval: 2048,
	WorkRewardProportion:     6,
	StakeRewardProportion:    3,
	BlockTaxProportion:       1,

	// Checkpoints ordered from oldest to newest.
	Checkpoints: nil,

	// Consensus rule change deployments.
	//
	// The miner confirmation window is defined as:
	//   target proof of work timespan / target proof of work spacing
	RuleChangeActivationQuorum:     2520, // 10 % of RuleChangeActivationInterval * TicketsPerBlock
	RuleChangeActivationMultiplier: 3,    // 75%
	RuleChangeActivationDivisor:    4,
	RuleChangeActivationInterval:   5040, // 1 week
	Deployments: map[uint32][]ConsensusDeployment{
		7: {{
			Vote: Vote{
				Id:          VoteIDFixLNSeqLocks,
				Description: "Modify sequence lock handling as defined in DCP0004",
				Mask:        0x0006, // Bits 1 and 2
				Choices: []Choice{{
					Id:          "abstain",
					Description: "abstain voting for change",
					Bits:        0x0000,
					IsAbstain:   true,
					IsNo:        false,
				}, {
					Id:          "no",
					Description: "keep the existing consensus rules",
					Bits:        0x0002, // Bit 1
					IsAbstain:   false,
					IsNo:        true,
				}, {
					Id:          "yes",
					Description: "change to the new consensus rules",
					Bits:        0x0004, // Bit 2
					IsAbstain:   false,
					IsNo:        false,
				}},
			},
			StartTime:  1548633600, // Jan 28th, 2019
			ExpireTime: 1580169600, // Jan 28th, 2020
		}},
	},

	// Enforce current block version once majority of the network has
	// upgraded.
	// 51% (51 / 100)
	// Reject previous block versions once a majority of the network has
	// upgraded.
	// 75% (75 / 100)
	BlockEnforceNumRequired: 51,
	BlockRejectNumRequired:  75,
	BlockUpgradeNumToCheck:  100,

	// AcceptNonStdTxs is a mempool param to either accept and relay
	// non standard txs to the network or reject them
	AcceptNonStdTxs: true,

	// Address encoding magics
	NetworkAddressPrefix: "T",
	PubKeyAddrID:         [2]byte{0x28, 0xf7}, // starts with Tk
	PubKeyHashAddrID:     [2]byte{0x0f, 0x21}, // starts with Ts
	PKHEdwardsAddrID:     [2]byte{0x0f, 0x01}, // starts with Te
	PKHSchnorrAddrID:     [2]byte{0x0e, 0xe3}, // starts with TS
	ScriptHashAddrID:     [2]byte{0x0e, 0xfc}, // starts with Tc
	PrivateKeyID:         [2]byte{0x23, 0x0e}, // starts with Pt

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x97}, // starts with tprv
	HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xd1}, // starts with tpub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	SLIP0044CoinType: 1,  // SLIP0044, Testnet (all coins)
	LegacyCoinType:   11, // for backwards compatibility

	// Valhalla PoS parameters
	MinimumStakeDiff:        20000000, // 0.2 Coin
	TicketPoolSize:          1024,
	TicketsPerBlock:         5,
	TicketMaturity:          16,
	TicketExpiry:            6144, // 6*TicketPoolSize
	CoinbaseMaturity:        16,
	SStxChangeMaturity:      1,
	TicketPoolSizeWeight:    4,
	StakeDiffAlpha:          1,
	StakeDiffWindowSize:     144,
	StakeDiffWindows:        20,
	StakeVersionInterval:    144 * 2 * 7, // ~1 week
	MaxFreshStakePerBlock:   20,          // 4*TicketsPerBlock
	StakeEnabledHeight:      16 + 16,     // CoinbaseMaturity + TicketMaturity
	StakeValidationHeight:   768,         // Arbitrary
	StakeBaseSigScript:      []byte{0x00, 0x00},
	StakeMajorityMultiplier: 3,
	StakeMajorityDivisor:    4,

	// Valhalla organization related parameters.
	// Organization address is TsnK11KjNihw9qSSEPDrQiZsedvDKa9VG3u
	OrganizationPkScript:        hexDecode("76a914e98b7a3e26aee2d5dc1112e2e8990617cfdcd06988ac"),
	OrganizationPkScriptVersion: 0,
	BlockOneLedger:              BlockOneLedgerTestNet,
}
