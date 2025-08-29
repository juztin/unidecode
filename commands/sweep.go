package commands

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/juztin/unidecode/actions"
	"github.com/juztin/unidecode/hex"
)

type Sweep struct {
	Token     common.Address `json:"token"`
	Recipient common.Address `json:"recipient"`
	AmountMin *big.Int       `json:"amountMin"`
}

func (s Sweep) MarshalJSON() ([]byte, error) {
	type Alias Sweep
	return json.Marshal(&struct {
		Alias
		Type string `json:"type"`
	}{(Alias)(s), SWEEP.String()})
}

func (Sweep) Type() Type {
	return SWEEP
}

func (Sweep) Actions() []actions.Action {
	return nil
}

func DecodeSweep(calldata []byte, offset int) (Sweep, error) {
	var s Sweep
	count, err := hex.Int(calldata[offset : offset+0x20])
	if err != nil {
		return s, fmt.Errorf("invalid %s input count; %w", err)
	} else if count != 0x60 {
		return s, fmt.Errorf("invalid %s input count; expected 0x60 but got 0x%x", SWEEP, count)
	}

	s = Sweep{
		Token:     common.BytesToAddress(calldata[offset+0x20 : offset+0x40]),
		Recipient: common.BytesToAddress(calldata[offset+0x40 : offset+0x60]),
		AmountMin: new(big.Int).SetBytes(calldata[offset+0x60 : offset+0x80]),
	}
	return s, nil
}
