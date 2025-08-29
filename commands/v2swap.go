package commands

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/juztin/unidecode/actions"
)

type V2SwapExactIn struct {
	// NOTE:
	// Recipient:
	//       1: msg.sender
	//       2: this
	//   other: value
	Recipient    common.Address `json:"recipient"`
	AmountIn     *big.Int       `json:"amountIn"`
	AmountOutMin *big.Int       `json:"amountOutMin"`
	PayerIsUser  bool           `json:"payerIsUser"` // payer = PayerIsUser ? msg.sender : this
	// TODO Implement v2 path decoding
	// Path ...
}

func (s V2SwapExactIn) MarshalJSON() ([]byte, error) {
	type Alias V2SwapExactIn
	return json.Marshal(&struct {
		Alias
		Type string `json:"type"`
	}{(Alias)(s), V2_SWAP_EXACT_IN.String()})
}

func (V2SwapExactIn) Type() Type {
	return V2_SWAP_EXACT_IN
}

func (V2SwapExactIn) Actions() []actions.Action {
	return nil
}

func DecodeV2SwapExactIn(calldata []byte, offset int) (V2SwapExactIn, error) {
	//count, err := hex.Int(calldata[offset : offset+0x20])
	//if err != nil {
	//}
	offset = offset + 0x20

	s := V2SwapExactIn{
		Recipient:    common.BytesToAddress(calldata[offset : offset+0x20]),
		AmountIn:     new(big.Int).SetBytes(calldata[offset+0x20 : offset+0x40]),
		AmountOutMin: new(big.Int).SetBytes(calldata[offset+0x40 : offset+0x60]),
	}

	// TODO Implement v2 path decoding

	return s, nil
}

type V2SwapExactOut struct {
	// NOTE:
	// Recipient:
	//       1: msg.sender
	//       2: this
	//   other: value
	Recipient   common.Address `json:"recipient"`
	AmountOut   *big.Int       `json:"amountOut"`
	AmountInMin *big.Int       `json:"amountInMin"`
	PayerIsUser bool           `json:"payerIsUser"` // payer = PayerIsUser ? msg.sender : this
	// TODO Implement v2 path decoding
	// Path ...
}

func (s V2SwapExactOut) MarshalJSON() ([]byte, error) {
	type Alias V2SwapExactOut
	return json.Marshal(&struct {
		Alias
		Type string `json:"type"`
	}{(Alias)(s), V2_SWAP_EXACT_OUT.String()})
}

func (V2SwapExactOut) Type() Type {
	return V2_SWAP_EXACT_OUT
}

func (V2SwapExactOut) Actions() []actions.Action {
	return nil
}

func DecodeV2SwapExactOut(calldata []byte, offset int) (V2SwapExactOut, error) {
	//count, err := hex.Int(calldata[offset : offset+0x20])
	//if err != nil {
	//}
	offset = offset + 0x20

	s := V2SwapExactOut{
		Recipient:   common.BytesToAddress(calldata[offset : offset+0x20]),
		AmountOut:   new(big.Int).SetBytes(calldata[offset+0x20 : offset+0x40]),
		AmountInMin: new(big.Int).SetBytes(calldata[offset+0x40 : offset+0x60]),
	}
	return s, nil
}
