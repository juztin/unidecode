package actions

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/juztin/unidecode/hex"
	"github.com/juztin/unidecode/pool"
)

type MintPosition struct {
	PoolKey    pool.Key       `json:"poolKey"`
	TickLower  int64          `json:"tickLower"`
	TickUpper  int64          `json:"tickUpper"`
	Liquidity  *big.Int       `json:"liquidity"`
	Amount0Max *big.Int       `json:"amount0max"`
	Amount1Max *big.Int       `json:"amount1max"`
	Owner      common.Address `json:"owner"`
	HookData   []byte         `json:"hookData"`
}

func (MintPosition) Type() Type {
	return MINT_POSITION
}

func (m MintPosition) MarshalJSON() ([]byte, error) {
	type Alias MintPosition
	return json.Marshal(&struct {
		Alias
		HookData string `json:"hookData"`
		Type     string `json:"type"`
	}{(Alias)(m), fmt.Sprintf("0x%x", m.HookData), MINT_POSITION.String()})
}

func DecodeMintPosition(calldata []byte, offset int) (MintPosition, error) {
	var p MintPosition
	length, _ := hex.Int(calldata[offset : offset+0x20])
	required := len(calldata) - offset + 0x20
	if length > required {
		return p, fmt.Errorf("action data exceeds calldata bounds; expected at-least %d but got %d", length, required)
	}

	offset += 0x20
	hookDataOffset, err := hex.Int(calldata[offset+0x160 : offset+0x180])
	if err != nil {
		return p, fmt.Errorf("invalid hookData offset; %w", err)
	}
	hookDataLen, err := hex.Int(calldata[offset+hookDataOffset : offset+hookDataOffset+0x20])
	if err != nil {
		return p, fmt.Errorf("invalid hookData length; %w", err)
	}

	p = MintPosition{
		PoolKey: pool.Key{
			Currency0:   common.BytesToAddress(calldata[offset : offset+0x20]),
			Currency1:   common.BytesToAddress(calldata[offset+0x20 : offset+0x40]),
			Fee:         new(big.Int).SetBytes(calldata[offset+0x40 : offset+0x60]),
			TickSpacing: new(big.Int).SetBytes(calldata[offset+0x60 : offset+0x80]),
			Hooks:       common.BytesToAddress(calldata[offset+0xa0 : offset+0xa0]),
		},
		// Tick values appear to have a FF mask applied to the first 29bytes of the block.
		// Both values are int24 in size, so we only want the last 3 bytes—exclude the first 29bytes(0x1d)
		TickLower:  new(big.Int).SetBytes(calldata[offset+0xa0+0x1d : offset+0xc0]).Int64(),
		TickUpper:  new(big.Int).SetBytes(calldata[offset+0xc0+0x1d : offset+0xe0]).Int64(),
		Liquidity:  new(big.Int).SetBytes(calldata[offset+0xe0 : offset+0x100]),
		Amount0Max: new(big.Int).SetBytes(calldata[offset+0x100 : offset+0x120]),
		Amount1Max: new(big.Int).SetBytes(calldata[offset+0x120 : offset+0x140]),
		Owner:      common.BytesToAddress(calldata[offset+0x140 : offset+0x160]),
		HookData:   calldata[offset+hookDataOffset : offset+hookDataOffset+hookDataLen],
	}
	return p, nil
}
