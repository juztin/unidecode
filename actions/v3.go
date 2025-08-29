package actions

import (
	"encoding/json"
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

func (m V3Mint) MarshalJSON() ([]byte, error) {
	type Alias V3Mint
	return json.Marshal(&struct {
		Alias
		Type string `json:"type"`
	}{(Alias)(m), V3_MINT.String()})
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
	TokenID        *big.Int  `json:"tokenId"`
	Amount0Desired *big.Int  `json:"amount0desired"`
	Amount1Desired *big.Int  `json:"amount1desired"`
	Amount0Min     *big.Int  `json:"amount0min"`
	Amount1Min     *big.Int  `json:"amount1min"`
	Deadline       time.Time `json:"deadline"`
}

func (V3IncreaseLiquidity) Type() Type {
	return V3_INCREASE_LIQUIDITY
}

func (l V3IncreaseLiquidity) MarshalJSON() ([]byte, error) {
	type Alias V3IncreaseLiquidity
	return json.Marshal(&struct {
		Alias
		Type string `json:"type"`
	}{(Alias)(l), V3_INCREASE_LIQUIDITY.String()})
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
	TokenID    *big.Int  `json:"tokenId"`
	Liquidity  *big.Int  `json:"liquidity"`
	Amount0Min *big.Int  `json:"amount0min"`
	Amount1Min *big.Int  `json:"amount1min"`
	Deadline   time.Time `json:"deadline"`
}

func (V3DecreaseLiquidity) Type() Type {
	return V3_DECREASE_LIQUIDITY
}

func (l V3DecreaseLiquidity) MarshalJSON() ([]byte, error) {
	type Alias V3DecreaseLiquidity
	return json.Marshal(&struct {
		Alias
		Type string `json:"type"`
	}{(Alias)(l), V3_DECREASE_LIQUIDITY.String()})
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
	TokenID    *big.Int       `json:"tokenId"`
	Recipient  common.Address `json:"recipient"`
	Amount0Max *big.Int       `json:"amount0max"`
	Amount1Max *big.Int       `json:"amount1max"`
}

func (V3Collect) Type() Type {
	return V3_COLLECT
}

func (c V3Collect) MarshalJSON() ([]byte, error) {
	type Alias V3Collect
	return json.Marshal(&struct {
		Alias
		Type string `json:"type"`
	}{(Alias)(c), V3_COLLECT.String()})
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
	TokenID *big.Int `json:"tokenId"`
}

func (V3Burn) Type() Type {
	return V3_BURN
}

func (c V3Burn) MarshalJSON() ([]byte, error) {
	type Alias V3Burn
	return json.Marshal(&struct {
		Alias
		Type string `json:"type"`
	}{(Alias)(c), V3_BURN.String()})
}

func DecodeV3Burn(calldata []byte, offset int) (V3Burn, error) {
	a := V3Burn{
		TokenID: new(big.Int).SetBytes(calldata[offset : offset+0x20]),
	}
	return a, nil
}
