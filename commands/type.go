package commands

import (
	"errors"
	"fmt"
)

type Type int

const (
	FLAG_ALLOW_REVERT           Type = 0x80
	COMMAND_TYPE_MASK           Type = 0x3f
	V3_SWAP_EXACT_IN            Type = 0x00
	V3_SWAP_EXACT_OUT           Type = 0x01
	PERMIT2_TRANSFER_FROM       Type = 0x02
	PERMIT2_PERMIT_BATCH        Type = 0x03
	SWEEP                       Type = 0x04
	TRANSFER                    Type = 0x05
	PAY_PORTION                 Type = 0x06
	V2_SWAP_EXACT_IN            Type = 0x08
	V2_SWAP_EXACT_OUT           Type = 0x09
	PERMIT2_PERMIT              Type = 0x0a
	WRAP_ETH                    Type = 0x0b
	UNWRAP_WETH                 Type = 0x0c
	PERMIT2_TRANSFER_FROM_BATCH Type = 0x0d
	BALANCE_CHECK_ERC20         Type = 0x0e
	V4_SWAP                     Type = 0x10
	V3_POSITION_MANAGER_PERMIT  Type = 0x11
	V3_POSITION_MANAGER_CALL    Type = 0x12
	V4_INITIALIZE_POOL          Type = 0x13
	V4_POSITION_MANAGER_CALL    Type = 0x14
	EXECUTE_SUB_PLAN            Type = 0x21
)

var errInvalidType = errors.New("invalid command type")

func (c Type) String() string {
	switch c {
	case 0x80:
		return "FLAG_ALLOW_REVERT"
	case 0x3f:
		return "COMMAND_TYPE_MASK"
	case 0x00:
		return "V3_SWAP_EXACT_IN"
	case 0x01:
		return "V3_SWAP_EXACT_OUT"
	case 0x02:
		return "PERMIT2_TRANSFER_FROM"
	case 0x03:
		return "PERMIT2_PERMIT_BATCH"
	case 0x04:
		return "SWEEP"
	case 0x05:
		return "TRANSFER"
	case 0x06:
		return "PAY_PORTION"
	case 0x08:
		return "V2_SWAP_EXACT_IN"
	case 0x09:
		return "V2_SWAP_EXACT_OUT"
	case 0x0a:
		return "PERMIT2_PERMIT"
	case 0x0b:
		return "WRAP_ETH"
	case 0x0c:
		return "UNWRAP_WETH"
	case 0x0d:
		return "PERMIT2_TRANSFER_FROM_BATCH"
	case 0x0e:
		return "BALANCE_CHECK_ERC20"
	case 0x10:
		return "V4_SWAP"
	case 0x11:
		return "V3_POSITION_MANAGER_PERMIT"
	case 0x12:
		return "V3_POSITION_MANAGER_CALL"
	case 0x13:
		return "V4_INITIALIZE_POOL"
	case 0x14:
		return "V4_POSITION_MANAGER_CALL"
	case 0x21:
		return "EXECUTE_SUB_PLAN"
	}
	return "UNKNOWN"
}

func Parse(b byte) (t Type, err error) {
	switch b {
	case 0x80:
		t = FLAG_ALLOW_REVERT
	case 0x3f:
		t = COMMAND_TYPE_MASK
	case 0x00:
		t = V3_SWAP_EXACT_IN
	case 0x01:
		t = V3_SWAP_EXACT_OUT
	case 0x02:
		t = PERMIT2_TRANSFER_FROM
	case 0x03:
		t = PERMIT2_PERMIT_BATCH
	case 0x04:
		t = SWEEP
	case 0x05:
		t = TRANSFER
	case 0x06:
		t = PAY_PORTION
	case 0x08:
		t = V2_SWAP_EXACT_IN
	case 0x09:
		t = V2_SWAP_EXACT_OUT
	case 0x0a:
		t = PERMIT2_PERMIT
	case 0x0b:
		t = WRAP_ETH
	case 0x0c:
		t = UNWRAP_WETH
	case 0x0d:
		t = PERMIT2_TRANSFER_FROM_BATCH
	case 0x0e:
		t = BALANCE_CHECK_ERC20
	case 0x10:
		t = V4_SWAP
	case 0x11:
		t = V3_POSITION_MANAGER_PERMIT
	case 0x12:
		t = V3_POSITION_MANAGER_CALL
	case 0x13:
		t = V4_INITIALIZE_POOL
	case 0x14:
		t = V4_POSITION_MANAGER_CALL
	case 0x21:
		t = EXECUTE_SUB_PLAN
	default:
		//err = errInvalidType
		err = fmt.Errorf("invalid command type 0x%x", b)
	}
	return
}

func DecodeType(b []byte) ([]Type, error) {
	var s []Type
	for i := 0; i < len(b); i++ {
		t, err := Parse(b[i])
		if err != nil {
			return s, err
		}
		s = append(s, t)
	}
	return s, nil
}
