package router

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/juztin/unidecode/path"
	"github.com/juztin/unidecode/pool"
)

type ExactInputSingleParams struct {
	PoolKey          pool.Key
	ZeroForOne       bool
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
	HookData         []byte
}

type ExactInputParams struct {
	CurrencyIn       common.Address
	PathKey          []path.Key
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
}

type ExactOutputSingleParams struct {
	PoolKey         pool.Key
	ZeroForOne      bool
	AmountOut       *big.Int
	AmountInMaximum *big.Int
	HookData        []byte
}

type ExactOutputParams struct {
	CurrencyOut     common.Address
	Path            []path.Key
	AmountOut       *big.Int
	AmountInMaximum *big.Int
}
