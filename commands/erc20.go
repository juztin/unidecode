package commands

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/juztin/unidecode/actions"
)

type BalanceCheckERC20 struct {
	Owner      common.Address `json:"owner"`
	Token      common.Address `json:"token"`
	MinBalance *big.Int       `json:"minBalance"`
}

func (b BalanceCheckERC20) MarshalJSON() ([]byte, error) {
	type Alias BalanceCheckERC20
	return json.Marshal(&struct {
		Alias
		Type string `json:"type"`
	}{(Alias)(b), BALANCE_CHECK_ERC20.String()})
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
