package commands

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/juztin/unidecode/actions"
	"github.com/juztin/unidecode/pool"
)

type V4InitializePool struct {
	Key       pool.Key
	SqrtPrice *big.Int
}

func (V4InitializePool) Type() Type {
	return V4_INITIALIZE_POOL
}

func (V4InitializePool) Actions() []actions.Action {
	return nil
}

func DecodeV4InitializePool(calldata []byte, offset int) (V4InitializePool, error) {
	currency0 := common.BytesToAddress(calldata[offset : offset+0x20])
	currency1 := common.BytesToAddress(calldata[offset+0x20 : offset+0x40])
	fee := new(big.Int).SetBytes(calldata[offset+0x40 : offset+0x60])
	tickSpacing := new(big.Int).SetBytes(calldata[offset+0x60 : offset+0x80])
	hooks := common.BytesToAddress(calldata[offset+0x80 : offset+0xa0])
	sqrtPrice := new(big.Int).SetBytes(calldata[offset+0xa0 : offset+0xc0])

	key := pool.NewKey(currency0, currency1, fee, tickSpacing, hooks)
	return V4InitializePool{
		Key:       key,
		SqrtPrice: sqrtPrice,
	}, nil
}
