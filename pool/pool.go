package pool

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Key struct {
	Currency0   common.Address
	Currency1   common.Address
	Fee         *big.Int
	TickSpacing *big.Int
	Hooks       common.Address
}

func NewKey(currency0, currency1 common.Address, fee *big.Int, tickSpacing *big.Int, hooks common.Address) Key {
	if currency0.Cmp(currency1) == 1 {
		currency0, currency1 = currency1, currency0
	}
	return Key{
		Currency0:   currency0,
		Currency1:   currency1,
		Fee:         fee,
		TickSpacing: tickSpacing,
		Hooks:       hooks,
	}
}

func (k Key) ID() common.Hash {
	b := common.LeftPadBytes(k.Currency0[:], 0x20)
	b = append(b, common.LeftPadBytes(k.Currency1[:], 0x20)...)
	b = append(b, common.LeftPadBytes(k.Fee.Bytes(), 0x20)...)
	b = append(b, common.LeftPadBytes(k.TickSpacing.Bytes(), 0x20)...)
	b = append(b, common.LeftPadBytes(k.Hooks[:], 0x20)...)
	return crypto.Keccak256Hash(b)
}
