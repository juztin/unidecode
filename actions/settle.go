package actions

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/juztin/unidecode/hex"
)

type Settle struct {
	Currency    common.Address
	Amount      *big.Int
	PayerIsUser bool
}

func (Settle) Type() Type {
	return SETTLE
}

func DecodeSettle(calldata []byte, offset int) (Settle, error) {
	var s Settle
	count, err := hex.Int(calldata[offset : offset+0x20])
	if err != nil {
		return s, fmt.Errorf("invalid %s input count; %w", err)
	} else if count != 0x60 {
		return s, fmt.Errorf("invalid %s input count; expected 0x60 but got 0x%x", SETTLE, count)
	}

	payerIsUser, err := hex.Bool(calldata[offset+0x60 : offset+0x80])
	if err != nil {
		return s, fmt.Errorf("invalid payerIsUser value; %w", err)
	}

	s = Settle{
		Currency:    common.BytesToAddress(calldata[offset+0x20 : offset+0x40]),
		Amount:      new(big.Int).SetBytes(calldata[offset+0x40 : offset+0x60]),
		PayerIsUser: payerIsUser,
	}
	return s, nil
}

type SettleAll struct {
	Currency  common.Address
	MaxAmount *big.Int
}

func (s SettleAll) Type() Type {
	return SETTLE_ALL
}

func DecodeSettleAll(calldata []byte, offset int) (SettleAll, error) {
	var s SettleAll
	count, err := hex.Int(calldata[offset : offset+0x20])
	if err != nil {
		return s, fmt.Errorf("invalid %s input count; %w", err)
	} else if count != 0x40 {
		return s, fmt.Errorf("invalid %s input count; expected 0x40 but got 0x%x", SETTLE, count)
	}

	s = SettleAll{
		Currency:  common.BytesToAddress(calldata[offset+0x20 : offset+0x40]),
		MaxAmount: new(big.Int).SetBytes(calldata[offset+0x40 : offset+0x60]),
	}
	return s, nil
}
