package commands

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/juztin/unidecode/actions"
	"github.com/juztin/unidecode/hex"
)

type PayPortion struct {
	Token     common.Address
	Recipient common.Address
	BIPs      *big.Int
}

func (PayPortion) Type() Type {
	return PAY_PORTION
}

func (PayPortion) Actions() []actions.Action {
	return nil
}

func DecodePayPortion(calldata []byte, offset int) (PayPortion, error) {
	var p PayPortion
	count, err := hex.Int(calldata[offset : offset+0x20])
	if err != nil {
		return p, fmt.Errorf("invalid %s input count; %w", err)
	} else if count != 0x60 {
		return p, fmt.Errorf("invalid %s input count; expected 0x60 but got 0x%x", SWEEP, count)
	}

	p = PayPortion{
		Token:     common.BytesToAddress(calldata[offset+0x20 : offset+0x40]),
		Recipient: common.BytesToAddress(calldata[offset+0x40 : offset+0x60]),
		BIPs:      new(big.Int).SetBytes(calldata[offset+0x60 : offset+0x80]),
	}
	return p, nil
}
