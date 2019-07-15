package txrules_test

import (
	"testing"

	"github.com/valhallacoin/vhcd/chaincfg"
	"github.com/valhallacoin/vhcd/vhcutil"
	. "github.com/valhallacoin/vhcwallet/wallet/txrules"
)

func TestStakePoolTicketFee(t *testing.T) {
	params := &chaincfg.MainNetParams
	tests := []struct {
		StakeDiff vhcutil.Amount
		Fee       vhcutil.Amount
		Height    int32
		PoolFee   float64
		Expected  vhcutil.Amount
	}{
		0: {10 * 1e8, 0.01 * 1e8, 25000, 1.00, 0.05861753 * 1e8},
		1: {20 * 1e8, 0.01 * 1e8, 25000, 1.00, 0.08284478 * 1e8},
		2: {5 * 1e8, 0.05 * 1e8, 50000, 2.59, 0.09559588 * 1e8},
		3: {15 * 1e8, 0.05 * 1e8, 50000, 2.59, 0.18520901 * 1e8},
	}
	for i, test := range tests {
		poolFeeAmt := StakePoolTicketFee(test.StakeDiff, test.Fee, test.Height,
			test.PoolFee, params)
		if poolFeeAmt != test.Expected {
			t.Errorf("Test %d: Got %v: Want %v", i, poolFeeAmt, test.Expected)
		}
	}
}
