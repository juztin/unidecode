package commands

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/juztin/unidecode/actions"
	"github.com/juztin/unidecode/hex"
)

type WrapWETH struct {
	Recipient common.Address `json:"recipient"`
	AmountMin *big.Int       `json:"amountMin"`
}

func (WrapWETH) Type() Type {
	return WRAP_ETH
}

func (w WrapWETH) MarshalJSON() ([]byte, error) {
	type Alias WrapWETH
	return json.Marshal(&struct {
		Alias
		Type string `json:"type"`
	}{(Alias)(w), V4_POSITION_MANAGER_CALL.String()})
}

func (WrapWETH) Actions() []actions.Action {
	return nil
}

func DecodeWrapWETH(calldata []byte, offset int) (WrapWETH, error) {
	var w WrapWETH

	count, err := hex.Int(calldata[offset : offset+0x20])
	if err != nil {
		return w, fmt.Errorf("invalid %s input count; %w", WRAP_ETH, err)
	} else if count != 0x40 {
		return w, fmt.Errorf("invalid %s input count; expected 0x40 but got 0x%x", WRAP_ETH, count)
	}

	w = WrapWETH{
		Recipient: common.BytesToAddress(calldata[offset+0x20 : offset+0x40]),
		AmountMin: new(big.Int).SetBytes(calldata[offset+0x40 : offset+0x60]),
	}
	return w, nil
}

type UnwrapWETH struct {
	Recipient common.Address `json"recipient"`
	AmountMin *big.Int       `json:"amountMin"`
}

func (UnwrapWETH) Type() Type {
	return UNWRAP_WETH
}

func (w UnwrapWETH) MarshalJSON() ([]byte, error) {
	type Alias UnwrapWETH
	return json.Marshal(&struct {
		Alias
		Type string `json:"type"`
	}{(Alias)(w), UNWRAP_WETH.String()})
}

func (UnwrapWETH) Actions() []actions.Action {
	return nil
}

func DecodeUnwrapWETH(calldata []byte, offset int) (UnwrapWETH, error) {
	var u UnwrapWETH

	count, err := hex.Int(calldata[offset : offset+0x20])
	if err != nil {
		return u, fmt.Errorf("invalid %s input count; %w", UNWRAP_WETH, err)
	} else if count != 0x40 {
		return u, fmt.Errorf("invalid %s input count; expected 0x40 but got 0x%x", UNWRAP_WETH, count)
	}

	u = UnwrapWETH{
		Recipient: common.BytesToAddress(calldata[offset+0x20 : offset+0x40]),
		AmountMin: new(big.Int).SetBytes(calldata[offset+0x40 : offset+0x60]),
	}
	return u, nil
}
