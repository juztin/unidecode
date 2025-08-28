package actions

import (
	"errors"
	"fmt"
)

type Type int

const (
	INCREASE_LIQUIDITY             Type = 0x00
	DECREASE_LIQUIDITY             Type = 0x01
	MINT_POSITION                  Type = 0x02
	BURN_POSITION                  Type = 0x03
	INCREASE_LIQUIDITY_FROM_DELTAS Type = 0x04
	MINT_POSITION_FROM_DELTAS      Type = 0x05
	SWAP_EXACT_IN_SINGLE           Type = 0x06
	SWAP_EXACT_IN                  Type = 0x07
	SWAP_EXACT_OUT_SINGLE          Type = 0x08
	SWAP_EXACT_OUT                 Type = 0x09
	DONATE                         Type = 0x0a
	SETTLE                         Type = 0x0b
	SETTLE_ALL                     Type = 0x0c
	SETTLE_PAIR                    Type = 0x0d
	TAKE                           Type = 0x0e
	TAKE_ALL                       Type = 0x0f
	TAKE_PORTION                   Type = 0x10
	TAKE_PAIR                      Type = 0x11
	CLOSE_CURRENCY                 Type = 0x12
	CLEAR_OR_TAKE                  Type = 0x13
	SWEEP                          Type = 0x14
	WRAP                           Type = 0x15
	UNWRAP                         Type = 0x16
	MINT_6909                      Type = 0x17
	BURN_6909                      Type = 0x18

	// - V3 --------------
	// Handlers found within `v3-periphery/contracts/NonfungiblePositionManager.sol`
	V3_MINT               Type = 0x88316456 // mint((address,address,uint24,int24,int24,uint256,uint256,uint256,uint256,address,uint256))
	V3_INCREASE_LIQUIDITY Type = 0x219f5d17 // increaseLiquidity((uint256,uint256,uint256,uint256,uint256,uint256))
	V3_DECREASE_LIQUIDITY Type = 0x0c49ccbe // decreaseLiquidity((uint256,uint128,uint256,uint256,uint256))
	V3_COLLECT            Type = 0xfc6f7865 // collect((uint256,address,uint128,uint128))
	V3_BURN               Type = 0x42966c68 // burn(uint256)
)

var errInvalidType = errors.New("invalid action type")

func (a Type) String() string {
	switch a {
	case 0x00:
		return "INCREASE_LIQUIDITY"
	case 0x01:
		return "DECREASE_LIQUIDITY"
	case 0x02:
		return "MINT_POSITION"
	case 0x03:
		return "BURN_POSITION"
	case 0x04:
		return "INCREASE_LIQUIDITY_FROM_DELTAS"
	case 0x05:
		return "MINT_POSITION_FROM_DELTAS"
	// swapping
	case 0x06:
		return "SWAP_EXACT_IN_SINGLE"
	case 0x07:
		return "SWAP_EXACT_IN"
	case 0x08:
		return "SWAP_EXACT_OUT_SINGLE"
	case 0x09:
		return "SWAP_EXACT_OUT"
	// donate
	// note this is not supported in the position manager or router
	case 0x0a:
		return "DONATE"
	// closing deltas on the pool manager
	// settling
	case 0x0b:
		return "SETTLE"
	case 0x0c:
		return "SETTLE_ALL"
	case 0x0d:
		return "SETTLE_PAIR"
	// taking
	case 0x0e:
		return "TAKE"
	case 0x0f:
		return "TAKE_ALL"
	case 0x10:
		return "TAKE_PORTION"
	case 0x11:
		return "TAKE_PAIR"
	case 0x12:
		return "CLOSE_CURRENCY"
	case 0x13:
		return "CLEAR_OR_TAKE"
	case 0x14:
		return "SWEEP"
	case 0x15:
		return "WRAP"
	case 0x16:
		return "UNWRAP"
	// minting/burning 6909s to close deltas
	// note this is not supported in the position manager or router
	case 0x17:
		return "MINT_6909"
	case 0x18:
		return "BURN_6909"

	// - V3 --------------
	case 0x88316456:
		return "V3_MINT"
	case 0x219f5d17:
		return "V3_INCREASE_LIQUIDITY"
	case 0x0c49ccbe:
		return "V3_DECREASE_LIQUIDITY"
	case 0xfc6f7865:
		return "V3_COLLECT"
	case 0x42966c68:
		return "V3_BURN"
	}
	return "UNKNOWN"
}

func Parse(b byte) (t Type, err error) {
	switch b {
	case 0x00:
		t = INCREASE_LIQUIDITY
	case 0x01:
		t = DECREASE_LIQUIDITY
	case 0x02:
		t = MINT_POSITION
	case 0x03:
		t = BURN_POSITION
	case 0x04:
		t = INCREASE_LIQUIDITY_FROM_DELTAS
	case 0x05:
		t = MINT_POSITION_FROM_DELTAS
	case 0x06:
		t = SWAP_EXACT_IN_SINGLE
	case 0x07:
		t = SWAP_EXACT_IN
	case 0x08:
		t = SWAP_EXACT_OUT_SINGLE
	case 0x09:
		t = SWAP_EXACT_OUT
	case 0x0a:
		t = DONATE
	case 0x0b:
		t = SETTLE
	case 0x0c:
		t = SETTLE_ALL
	case 0x0d:
		t = SETTLE_PAIR
	case 0x0e:
		t = TAKE
	case 0x0f:
		t = TAKE_ALL
	case 0x10:
		t = TAKE_PORTION
	case 0x11:
		t = TAKE_PAIR
	case 0x12:
		t = CLOSE_CURRENCY
	case 0x13:
		t = CLEAR_OR_TAKE
	case 0x14:
		t = SWEEP
	case 0x15:
		t = WRAP
	case 0x16:
		t = UNWRAP
	case 0x17:
		t = MINT_6909
	case 0x18:
		t = BURN_6909
	default:
		//err = errInvalidType
		err = fmt.Errorf("invalid action type 0x%x", b)
	}
	return
}

func ParseV3(i int64) (t Type, err error) {
	switch i {
	case 0x88316456:
		t = V3_MINT
	case 0x219f5d17:
		t = V3_INCREASE_LIQUIDITY
	case 0x0c49ccbe:
		t = V3_DECREASE_LIQUIDITY
	case 0xfc6f7865:
		t = V3_COLLECT
	case 0x42966c68:
		t = V3_BURN
	default:
		err = fmt.Errorf("invalid action type 0x%x", i)
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
