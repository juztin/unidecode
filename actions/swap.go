package actions

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/juztin/unidecode/hex"
	"github.com/juztin/unidecode/path"
	"github.com/juztin/unidecode/pool"
)

type SwapExactOut struct {
	CurrencyOut     common.Address
	Path            []path.Key
	AmountOut       *big.Int
	AmountInMaximum *big.Int
}

func (SwapExactOut) Type() Type {
	return SWAP_EXACT_OUT
}

func DecodeSwapExactOut(calldata []byte, offset int) (SwapExactOut, error) {
	var s SwapExactOut
	length, _ := hex.Int(calldata[offset : offset+0x20])
	required := len(calldata) - offset + 0x20
	if length > required {
		return s, fmt.Errorf("action data exceeds calldata bounds; expected at-least %d but got %d", length, required)
	}

	offset += 0x20
	start, err := hex.Int(calldata[offset : offset+0x20])
	if err != nil {
		return s, fmt.Errorf("invalid start value; %w", calldata[offset:offset+0x20], err)
	}
	offset += start

	s = SwapExactOut{}
	s.CurrencyOut = common.BytesToAddress(calldata[offset : offset+0x20])
	s.AmountOut = new(big.Int).SetBytes(calldata[offset+0x40 : offset+0x60])
	s.AmountInMaximum = new(big.Int).SetBytes(calldata[offset+0x60 : offset+0x80])

	pathsStart, err := hex.Int(calldata[offset+0x20 : offset+0x40])
	if err != nil {
		return s, fmt.Errorf("invalid path start loc; %w", err)
	}

	s.Path, err = path.DecodeMany(calldata, offset+pathsStart)
	if err != nil {
		return s, fmt.Errorf("failed to decode paths; %w", err)
	}
	return s, nil
}

type SwapExactOutSingle struct {
	PoolKey         pool.Key
	ZeroForOne      bool
	AmountOut       *big.Int
	AmountInMaximum *big.Int
	HookData        []byte
}

func (s SwapExactOutSingle) Type() Type {
	return SWAP_EXACT_OUT_SINGLE
}

func DecodeSwapExactOutSingle(calldata []byte, offset int) (SwapExactOutSingle, error) {
	var s SwapExactOutSingle
	length, _ := hex.Int(calldata[offset : offset+0x20])
	required := len(calldata) - offset + 0x20
	if length > required {
		return s, fmt.Errorf("action data exceeds calldata bounds; expected at-least %d but got %d", length, required)
	}

	offset += 0x20
	start, err := hex.Int(calldata[offset : offset+0x20])
	if err != nil {
		return s, fmt.Errorf("invalid start value; %w", calldata[offset:offset+0x20], err)
	}
	offset += start

	s = SwapExactOutSingle{}
	s.PoolKey = pool.Key{
		Currency0:   common.BytesToAddress(calldata[offset : offset+0x20]),
		Currency1:   common.BytesToAddress(calldata[offset+0x20 : offset+0x40]),
		Fee:         new(big.Int).SetBytes(calldata[offset+0x40 : offset+0x60]),
		TickSpacing: new(big.Int).SetBytes(calldata[offset+0x60 : offset+0x80]),
		Hooks:       common.BytesToAddress(calldata[offset+0x80 : offset+0xa0]),
	}
	s.ZeroForOne, err = hex.Bool(calldata[offset+0xa0 : offset+0xc0])
	if err != nil {
		fmt.Println("invalid zeroForOne value; %w", err)
	}

	s.AmountOut = new(big.Int).SetBytes(calldata[offset+0xc0 : offset+0xe0])
	s.AmountInMaximum = new(big.Int).SetBytes(calldata[offset+0xe0 : offset+0x100])
	// TODO check hookData loc
	s.HookData = calldata[offset+0x100 : offset+0x120]
	return s, nil
}

type SwapExactIn struct {
	CurrencyIn       common.Address
	Path             []path.Key
	AmountIn         *big.Int
	AmountOutMaximum *big.Int
}

func (SwapExactIn) Type() Type {
	return SWAP_EXACT_IN
}

func DecodeSwapExactIn(calldata []byte, offset int) (SwapExactIn, error) {
	var s SwapExactIn
	length, _ := hex.Int(calldata[offset : offset+0x20])
	required := len(calldata) - offset + 0x20
	if length > required {
		return s, fmt.Errorf("action data exceeds calldata bounds; expected at-least %d but got %d", length, required)
	}

	offset += 0x20
	start, err := hex.Int(calldata[offset : offset+0x20])
	if err != nil {
		return s, fmt.Errorf("invalid start value; %w", calldata[offset:offset+0x20], err)
	}
	offset += start

	s = SwapExactIn{}
	s.CurrencyIn = common.BytesToAddress(calldata[offset : offset+0x20])
	s.AmountIn = new(big.Int).SetBytes(calldata[offset+0x40 : offset+0x60])
	s.AmountOutMaximum = new(big.Int).SetBytes(calldata[offset+0x60 : offset+0x80])

	pathsStart, err := hex.Int(calldata[offset+0x20 : offset+0x40])
	if err != nil {
		return s, fmt.Errorf("invalid path start loc; %w", err)
	}

	s.Path, err = path.DecodeMany(calldata, offset+pathsStart)
	if err != nil {
		return s, fmt.Errorf("failed to decode paths; %w", err)
	}
	return s, nil
}

type SwapExactInSingle struct {
	PoolKey          pool.Key
	ZeroForOne       bool
	AmountIn         *big.Int
	AmountOutMaximum *big.Int
	HookData         []byte
}

func (SwapExactInSingle) Type() Type {
	return SWAP_EXACT_IN_SINGLE
}

func DecodeSwapExactInSingle(calldata []byte, offset int) (SwapExactInSingle, error) {
	var s SwapExactInSingle
	length, _ := hex.Int(calldata[offset : offset+0x20])
	required := len(calldata) - offset + 0x20
	if length > required {
		return s, fmt.Errorf("action data exceeds calldata bounds; expected at-least %d but got %d", length, required)
	}

	offset += 0x20
	start, err := hex.Int(calldata[offset : offset+0x20])
	if err != nil {
		return s, fmt.Errorf("invalid start value; %w", calldata[offset:offset+0x20], err)
	}
	offset += start

	s = SwapExactInSingle{}
	s.PoolKey = pool.Key{
		Currency0:   common.BytesToAddress(calldata[offset : offset+0x20]),
		Currency1:   common.BytesToAddress(calldata[offset+0x20 : offset+0x40]),
		Fee:         new(big.Int).SetBytes(calldata[offset+0x40 : offset+0x60]),
		TickSpacing: new(big.Int).SetBytes(calldata[offset+0x60 : offset+0x80]),
		Hooks:       common.BytesToAddress(calldata[offset+0x80 : offset+0xa0]),
	}
	s.ZeroForOne, err = hex.Bool(calldata[offset+0xa0 : offset+0xc0])
	if err != nil {
		fmt.Println("invalid zeroForOne value; %w", err)
	}

	s.AmountIn = new(big.Int).SetBytes(calldata[offset+0xc0 : offset+0xe0])
	s.AmountOutMaximum = new(big.Int).SetBytes(calldata[offset+0xe0 : offset+0x100])
	// TODO check hookData loc
	s.HookData = calldata[offset+0x100 : offset+0x120]
	return s, nil
}
