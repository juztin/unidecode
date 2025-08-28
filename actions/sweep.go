package actions

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/juztin/unidecode/hex"
)

type Sweep struct {
	Currency  common.Address
	Recipient common.Address
}

func (Sweep) Type() Type {
	return SWEEP
}

func DecodeSweep(calldata []byte, offset int) (Sweep, error) {
	var s Sweep
	count, err := hex.Int(calldata[offset : offset+0x20])
	if err != nil {
		return s, fmt.Errorf("invalid %s input count; %w", err)
	} else if count != 0x40 {
		return s, fmt.Errorf("invalid %s input count; expected 0x40 but got 0x%x", SWEEP, count)
	}

	s = Sweep{
		Currency:  common.BytesToAddress(calldata[offset+0x20 : offset+0x40]),
		Recipient: common.BytesToAddress(calldata[offset+0x40 : offset+0x60]),
	}
	return s, nil
}
