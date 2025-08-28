package commands

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/juztin/unidecode/actions"
	"github.com/juztin/unidecode/hex"
)

type Transfer struct {
	Token     common.Address
	Recipient common.Address
	Value     *big.Int
}

func (Transfer) Type() Type {
	return TRANSFER
}

func (Transfer) Actions() []actions.Action {
	return nil
}

func DecodeTransfer(calldata []byte, offset int) (Transfer, error) {
	var t Transfer
	count, err := hex.Int(calldata[offset : offset+0x20])
	if err != nil {
		return t, fmt.Errorf("invalid %s input count; %w", err)
	} else if count != 0x60 {
		return t, fmt.Errorf("invalid %s input count; expected 0x60 but got 0x%x", TRANSFER, count)
	}

	t = Transfer{
		Token:     common.BytesToAddress(calldata[offset+0x20 : offset+0x40]),
		Recipient: common.BytesToAddress(calldata[offset+0x40 : offset+0x60]),
		Value:     new(big.Int).SetBytes(calldata[offset+0x60 : offset+0x80]),
	}
	return t, nil
}
