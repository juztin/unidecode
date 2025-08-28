package actions

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/juztin/unidecode/hex"
)

type V3Mint struct {
	Token0         common.Address
	Token1         common.Address
	Fee            uint
	TickLower      int
	TickUpper      int
	Amount0Desired *big.Int
	Amount1Desired *big.Int
	Amount0Min     *big.Int
	Amount1Min     *big.Int
	Recipient      common.Address
	Deadline       time.Time
}

func (V3Mint) Type() Type {
	return V3_MINT
}

func DecodeV3Mint(calldata []byte, offset int) (V3Mint, error) {
	a := V3Mint{
		Token0:         common.BytesToAddress(calldata[offset : offset+0x20]),
		Token1:         common.BytesToAddress(calldata[offset+0x20 : offset+0x40]),
		Fee:            uint(new(big.Int).SetBytes(calldata[offset+0x40 : offset+0x60]).Uint64()),
		TickLower:      int(new(big.Int).SetBytes(calldata[offset+0x60 : offset+0x80]).Int64()),
		TickUpper:      int(new(big.Int).SetBytes(calldata[offset+0x80 : offset+0xa0]).Int64()),
		Amount0Desired: new(big.Int).SetBytes(calldata[offset+0xa0 : offset+0xc0]),
		Amount1Desired: new(big.Int).SetBytes(calldata[offset+0xc0 : offset+0xe0]),
		Amount0Min:     new(big.Int).SetBytes(calldata[offset+0xe0 : offset+0x100]),
		Amount1Min:     new(big.Int).SetBytes(calldata[offset+0x100 : offset+0x120]),
		Recipient:      common.BytesToAddress(calldata[offset+0x120 : offset+0x140]),
	}
	deadline, err := hex.Int64(calldata[offset+0x140 : offset+0x160])
	if err != nil {
		return a, fmt.Errorf("invalid deadline; %w", err)
	}
	a.Deadline = time.Unix(deadline, 0)
	return a, nil
}

type V3IncreaseLiquidity struct {
	TokenID        *big.Int
	Amount0Desired *big.Int
	Amount1Desired *big.Int
	Amount0Min     *big.Int
	Amount1Min     *big.Int
	Deadline       time.Time
}

func (V3IncreaseLiquidity) Type() Type {
	return V3_INCREASE_LIQUIDITY
}

func DecodeV3IncreaseLiquidity(calldata []byte, offset int) (V3IncreaseLiquidity, error) {
	a := V3IncreaseLiquidity{
		TokenID:        new(big.Int).SetBytes(calldata[offset : offset+0x20]),
		Amount0Desired: new(big.Int).SetBytes(calldata[offset+0x20 : offset+0x40]),
		Amount1Desired: new(big.Int).SetBytes(calldata[offset+0x40 : offset+0x60]),
		Amount0Min:     new(big.Int).SetBytes(calldata[offset+0x60 : offset+0x80]),
		Amount1Min:     new(big.Int).SetBytes(calldata[offset+0x80 : offset+0xa0]),
	}
	deadline, err := hex.Int64(calldata[offset+0x140 : offset+0x160])
	if err != nil {
		return a, fmt.Errorf("invalid deadline; %w", err)
	}
	a.Deadline = time.Unix(deadline, 0)
	return a, nil
}

type V3DecreaseLiquidity struct {
	TokenID    *big.Int
	Liquidity  *big.Int
	Amount0Min *big.Int
	Amount1Min *big.Int
	Deadline   time.Time
}

func (V3DecreaseLiquidity) Type() Type {
	return V3_DECREASE_LIQUIDITY
}

func DecodeV3DecreaseLiquidity(calldata []byte, offset int) (V3DecreaseLiquidity, error) {
	a := V3DecreaseLiquidity{
		TokenID:    new(big.Int).SetBytes(calldata[offset : offset+0x20]),
		Liquidity:  new(big.Int).SetBytes(calldata[offset+0x20 : offset+0x40]),
		Amount0Min: new(big.Int).SetBytes(calldata[offset+0x40 : offset+0x60]),
		Amount1Min: new(big.Int).SetBytes(calldata[offset+0x60 : offset+0x80]),
	}
	return a, nil
}

type V3Collect struct {
	TokenID    *big.Int
	Recipient  common.Address
	Amount0Max *big.Int
	Amount1Max *big.Int
}

func (V3Collect) Type() Type {
	return V3_COLLECT
}

func DecodeV3Collect(calldata []byte, offset int) (V3Collect, error) {
	a := V3Collect{
		TokenID:    new(big.Int).SetBytes(calldata[offset : offset+0x20]),
		Recipient:  common.BytesToAddress(calldata[offset+0x20 : offset+0x40]),
		Amount0Max: new(big.Int).SetBytes(calldata[offset+0x40 : offset+0x60]),
		Amount1Max: new(big.Int).SetBytes(calldata[offset+0x60 : offset+0x80]),
	}
	return a, nil
}

type V3Burn struct {
	TokenID *big.Int
}

func (V3Burn) Type() Type {
	return V3_BURN
}

func DecodeV3Burn(calldata []byte, offset int) (V3Burn, error) {
	a := V3Burn{
		TokenID: new(big.Int).SetBytes(calldata[offset : offset+0x20]),
	}
	return a, nil
}
