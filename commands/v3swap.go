package commands

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/juztin/unidecode/actions"
)

type V3SwapExactIn struct {
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
	// TODO Implement v3 path decoding
	// Path ...
}

func (V3SwapExactIn) Type() Type {
	return V3_SWAP_EXACT_IN
}

func (V3SwapExactIn) Actions() []actions.Action {
	return nil
}

func DecodeV3SwapExactIn(calldata []byte, offset int) (V3SwapExactIn, error) {
	//count, err := hex.Int(calldata[offset : offset+0x20])
	//if err != nil {
	//}
	offset = offset + 0x20

	s := V3SwapExactIn{
		Recipient:    common.BytesToAddress(calldata[offset : offset+0x20]),
		AmountIn:     new(big.Int).SetBytes(calldata[offset+0x20 : offset+0x40]),
		AmountOutMin: new(big.Int).SetBytes(calldata[offset+0x40 : offset+0x60]),
	}

	// TODO Implement v3 path decoding
	//pathCount, err := hex.Int(calldata[offset+0x60: offset+0x80])
	//if err != nil {
	//}
	//pathItems, err := hex.Int(calldata[offset+0x80: offset+0xa0])
	//if err != nil {
	//}

	return s, nil
}

type V3SwapExactOut struct {
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
	// TODO Implement v3 path decoding
	// Path ...
}

func (V3SwapExactOut) Type() Type {
	return V3_SWAP_EXACT_OUT
}

func (V3SwapExactOut) Actions() []actions.Action {
	return nil
}

func DecodeV3SwapExactOut(calldata []byte, offset int) (V3SwapExactOut, error) {
	//count, err := hex.Int(calldata[offset : offset+0x20])
	//if err != nil {
	//}
	offset = offset + 0x20

	s := V3SwapExactOut{
		Recipient:   common.BytesToAddress(calldata[offset : offset+0x20]),
		AmountOut:   new(big.Int).SetBytes(calldata[offset+0x20 : offset+0x40]),
		AmountInMin: new(big.Int).SetBytes(calldata[offset+0x40 : offset+0x60]),
	}
	return s, nil
}
