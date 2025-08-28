package commands

import (
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
	Recipient    common.Address
	AmountIn     *big.Int
	AmountOutMin *big.Int
	// payer = PayerIsUser ? msg.sender : this
	PayerIsUser bool
	// TODO Implement v2 path decoding
	// Path ...
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
	Recipient   common.Address
	AmountOut   *big.Int
	AmountInMin *big.Int
	// payer = PayerIsUser ? msg.sender : this
	PayerIsUser bool
	// TODO Implement v2 path decoding
	// Path ...
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
