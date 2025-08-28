package actions

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/juztin/unidecode/hex"
)

var (
	NATIVE = common.HexToAddress("0000000000000000000000000000000000000000")
	WETH9  = common.HexToAddress("C02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
)

type Wrap struct {
	Amount   *big.Int
	Currency common.Address
	Wrapped  common.Address
}

func (Wrap) Type() Type {
	return WRAP
}

func DecodeWrap(calldata []byte, offset int) (Wrap, error) {
	var w Wrap
	count, err := hex.Int(calldata[offset : offset+0x20])
	if err != nil {
		return w, fmt.Errorf("invalid %s input count; %w", err)
	} else if count != 0x20 {
		return w, fmt.Errorf("invalid %s input count; expected 0x20 but got 0x%x", WRAP, count)
	}

	offset += 0x20
	w = Wrap{
		Amount:   new(big.Int).SetBytes(calldata[offset : offset+0x20]),
		Currency: NATIVE,
		Wrapped:  WETH9,
	}
	return w, nil
}

type Unwrap struct {
	Amount   *big.Int
	Currency common.Address
	Wrapped  common.Address
}

func (Unwrap) Type() Type {
	return UNWRAP
}

func DecodeUnwrap(calldata []byte, offset int) (Unwrap, error) {
	var w Unwrap
	count, err := hex.Int(calldata[offset : offset+0x20])
	if err != nil {
		return w, fmt.Errorf("invalid %s input count; %w", err)
	} else if count != 0x20 {
		return w, fmt.Errorf("invalid %s input count; expected 0x20 but got 0x%x", UNWRAP, count)
	}

	offset += 0x20
	w = Unwrap{
		Amount:   new(big.Int).SetBytes(calldata[offset : offset+0x20]),
		Currency: NATIVE,
		Wrapped:  WETH9,
	}
	return w, nil
}
