package commands

import (
	"encoding/json"
	"fmt"

	"github.com/juztin/unidecode/actions"
	"github.com/juztin/unidecode/hex"
)

type V4Swap struct {
	actions []actions.Action
}

func (V4Swap) Type() Type {
	return V4_SWAP
}

func (s V4Swap) MarshalJSON() ([]byte, error) {
	type Alias V4Swap
	return json.Marshal(&struct {
		Alias
		Type    string           `json:"type"`
		Actions []actions.Action `json:"actions"`
	}{(Alias)(s), V4_SWAP.String(), s.actions})
}

func (s V4Swap) Actions() []actions.Action {
	return s.actions
}

func DecodeV4Swap(calldata []byte, offset int) (V4Swap, error) {
	var s V4Swap

	dataLen, err := hex.Int(calldata[offset : offset+0x20])
	if err != nil {
		return s, fmt.Errorf("invalid %s input count; %w", V4_SWAP, err)
	} else if dataLen > len(calldata)-offset {
		return s, fmt.Errorf("invalid %s data length", V4_SWAP)
	}
	offset += 0x20

	// Get action start location
	actionStart, err := hex.Int(calldata[offset : offset+0x20])
	if err != nil {
		return s, fmt.Errorf("invalid action start memory location; %w", err)
	}

	actionLen, err := hex.Int(calldata[offset+actionStart : offset+actionStart+0x20])
	if err != nil {
		return s, fmt.Errorf("invalid action length value; %w", err)
	}

	offset = offset + actionStart + 0x20
	actionTypes, err := actions.DecodeType(calldata[offset : offset+actionLen])
	if err != nil {
		return s, fmt.Errorf("invalid actions; %w", err)
	}

	if actionLen > 0x20 {
		// TODO
	} else {
		offset += 0x20
	}

	paramsLen, err := hex.Int(calldata[offset : offset+0x20])
	if err != nil {
		return s, fmt.Errorf("invalid params length; %w", err)
	} else if paramsLen != actionLen {
		return s, fmt.Errorf("params length mismatch; got %d but expected %d", paramsLen, actionLen)
	}
	offset += 0x20

	var paramsLoc []int
	for i := 0; i < paramsLen; i++ {
		paramOffset := offset + i*0x20
		loc, err := hex.Int(calldata[paramOffset : paramOffset+0x20])
		if err != nil {
			return s, fmt.Errorf("invalid param location for param %d; %w", i, err)
		}
		paramsLoc = append(paramsLoc, loc)
	}

	s = V4Swap{}
	for i, t := range actionTypes {
		a, err := actions.Decode(t, calldata, offset+paramsLoc[i])
		if err != nil {
			return s, fmt.Errorf("invalid action data at index %d for %s; %w", i, t, err)
		}
		s.actions = append(s.actions, a)
	}

	return s, nil
}
