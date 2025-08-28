package commands

import (
	"bytes"
	"fmt"
	"time"

	"github.com/juztin/unidecode/actions"
	"github.com/juztin/unidecode/hex"
)

// modifyLiquidities(bytes,uint256) | keccak256
var modifyLiquiditiesSig = []byte{0xdd, 0x46, 0x50, 0x8f}

type V4PositionManagerCall struct {
	Deadline time.Time
	actions  []actions.Action
}

func (V4PositionManagerCall) Type() Type {
	return V4_POSITION_MANAGER_CALL
}

func (p V4PositionManagerCall) Actions() []actions.Action {
	return p.actions
}

func DecodeV4PositionManagerCall(calldata []byte, offset int) (V4PositionManagerCall, error) {
	var p V4PositionManagerCall
	length, err := hex.Int(calldata[offset : offset+0x20])
	if err != nil {
		return p, err
	} else if length > len(calldata) {
		return p, fmt.Errorf("invalid %s data length", V4_POSITION_MANAGER_CALL)
	}
	offset += 0x20

	methodSig := calldata[offset : offset+0x04]
	// if (selector != V4_POSITION_MANAGER.modifyLiquidities.selector) {  // universal-router/contracts/modules/V3ToV4Migrator.sol:79
	if bytes.Compare(methodSig, modifyLiquiditiesSig) != 0 {
		return p, fmt.Errorf("invalid function selector; expected %x but got %x", modifyLiquiditiesSig, methodSig)
	}
	offset += 0x04

	// ModifyLiquidities(bytes calldata unlockData, uint256 deadline)
	unlockDataStart, err := hex.Int(calldata[offset : offset+0x20])
	//   [arg0] unlockData
	deadline, err := hex.Int64(calldata[offset+0x20 : offset+0x40])
	if err != nil {
		return p, fmt.Errorf("invalid deadline; %w", err)
	}
	//   [arg1] deadline
	p.Deadline = time.Unix(deadline, 0)
	// Advance offset to start of unlockData bytes
	offset += unlockDataStart

	dataLen, err := hex.Int(calldata[offset : offset+0x20])
	if err != nil {
		return p, fmt.Errorf("invalid %s input count; %w", V4_POSITION_MANAGER_CALL, err)
	} else if dataLen > len(calldata)-offset-(0x40-0x04) { // 0x40-0x4 (first 2 blocks and method sig bytes)
		return p, fmt.Errorf("invalid %s data length", V4_POSITION_MANAGER_CALL)
	}
	offset += 0x20

	// Get action start location
	actionStart, err := hex.Int(calldata[offset : offset+0x20])
	if err != nil {
		return p, fmt.Errorf("invalid action start memory location; %w", err)
	}

	actionLen, err := hex.Int(calldata[offset+actionStart : offset+actionStart+0x20])
	if err != nil {
		return p, fmt.Errorf("invalid action length value; %w", err)
	}

	offset = offset + actionStart + 0x20
	actionTypes, err := actions.DecodeType(calldata[offset : offset+actionLen])
	if err != nil {
		//return p, fmt.Errorf("invalid actions %x; %w", calldata[offset:offset+actionLen], err)
		return p, err
	}

	if actionLen > 0x20 {
		// TODO
	} else {
		offset += 0x20
	}

	paramsLen, err := hex.Int(calldata[offset : offset+0x20])
	if err != nil {
		return p, fmt.Errorf("invalid params length; %w", err)
	} else if paramsLen != actionLen {
		return p, fmt.Errorf("params length mismatch; got %d but expected %d", paramsLen, actionLen)
	}
	offset += 0x20

	var paramsLoc []int
	for i := 0; i < paramsLen; i++ {
		paramOffset := offset + i*0x20
		loc, err := hex.Int(calldata[paramOffset : paramOffset+0x20])
		if err != nil {
			return p, fmt.Errorf("invalid param location for param %d; %w", i, err)
		}
		paramsLoc = append(paramsLoc, loc)
	}

	for i, t := range actionTypes {
		a, err := actions.Decode(t, calldata, offset+paramsLoc[i])
		if err != nil {
			return p, fmt.Errorf("invalid action data at index %d for %s; %w", i, t, err)
		}
		p.actions = append(p.actions, a)
	}

	return p, nil
}
