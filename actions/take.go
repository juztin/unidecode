package actions

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/juztin/unidecode/hex"
)

type Take struct {
	Currency  common.Address `json:"currency"`
	Recipient common.Address `json:"recipient"`
	Amount    *big.Int       `json:"amount"`
}

func (Take) Type() Type {
	return TAKE
}

func (t Take) MarshalJSON() ([]byte, error) {
	type Alias Take
	return json.Marshal(&struct {
		Alias
		Type string `json:"type"`
	}{(Alias)(t), TAKE.String()})
}

func DecodeTake(calldata []byte, offset int) (Take, error) {
	var t Take
	count, err := hex.Int(calldata[offset : offset+0x20])
	if err != nil {
		return t, fmt.Errorf("invalid %s input count; %w", err)
	} else if count != 0x60 {
		return t, fmt.Errorf("invalid %s input count; expected 0x60 but got 0x%x", TAKE, count)
	}

	t = Take{
		Currency:  common.BytesToAddress(calldata[offset+0x20 : offset+0x40]),
		Recipient: common.BytesToAddress(calldata[offset+0x40 : offset+0x60]),
		Amount:    new(big.Int).SetBytes(calldata[offset+0x60 : offset+0x80]),
	}
	return t, nil
}

type TakeAll struct {
	Currency  common.Address `json:"currency"`
	MinAmount *big.Int       `json:"minAmount"`
}

func (t TakeAll) Type() Type {
	return TAKE_ALL
}

func (t TakeAll) MarshalJSON() ([]byte, error) {
	type Alias TakeAll
	return json.Marshal(&struct {
		Alias
		Type string `json:"type"`
	}{(Alias)(t), TAKE_ALL.String()})
}

func DecodeTakeAll(calldata []byte, offset int) (TakeAll, error) {
	var t TakeAll
	count, err := hex.Int(calldata[offset : offset+0x20])
	if err != nil {
		return t, fmt.Errorf("invalid %s input count; %w", err)
	} else if count != 0x40 {
		return t, fmt.Errorf("invalid %s input count; expected 0x40 but got 0x%x", TAKE, count)
	}

	t = TakeAll{
		Currency:  common.BytesToAddress(calldata[offset+0x20 : offset+0x40]),
		MinAmount: new(big.Int).SetBytes(calldata[offset+0x40 : offset+0x60]),
	}
	return t, nil
}

type TakePortion struct {
	Currency  common.Address `json:"currency"`
	Recipient common.Address `json:"recipient"`
	BIPs      *big.Int       `json:"bips"`
}

func (t TakePortion) Type() Type {
	return TAKE_PORTION
}

func (t TakePortion) MarshalJSON() ([]byte, error) {
	type Alias TakePortion
	return json.Marshal(&struct {
		Alias
		Type string `json:"type"`
	}{(Alias)(t), TAKE_PORTION.String()})
}

func DecodeTakePortion(calldata []byte, offset int) (TakePortion, error) {
	var t TakePortion
	count, err := hex.Int(calldata[offset : offset+0x20])
	if err != nil {
		return t, fmt.Errorf("invalid %s input count; %w", err)
	} else if count != 0x60 {
		return t, fmt.Errorf("invalid %s input count; expected 0x60 but got 0x%x", TAKE, count)
	}

	t = TakePortion{
		Currency:  common.BytesToAddress(calldata[offset+0x20 : offset+0x40]),
		Recipient: common.BytesToAddress(calldata[offset+0x40 : offset+0x60]),
		BIPs:      new(big.Int).SetBytes(calldata[offset+0x60 : offset+0x80]),
	}
	return t, nil
}
