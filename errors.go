package unidecode

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/rpc"
)

func Error(err error) error {
	if err == nil {
		return nil
	}

	if rpcErr, ok := err.(rpc.DataError); ok {
		if data, ok := rpcErr.ErrorData().(string); ok {
			if len(data) < 10 {
				return fmt.Errorf("Unknown error: %s; %w", data, rpcErr)
			}

			code := data[:10]
			switch code {
			case "0xd0df97cc":
				err = errors.New("ActionNotSupported")
			case "0x01b673a7":
				err = errors.New("AddLiquidityDirectToHook")
			case "0xd81b2f2e":
				err = errors.New("AllowanceExpired")
			case "0x25fbd8be":
				err = errors.New("AlreadySubscribed")
			case "0x5090d6c6":
				err = errors.New("AlreadyUnlocked")
			case "0xa3281672":
				err = errors.New("BalanceTooLow")
			case "0xace94481":
				err = errors.New("BurnNotificationReverted")
			case "0x8c6e5d71":
				err = errors.New("CallerNotWhitelisted")
			case "0xaefeb924":
				err = errors.New("CannotUpdateEmptyPosition")
			case "0x39492f34":
				err = errors.New("CauseRevert")
			case "0xac8429db":
				err = errors.New("CheckParameters")
			case "0x0474c5c1":
				err = errors.New("CompetitionNotOver")
			case "0xafa178b0":
				err = errors.New("CompetitionOver")
			case "0x6f5ffb7e":
				err = errors.New("ContractLocked")
			case "0x6e6c9830":
				err = errors.New("CurrenciesOutOfOrderOrEqual")
			case "0x5212cba1":
				err = errors.New("CurrencyNotSettled")
			case "0x773a6187":
				err = errors.New("DeadlineBeforeEndTime")
			case "0xbfb22adf":
				err = errors.New("DeadlinePassed")
			case "0xb08ce5b3":
				err = errors.New("DeadlineReached")
			case "0x0d89438e":
				err = errors.New("DelegateCallNotAllowed")
			case "0xf217bc40":
				err = errors.New("DeltaNotNegative")
			case "0x5a1aed16":
				err = errors.New("DeltaNotPositive")
			case "0xfff08303":
				err = errors.New("DuplicateFeeOutput")
			case "0x43133453":
				err = errors.New("EndTimeBeforeStartTime")
			case "0xf27f64e4":
				err = errors.New("ERC20TransferFailed")
			case "0xd845fc92":
				err = errors.New("Error1")
			case "0x9407b1cb":
				err = errors.New("Error2")
			case "0xd04addd7":
				err = errors.New("Error36Bytes")
			case "0xa55c4f1f":
				err = errors.New("Error4Bytes")
			case "0x251d6607":
				err = errors.New("Error68Bytes")
			case "0x2dea0607":
				err = errors.New("ErrorBytes")
			case "0x1231ae40":
				err = errors.New("ETHNotAccepted")
			case "0x0ace433b":
				err = errors.New("ExactInputNotSupported")
			case "0x21b865b3":
				err = errors.New("ExactOutputNotSupported")
			case "0x24d35a26":
				err = errors.New("ExcessiveInvalidation")
			case "0x2c4029e9":
				err = errors.New("ExecutionFailed")
			case "0x82e75656":
				err = errors.New("FeeTooLarge")
			case "0xe7002877":
				err = errors.New("FromAddressIsNotOwner")
			case "0xed43c3a6":
				err = errors.New("GasLimitTooLow")
			case "0xe65af6a0":
				err = errors.New("HookAddressNotValid")
			case "0xa9e35b2f":
				err = errors.New("HookCallFailed")
			case "0xfa0b71d6":
				err = errors.New("HookDeltaExceedsSwapAmount")
			case "0x0a85dc29":
				err = errors.New("HookNotImplemented")
			case "0x7c1f8113":
				err = errors.New("IncorrectAmounts")
			case "0x4e23d035":
				err = errors.New("IndexOutOfBounds")
			case "0xd303758b":
				err = errors.New("InputAndOutputDecay")
			case "0xedc7e2e4":
				err = errors.New("InputAndOutputFees")
			case "0xaaad13f7":
				err = errors.New("InputLengthMismatch")
			case "0xa6b844f5":
				err = errors.New("InputOutputScaling")
			case "0xf96fb071":
				err = errors.New("InsufficientAllowance")
			case "0xf4d678b8":
				err = errors.New("InsufficientBalance")
			case "0x6a12f104":
				err = errors.New("InsufficientETH")
			case "0x675cae38":
				err = errors.New("InsufficientToken")
			case "0xf801e525":
				err = errors.New("InvalidAction")
			case "0x6d1eca28":
				err = errors.New("InvalidAddressLength")
			case "0x3728b83d":
				err = errors.New("InvalidAmount")
			case "0x41ca9285":
				err = errors.New("InvalidArrLength")
			case "0xdeaa01e6":
				err = errors.New("InvalidBips")
			case "0x23639643":
				err = errors.New("InvalidBytecode")
			case "0x48f5c3ed":
				err = errors.New("InvalidCaller")
			case "0xd76a1e9e":
				err = errors.New("InvalidCommandType")
			case "0xb0669cbc":
				err = errors.New("InvalidContractSignature")
			case "0xd7815be1":
				err = errors.New("InvalidCosignature")
			case "0xac9143e7":
				err = errors.New("InvalidCosignerInput")
			case "0xa305df82":
				err = errors.New("InvalidCosignerOutput")
			case "0x769d11e4":
				err = errors.New("InvalidDeadline")
			case "0x0e996766":
				err = errors.New("InvalidDecayCurve")
			case "0x38bbd576":
				err = errors.New("InvalidEthSender")
			case "0x96206246":
				err = errors.New("InvalidFeeForExactOut")
			case "0xeddf07f5":
				err = errors.New("InvalidFeeToken")
			case "0xf3eb44e5":
				err = errors.New("InvalidGasPrice")
			case "0x1e048e1d":
				err = errors.New("InvalidHookResponse")
			case "0x756688fe":
				err = errors.New("InvalidNonce")
			case "0x20db8267":
				err = errors.New("InvalidPath")
			case "0x1213a0ab":
				err = errors.New("InvalidPoolFee")
			case "0xdcdedda9":
				err = errors.New("InvalidPoolToken")
			case "0x00bfc921":
				err = errors.New("InvalidPrice")
			case "0x4f2461b8":
				err = errors.New("InvalidPriceOrLiquidity")
			case "0x4ddf4a64":
				err = errors.New("InvalidReactor")
			case "0x7b9c8916":
				err = errors.New("InvalidReserves")
			case "0x8c7d9949":
				err = errors.New("InvalidSender")
			case "0x8baa579f":
				err = errors.New("InvalidSignature")
			case "0x4be6321b":
				err = errors.New("InvalidSignatureLength")
			case "0x815e1d64":
				err = errors.New("InvalidSigner")
			case "0x61487524":
				err = errors.New("InvalidSqrtPrice")
			case "0x8b86327a":
				err = errors.New("InvalidTick")
			case "0xed15e6cf":
				err = errors.New("InvalidTokenId")
			case "0x9096cccb":
				err = errors.New("KeyNotSet")
			case "0xff633a38":
				err = errors.New("LengthMismatch")
			case "0x78895c13":
				err = errors.New("LiquidityNotAllowed")
			case "0x14002113":
				err = errors.New("LPFeeTooLarge")
			case "0x54e3ca0d":
				err = errors.New("ManagerLocked")
			case "0x31e30ad0":
				err = errors.New("MaximumAmountExceeded")
			case "0x12816f22":
				err = errors.New("MinimumAmountInsufficient")
			case "0xb3ca2e28":
				err = errors.New("MockValidationError")
			case "0xe94f10e2":
				err = errors.New("ModifyLiquidityNotificationReverted")
			case "0x933fe52f":
				err = errors.New("MsgSenderNotReactor")
			case "0xbda73abf":
				err = errors.New("MustClearExactPositiveDelta")
			case "0xf4b3b1bc":
				err = errors.New("NativeTransferFailed")
			case "0x7c402b21":
				err = errors.New("NoCodeSubscriber")
			case "0xb9ec1e96":
				err = errors.New("NoExclusiveOverride")
			case "0xa74f97ab":
				err = errors.New("NoLiquidityToReceiveFees")
			case "0x1fb09b80":
				err = errors.New("NonceAlreadyUsed")
			case "0xb0ec849e":
				err = errors.New("NonzeroNativeValue")
			case "0x80e05c00":
				err = errors.New("NoSelfPermit")
			case "0x5a8b05b0":
				err = errors.New("NoSwapOccurred")
			case "0xb418cb98":
				err = errors.New("NotAllowedReenter")
			case "0xc419b2d2":
				err = errors.New("NotAllowedToDeploy")
			case "0x0ca968d8":
				err = errors.New("NotApproved")
			case "0xbb25d4c5":
				err = errors.New("NotAuthorizedForToken")
			case "0x6f05727d":
				err = errors.New("NotAuthorizedNotifer")
			case "0x4323a555":
				err = errors.New("NotEnoughLiquidity")
			case "0x4e38b830":
				err = errors.New("NotEnoughLiquidity")
			case "0x75c1bb14":
				err = errors.New("NotExclusiveFiller")
			case "0xd6234725":
				err = errors.New("NotImplemented")
			case "0xae18210a":
				err = errors.New("NotPoolManager")
			case "0x29c3b7ee":
				err = errors.New("NotSelf")
			case "0x237e6c28":
				err = errors.New("NotSubscribed")
			case "0x5d1d0f9f":
				err = errors.New("OnlyMintAllowed")
			case "0xee3b3d4b":
				err = errors.New("OrderAlreadyFilled")
			case "0xc6035520":
				err = errors.New("OrderNotFillable")
			case "0x06ee9878":
				err = errors.New("OrdersLengthIncorrect")
			case "0xe70ff93c":
				err = errors.New("Permit2NotDeployed")
			case "0x7983c051":
				err = errors.New("PoolAlreadyInitialized")
			case "0xd4b05fe0":
				err = errors.New("PoolManagerMustBeLocked")
			case "0x486aa307":
				err = errors.New("PoolNotInitialized")
			case "0x7c9c6e8f":
				err = errors.New("PriceLimitAlreadyExceeded")
			case "0x9e4d7cc7":
				err = errors.New("PriceLimitOutOfBounds")
			case "0xf5c787f1":
				err = errors.New("PriceOverflow")
			case "0xc79e5948":
				err = errors.New("ProtocolFeeCurrencySynced")
			case "0xa7abe2f7":
				err = errors.New("ProtocolFeeTooLarge")
			case "0xecbd9804":
				err = errors.New("QuoteSwap")
			case "0x93dafdf1":
				err = errors.New("SafeCastOverflow")
			case "0x5a9165ff":
				err = errors.New("SignatureDeadlineExpired")
			case "0xcd21db4f":
				err = errors.New("SignatureExpired")
			case "0x3b99b53d":
				err = errors.New("SliceOutOfBounds")
			case "0x81ea5e9e":
				err = errors.New("SubscriptionReverted")
			case "0xbe8b8507":
				err = errors.New("SwapAmountCannotBeZero")
			case "0x421e7f54":
				err = errors.New("TestRevert")
			case "0xb8e3c385":
				err = errors.New("TickLiquidityOverflow")
			case "0xd5e2f7ab":
				err = errors.New("TickLowerOutOfBounds")
			case "0xd4d8f3e6":
				err = errors.New("TickMisaligned")
			case "0xc4433ed5":
				err = errors.New("TicksMisordered")
			case "0xb70024f8":
				err = errors.New("TickSpacingTooLarge")
			case "0xe9e90588":
				err = errors.New("TickSpacingTooSmall")
			case "0x1ad777f8":
				err = errors.New("TickUpperOutOfBounds")
			case "0x5bf6f916":
				err = errors.New("TransactionDeadlinePassed")
			case "0x82b42900":
				err = errors.New("Unauthorized")
			case "0x30d21641":
				err = errors.New("UnauthorizedDynamicLPFeeUpdate")
			case "0xe0752a5a":
				err = errors.New("UnexpectedCallSuccess")
			case "0x6190b2b0":
				err = errors.New("UnexpectedRevertBytes")
			case "0xc4bd89a9":
				err = errors.New("UnsafeCast")
			case "0x5cda29d7":
				err = errors.New("UnsupportedAction")
			case "0xea3559ef":
				err = errors.New("UnsupportedProtocolError")
			case "0xae52ad0c":
				err = errors.New("V2InvalidPath")
			case "0x849eaf98":
				err = errors.New("V2TooLittleReceived")
			case "0x8ab0bc16":
				err = errors.New("V2TooMuchRequested")
			case "0xd4e0248e":
				err = errors.New("V3InvalidAmountOut")
			case "0x32b13d91":
				err = errors.New("V3InvalidCaller")
			case "0x316cf0eb":
				err = errors.New("V3InvalidSwap")
			case "0x39d35496":
				err = errors.New("V3TooLittleReceived")
			case "0x739dbe52":
				err = errors.New("V3TooMuchRequested")
			case "0x8b063d73":
				err = errors.New("V4TooLittleReceived")
			case "0x12bacdd3":
				err = errors.New("V4TooMuchRequested")
			case "0x29551db8":
				err = errors.New("WorseAddress")
			case "0x90bfb865":
				err = errors.New("WrappedError")
			default:
				err = fmt.Errorf("Unknown error code: %s; %w", code, rpcErr)
			}
		}
	}
	return err
}
