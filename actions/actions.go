package actions

import "errors"

type Action interface {
	Type() Type
}

var errUnsupportedAction = errors.New("Unsupported action")

func Decode(t Type, calldata []byte, offset int) (Action, error) {
	switch t {
	case MINT_POSITION:
		return DecodeMintPosition(calldata, offset)
	case SWAP_EXACT_IN_SINGLE:
		return DecodeSwapExactInSingle(calldata, offset)
	case SWAP_EXACT_IN:
		return DecodeSwapExactIn(calldata, offset)
	case SWAP_EXACT_OUT_SINGLE:
		return DecodeSwapExactOutSingle(calldata, offset)
	case SWAP_EXACT_OUT:
		return DecodeSwapExactOut(calldata, offset)
	case SETTLE:
		return DecodeSettle(calldata, offset)
	case SETTLE_ALL:
		return DecodeSettleAll(calldata, offset)
	case TAKE:
		return DecodeTake(calldata, offset)
	case TAKE_ALL:
		return DecodeTakeAll(calldata, offset)
	case TAKE_PORTION:
		return DecodeTakePortion(calldata, offset)
	case SWEEP:
		return DecodeSweep(calldata, offset)
	case WRAP:
		return DecodeWrap(calldata, offset)
	case UNWRAP:
		return DecodeUnwrap(calldata, offset)

	// Below actions are not handled within V4Router.sol within `_handleAction(action, params)`
	case INCREASE_LIQUIDITY,
		DECREASE_LIQUIDITY,
		BURN_POSITION,
		INCREASE_LIQUIDITY_FROM_DELTAS,
		MINT_POSITION_FROM_DELTAS,
		DONATE,
		SETTLE_PAIR,
		TAKE_PAIR,
		CLOSE_CURRENCY,
		CLEAR_OR_TAKE,
		MINT_6909,
		BURN_6909:
		return nil, errUnsupportedAction

	// - V3 --------------------
	case V3_MINT:
		return DecodeV3Mint(calldata, offset)
	case V3_INCREASE_LIQUIDITY:
		return DecodeV3IncreaseLiquidity(calldata, offset)
	case V3_DECREASE_LIQUIDITY:
		return DecodeV3DecreaseLiquidity(calldata, offset)
	case V3_COLLECT:
		return DecodeV3Collect(calldata, offset)
	case V3_BURN:
		return DecodeV3Burn(calldata, offset)
	}
	return nil, errInvalidType
}
