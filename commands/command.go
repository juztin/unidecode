package commands

import (
	"errors"

	"github.com/juztin/unidecode/actions"
)

type Command interface {
	Type() Type
	Actions() []actions.Action
}

var errNotImplemented = errors.New("Not implemented")

func Decode(t Type, calldata []byte, offset int) (Command, error) {
	switch t {
	case V3_SWAP_EXACT_IN:
		return DecodeV3SwapExactIn(calldata, offset)
	case V3_SWAP_EXACT_OUT:
		return DecodeV3SwapExactOut(calldata, offset)
	case PERMIT2_TRANSFER_FROM:
		return DecodePermit2TransferFrom(calldata, offset)
	case PERMIT2_PERMIT_BATCH:
		return DecodePermit2PermitBatch(calldata, offset)
	case SWEEP:
		return DecodeSweep(calldata, offset)
	case TRANSFER:
		return DecodeTransfer(calldata, offset)
	case PAY_PORTION:
		return DecodePayPortion(calldata, offset)
	case V2_SWAP_EXACT_IN:
		return DecodeV2SwapExactIn(calldata, offset)
	case V2_SWAP_EXACT_OUT:
		return DecodeV2SwapExactOut(calldata, offset)
	case PERMIT2_PERMIT:
		return DecodePermit2Permit(calldata, offset)
	case WRAP_ETH:
		return DecodeWrapWETH(calldata, offset)
	case UNWRAP_WETH:
		return DecodeUnwrapWETH(calldata, offset)
	case PERMIT2_TRANSFER_FROM_BATCH:
		return DecodePermit2TransferFromBatch(calldata, offset)
	case BALANCE_CHECK_ERC20:
		return DecodeBalanceCheckERC20(calldata, offset)
	case V3_POSITION_MANAGER_PERMIT:
		return DecodeV3PositionManagerPermit(calldata, offset)
	case V3_POSITION_MANAGER_CALL:
		return DecodeV3PositionManagerCall(calldata, offset)
	case V4_SWAP:
		return DecodeV4Swap(calldata, offset)
	case V4_INITIALIZE_POOL:
		return DecodeV4InitializePool(calldata, offset)
	case V4_POSITION_MANAGER_CALL:
		return DecodeV4PositionManagerCall(calldata, offset)
	// NOT IMPLEMENTED
	case EXECUTE_SUB_PLAN,
		FLAG_ALLOW_REVERT,
		COMMAND_TYPE_MASK:
		return nil, errNotImplemented
	}
	return nil, errInvalidType
}
