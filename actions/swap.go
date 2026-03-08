package actions

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/juztin/unidecode/hex"
	"github.com/juztin/unidecode/path"
	"github.com/juztin/unidecode/pool"
)

// SwapExactOut Solidity representation
//
// see: v4-periphery/src/interfaces/IV4Router.sol:44
type SwapExactOut struct {
	CurrencyOut     common.Address `json:"currencyOut"`
	Path            []path.Key     `json:"path"`
	AmountOut       *big.Int       `json:"amountOut"`
	AmountInMaximum *big.Int       `json:"amountInMaximum"`
}

func (SwapExactOut) Type() Type {
	return SWAP_EXACT_OUT
}

func (s SwapExactOut) MarshalJSON() ([]byte, error) {
	type Alias SwapExactOut
	return json.Marshal(&struct {
		Alias
		Type string `json:"type"`
	}{(Alias)(s), SWAP_EXACT_OUT.String()})
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

// SwapExactOutSingle Solidity representation
//
// v4-periphery/src/interfaces/IV4Router.sol:35
type SwapExactOutSingle struct {
	PoolKey         pool.Key `json:"poolKey"`
	ZeroForOne      bool     `json:"zeroForOne"`
	AmountOut       *big.Int `json:"amountOut"`
	AmountInMaximum *big.Int `json:"amountInMaximum"`
	HookData        []byte   `json:"hookData"`
}

func (s SwapExactOutSingle) Type() Type {
	return SWAP_EXACT_OUT_SINGLE
}

func (s SwapExactOutSingle) MarshalJSON() ([]byte, error) {
	type Alias SwapExactOutSingle
	return json.Marshal(&struct {
		Alias
		Type string `json:"type"`
	}{(Alias)(s), SWAP_EXACT_OUT_SINGLE.String()})
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
		return s, fmt.Errorf("invalid start value: 0x%x; %w", calldata[offset:offset+0x20], err)
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
	s.HookData, err = hookDataFrom(calldata, offset)
	return s, err
}

// SwapExactIn Solidity representation
//
// v4-periphery/src/interfaces/IV4Router.sol:27
type SwapExactIn struct {
	CurrencyIn       common.Address `json:"currencyIn"`
	Path             []path.Key     `json:"path"`
	AmountIn         *big.Int       `json:"amountIn"`
	AmountOutMinimum *big.Int       `json:"amountOutMinimum"`
}

func (SwapExactIn) Type() Type {
	return SWAP_EXACT_IN
}

func (s SwapExactIn) MarshalJSON() ([]byte, error) {
	type Alias SwapExactIn
	return json.Marshal(&struct {
		Alias
		Type string `json:"type"`
	}{(Alias)(s), SWAP_EXACT_IN.String()})
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
		return s, fmt.Errorf("invalid start value: 0x%x; %w", calldata[offset:offset+0x20], err)
	}
	offset += start

	s = SwapExactIn{}
	s.CurrencyIn = common.BytesToAddress(calldata[offset : offset+0x20])
	s.AmountIn = new(big.Int).SetBytes(calldata[offset+0x40 : offset+0x60])
	s.AmountOutMinimum = new(big.Int).SetBytes(calldata[offset+0x60 : offset+0x80])

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

// SwapExactInSingle Solidity representation
//
// v4-periphery/src/interfaces/IV4Router.sol:18
type SwapExactInSingle struct {
	PoolKey          pool.Key `json:"poolKey"`
	ZeroForOne       bool     `json:"zeroForOne"`
	AmountIn         *big.Int `json:"amountInt"`
	AmountOutMinimum *big.Int `json:"amountOutMinimum"`
	HookData         []byte   `json:"hookData"`
}

func (SwapExactInSingle) Type() Type {
	return SWAP_EXACT_IN_SINGLE
}

func (s SwapExactInSingle) MarshalJSON() ([]byte, error) {
	type Alias SwapExactInSingle
	return json.Marshal(&struct {
		Alias
		HookData string `json:"hookData"`
	}{(Alias)(s), fmt.Sprintf("0x%x", s.HookData)})
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
		return s, fmt.Errorf("invalid start value: 0x%x; %w", calldata[offset:offset+0x20], err)
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
	s.AmountOutMinimum = new(big.Int).SetBytes(calldata[offset+0xe0 : offset+0x100])
	s.HookData, err = hookDataFrom(calldata, offset)
	return s, err
}

func hookDataFrom(calldata []byte, offset int) ([]byte, error) {
	hookDataOffset, err := hex.Int(calldata[offset+0x100 : offset+0x120])
	if err != nil {
		return nil, fmt.Errorf("invalid hook-data start value: 0x%x", calldata[offset+0x100:offset+0x120])
	}
	hookDataLen, err := hex.Int(calldata[offset+hookDataOffset : offset+hookDataOffset+0x20])
	var hookData []byte
	if err != nil {
		err = fmt.Errorf("invalid hook-data length: 0x%x; %w", calldata[offset+hookDataOffset:offset+hookDataOffset+0x20], err)
	} else if hookDataLen == 0 {
		hookData = []byte{0x0}
	} else {
		hookData = calldata[offset+hookDataOffset+0x20 : offset+hookDataOffset+0x20+hookDataLen]
	}
	return hookData, nil
}
