package commands

import (
	"bytes"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/juztin/unidecode/actions"
	"github.com/juztin/unidecode/hex"
)

// permit(address,uint256,uint256,uint8,bytes32,bytes32) | keccak256
var permitSig = []byte{0x7a, 0xc2, 0xff, 0x7b}

type Sig struct {
	V uint8
	//R []byte
	//S []byte
	R [32]byte
	S [32]byte
}

type V3PositionManagerPermit struct {
	Spender  common.Address
	Amount   *big.Int
	Deadline time.Time
	Sig      Sig
}

func (V3PositionManagerPermit) Type() Type {
	return V3_POSITION_MANAGER_PERMIT
}

func (V3PositionManagerPermit) Actions() []actions.Action {
	return nil
}

func DecodeV3PositionManagerPermit(calldata []byte, offset int) (V3PositionManagerPermit, error) {
	var p V3PositionManagerPermit
	length, err := hex.Int(calldata[offset : offset+0x20])
	if err != nil {
		return p, err
	} else if length > len(calldata) {
		return p, fmt.Errorf("invalid %s data length", V4_POSITION_MANAGER_CALL)
	}
	offset += 0x20

	methodSig := calldata[offset : offset+0x04]
	if bytes.Compare(methodSig, permitSig) != 0 {
		return p, fmt.Errorf("invalid function selector; expected %x but got %x", permitSig, methodSig)
	}
	offset += 0x04

	p = V3PositionManagerPermit{
		Spender: common.BytesToAddress(calldata[offset : offset+0x20]),
		Amount:  new(big.Int).SetBytes(calldata[offset+0x20 : offset+0x40]),
	}
	deadline, err := hex.Int64(calldata[offset+0x40 : offset+0x60])
	if err != nil {
		return p, fmt.Errorf("invalid deadline; %w", err)
	}
	p.Deadline = time.Unix(deadline, 0)

	// Grabbing just the last 8 bytes (+24bytes/0x18)
	v, err := hex.Int(calldata[offset+0x60+0x18 : offset+0x80])
	if err != nil {
		return p, fmt.Errorf("invalid V signature value; %w", err)
	}
	p.Sig = Sig{
		V: uint8(v),
		//R: calldata[offset+0xa0 : offset+0xc0],
		//S: calldata[offset+0xc0 : offset+0xe0],
	}
	copy(p.Sig.R[:], calldata[offset+0x80:offset+0xa0])
	copy(p.Sig.S[:], calldata[offset+0xa0:offset+0xc0])

	return p, nil
}

// universal-router/contracts/base/Dispatcher.sol:254

type V3PositionManagerCall struct {
	Action actions.Action
}

func (V3PositionManagerCall) Type() Type {
	return V3_POSITION_MANAGER_CALL
}

func (p V3PositionManagerCall) Actions() []actions.Action {
	return []actions.Action{p.Action}
}

func DecodeV3PositionManagerCall(calldata []byte, offset int) (V3PositionManagerCall, error) {
	var p V3PositionManagerCall
	length, err := hex.Int(calldata[offset : offset+0x20])
	if err != nil {
		return p, err
	} else if length > len(calldata) {
		return p, fmt.Errorf("invalid %s data length", V3_POSITION_MANAGER_CALL)
	}

	methodSig, err := hex.Int64(calldata[offset+0x20 : offset+0x24])
	if err != nil {
		return p, err
	}
	offset += 0x24

	t, err := actions.ParseV3(methodSig)
	if err != nil {
		return p, err
	}
	a, err := actions.Decode(t, calldata, offset)
	if err != nil {
		return p, err
	}
	p.Action = a
	return p, nil
}
