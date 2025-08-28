package actions

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/juztin/unidecode/hex"
	"github.com/juztin/unidecode/pool"
)

type MintPosition struct {
	PoolKey    pool.Key
	TickLower  int64
	TickUpper  int64
	Liquidity  *big.Int
	Amount0Max *big.Int
	Amount1Max *big.Int
	Owner      common.Address
	HookData   []byte
}

func (MintPosition) Type() Type {
	return MINT_POSITION
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
