package commands

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/juztin/unidecode/actions"
)

type BalanceCheckERC20 struct {
	Owner      common.Address
	Token      common.Address
	MinBalance *big.Int
}

func (BalanceCheckERC20) Type() Type {
	return BALANCE_CHECK_ERC20
}

func (BalanceCheckERC20) Actions() []actions.Action {
	return nil
}

func DecodeBalanceCheckERC20(calldata []byte, offset int) (BalanceCheckERC20, error) {
	return BalanceCheckERC20{
		Owner:      common.BytesToAddress(calldata[offset+0x20 : offset+0x40]),
		Token:      common.BytesToAddress(calldata[offset+0x40 : offset+0x60]),
		MinBalance: new(big.Int).SetBytes(calldata[offset+0x60 : offset+0x80]),
	}, nil
}
