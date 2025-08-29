package main

import (
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/juztin/unidecode"
	"github.com/juztin/unidecode/actions"
	"github.com/juztin/unidecode/commands"
)

var (
	actionArgFmt       = "%9s %s: %s\n"
	actionArgSubFmt    = "%11s %s: %s\n"
	actionArgSubIdxFmt = "%11s [%d] %s\n"
	actionArgSubSubFmt = "%13s %s: %s\n"
)

func checkErr(errFmt string, err error) {
	if err != nil {
		if errFmt != "" {
			fmt.Fprintf(os.Stderr, errFmt, err)
			fmt.Fprintln(os.Stderr)
		} else {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}
}

func calldataFromArg(arg string) ([]byte, error) {
	if len(arg) == 0 {
		return nil, fmt.Errorf("empty/missing CALLDATA")
	}
	if strings.HasPrefix(arg, "0x") {
		arg = os.Args[1][2:]
	}
	b, err := hex.DecodeString(arg)
	return b, err
}

func calldataCmd_JSON(pretty bool, args ...string) {
	b, err := calldataFromArg(args[0])
	checkErr("", err)

	t := unidecode.MessageType(b)
	if t != unidecode.ExecuteMessage {
		checkErr("currently only supports EXECUTE messages; %s", fmt.Errorf("got %x but expected %s", b, unidecode.ExecuteMessage))
	}
	execute, err := unidecode.DecodeExecute(b)
	checkErr("", err)

	if pretty {
		b, err = json.MarshalIndent(execute, "", "  ")
	} else {
		b, err = json.Marshal(execute)
	}
	checkErr("", err)

	fmt.Println(string(b))
}

func printActions(cmd commands.Command) {
	for _, action := range cmd.Actions() {
		actionType := action.Type()
		fmt.Fprintf(os.Stdout, "      - %s:\n", strings.ToUpper(actionType.String()))

		switch actionType {
		case actions.MINT_POSITION:
			a := action.(actions.MintPosition)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "PoolKey", a.PoolKey.ID())
			fmt.Fprintf(os.Stdout, actionArgSubFmt, "", "Currency0", a.PoolKey.Currency0)
			fmt.Fprintf(os.Stdout, actionArgSubFmt, "", "Currency1", a.PoolKey.Currency1)
			fmt.Fprintf(os.Stdout, actionArgSubFmt, "", "Fee", a.PoolKey.Fee)
			fmt.Fprintf(os.Stdout, actionArgSubFmt, "", "TickSpacing", a.PoolKey.TickSpacing)
			fmt.Fprintf(os.Stdout, actionArgSubFmt, "", "Hooks", a.PoolKey.Hooks)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "TickLower", a.TickLower)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "TickUpper", a.TickUpper)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Liquidity", a.Liquidity)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Amount0Max", a.Amount0Max)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Amount1Max", a.Amount1Max)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Owner", a.Owner)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "HookData", fmt.Sprintf("%x", a.HookData))
		case actions.SWAP_EXACT_IN_SINGLE:
			a := action.(actions.SwapExactInSingle)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "PoolKey", a.PoolKey.ID())
			fmt.Fprintf(os.Stdout, actionArgSubFmt, "", "Currency0", a.PoolKey.Currency0)
			fmt.Fprintf(os.Stdout, actionArgSubFmt, "", "Currency1", a.PoolKey.Currency1)
			fmt.Fprintf(os.Stdout, actionArgSubFmt, "", "Fee", a.PoolKey.Fee)
			fmt.Fprintf(os.Stdout, actionArgSubFmt, "", "TickSpacing", a.PoolKey.TickSpacing)
			fmt.Fprintf(os.Stdout, actionArgSubFmt, "", "Hooks", a.PoolKey.Hooks)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "ZeroForOne", strconv.FormatBool(a.ZeroForOne))
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "AmountIn", a.AmountIn)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "AmountOutMaximum", a.AmountOutMaximum)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "HookData", fmt.Sprintf("%x", a.HookData))
		case actions.SWAP_EXACT_IN:
			a := action.(actions.SwapExactIn)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "CurrencyIn", a.CurrencyIn)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Path", "")
			for i, p := range a.Path {
				fmt.Fprintf(os.Stdout, actionArgSubIdxFmt, "", i, "")
				fmt.Fprintf(os.Stdout, actionArgSubSubFmt, "", "IntermediateCurrency", p.IntermediateCurrency)
				fmt.Fprintf(os.Stdout, actionArgSubSubFmt, "", "Fee", p.Fee)
				fmt.Fprintf(os.Stdout, actionArgSubSubFmt, "", "TickSpacing", p.TickSpacing)
				fmt.Fprintf(os.Stdout, actionArgSubSubFmt, "", "Hooks", p.Hooks)
				fmt.Fprintf(os.Stdout, actionArgSubSubFmt, "", "HookData", fmt.Sprintf("%x", p.HookData))
			}
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "AmountIn", a.AmountIn)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "AmountOutMaximum", a.AmountOutMaximum)
		case actions.SWAP_EXACT_OUT_SINGLE:
			a := action.(actions.SwapExactOutSingle)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "PoolKey", a.PoolKey.ID())
			fmt.Fprintf(os.Stdout, actionArgSubFmt, "", "Currency0", a.PoolKey.Currency0)
			fmt.Fprintf(os.Stdout, actionArgSubFmt, "", "Currency1", a.PoolKey.Currency1)
			fmt.Fprintf(os.Stdout, actionArgSubFmt, "", "Fee", a.PoolKey.Fee)
			fmt.Fprintf(os.Stdout, actionArgSubFmt, "", "TickSpacing", a.PoolKey.TickSpacing)
			fmt.Fprintf(os.Stdout, actionArgSubFmt, "", "Hooks", a.PoolKey.Hooks)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "ZeroForOne", strconv.FormatBool(a.ZeroForOne))
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "AmountOut", a.AmountOut)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "AmountInMaximum", a.AmountInMaximum)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "HookData", fmt.Sprintf("%x", a.HookData))
		case actions.SWAP_EXACT_OUT:
			a := action.(actions.SwapExactOut)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "CurrencyOut", a.CurrencyOut)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Path", "")
			for i, p := range a.Path {
				fmt.Fprintf(os.Stdout, actionArgSubIdxFmt, "", i, "")
				fmt.Fprintf(os.Stdout, actionArgSubSubFmt, "", "IntermediateCurrency", p.IntermediateCurrency)
				fmt.Fprintf(os.Stdout, actionArgSubSubFmt, "", "Fee", p.Fee)
				fmt.Fprintf(os.Stdout, actionArgSubSubFmt, "", "TickSpacing", p.TickSpacing)
				fmt.Fprintf(os.Stdout, actionArgSubSubFmt, "", "Hooks", p.Hooks)
				fmt.Fprintf(os.Stdout, actionArgSubSubFmt, "", "HookData", fmt.Sprintf("%x", p.HookData))
			}
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "AmountOut", a.AmountOut)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "AmountInMaximum", a.AmountInMaximum)
		case actions.SETTLE:
			a := action.(actions.Settle)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Currency", a.Currency)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Amount", a.Amount)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "PayerIsUser", strconv.FormatBool(a.PayerIsUser))
		case actions.SETTLE_ALL:
			a := action.(actions.SettleAll)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Currency", a.Currency)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "MaxAmount", a.MaxAmount)
		case actions.TAKE:
			a := action.(actions.Take)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Currency", a.Currency)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Recipient", a.Recipient)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Amount", a.Amount)
		case actions.TAKE_ALL:
			a := action.(actions.TakeAll)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Currency", a.Currency)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "MinAmount", a.MinAmount)
		case actions.TAKE_PORTION:
			a := action.(actions.TakePortion)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Currency", a.Currency)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Recipient", a.Recipient)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "BIPs", a.BIPs)
		case actions.SWEEP:
			a := action.(actions.Sweep)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Currency", a.Currency)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Recipient", a.Recipient)
		case actions.WRAP:
			a := action.(actions.Wrap)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Amount", a.Amount)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Currency", a.Currency)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Wrapped", a.Wrapped)
		case actions.UNWRAP:
			a := action.(actions.Unwrap)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Amount", a.Amount)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Currency", a.Currency)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Wrapped", a.Wrapped)
		case actions.V3_MINT:
			a := action.(actions.V3Mint)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Token0", a.Token0)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Token1", a.Token1)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Fee", a.Fee)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "TickLower", a.TickLower)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "TickUpper", a.TickUpper)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Amount0Desired", a.Amount0Desired)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Amount1Desired", a.Amount1Desired)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Amount0Min", a.Amount0Min)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Amount1Min", a.Amount1Min)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Recipient", a.Recipient)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Deadline", a.Deadline)
		case actions.V3_INCREASE_LIQUIDITY:
			a := action.(actions.V3IncreaseLiquidity)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "TokenID", a.TokenID)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Amount0Desired", a.Amount0Desired)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Amount1Desired", a.Amount1Desired)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Amount0Min", a.Amount0Min)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Amount1Min", a.Amount1Min)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Deadline", a.Deadline)
		case actions.V3_DECREASE_LIQUIDITY:
			a := action.(actions.V3DecreaseLiquidity)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "TokenID", a.TokenID)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Liquidity", a.Liquidity)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Amount0Min", a.Amount0Min)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Amount1Min", a.Amount1Min)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Deadline", a.Deadline)
		case actions.V3_COLLECT:
			a := action.(actions.V3Collect)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "TokenID", a.TokenID)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Recipient", a.Recipient)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Amount0Max", a.Amount0Max)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Amount1Max", a.Amount1Max)
		case actions.V3_BURN:
			a := action.(actions.V3Burn)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "TokenID", a.TokenID)

		// Not Implemented
		case actions.INCREASE_LIQUIDITY,
			actions.DECREASE_LIQUIDITY,
			actions.BURN_POSITION,
			actions.INCREASE_LIQUIDITY_FROM_DELTAS,
			actions.MINT_POSITION_FROM_DELTAS,
			actions.DONATE,
			actions.SETTLE_PAIR,
			actions.TAKE_PAIR,
			actions.CLOSE_CURRENCY,
			actions.CLEAR_OR_TAKE,
			actions.MINT_6909,
			actions.BURN_6909:
		}
	}
}

func calldataCmd(args ...string) {
	b, err := calldataFromArg(args[0])
	checkErr("", err)

	t := unidecode.MessageType(b)
	if t != unidecode.ExecuteMessage {
		checkErr("currently only supports EXECUTE messages; %s", fmt.Errorf("got %x but expected %s", b, unidecode.ExecuteMessage))
	}

	execute, err := unidecode.DecodeExecute(b)
	checkErr("", err)

	fmt.Fprintf(os.Stdout, "EXECUTE:\n  Deadline: %s\n", execute.Deadline)
	for _, cmd := range execute.Commands {
		cmdType := cmd.Type()
		fmt.Fprintf(os.Stdout, "    - %s:\n", strings.ToUpper(cmdType.String()))

		switch cmdType {
		case commands.V4_SWAP:
			swap := cmd.(commands.V4Swap)
			printActions(swap)
		case commands.V3_SWAP_EXACT_IN:
			c := cmd.(commands.V3SwapExactIn)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Recipient", c.Recipient)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "AmountIn", c.AmountIn)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "AmountOutMin", c.AmountOutMin)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "PayerIsUser", c.PayerIsUser)
		case commands.V3_SWAP_EXACT_OUT:
			c := cmd.(commands.V3SwapExactOut)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Recipient", c.Recipient)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "AmountOut", c.AmountOut)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "AmountInMin", c.AmountInMin)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "PayerIsUser", c.PayerIsUser)
		case commands.PERMIT2_TRANSFER_FROM:
			c := cmd.(commands.Permit2TransferFrom)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Token", c.Token)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Recipient", c.Recipient)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Amount", c.Amount)
		case commands.PERMIT2_PERMIT_BATCH:
			c := cmd.(commands.Permit2PermitBatch)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Details", "")
			for i, p := range c.Details {
				fmt.Fprintf(os.Stdout, actionArgSubIdxFmt, "", i, "")
				fmt.Fprintf(os.Stdout, actionArgSubSubFmt, "", "Token", p.Token)
				fmt.Fprintf(os.Stdout, actionArgSubSubFmt, "", "Amount", p.Amount)
				fmt.Fprintf(os.Stdout, actionArgSubSubFmt, "", "Expiration", p.Expiration)
				fmt.Fprintf(os.Stdout, actionArgSubSubFmt, "", "Nonce", fmt.Sprintf("%d", p.Nonce))
			}
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Spender", c.Spender)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "SigDeadline", c.SigDeadline)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Sig", fmt.Sprintf("%x", c.Sig))
		case commands.SWEEP:
			c := cmd.(commands.Sweep)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Token", c.Token)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Recipient", c.Recipient)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "AmountMin", c.AmountMin)
		case commands.TRANSFER:
			c := cmd.(commands.Transfer)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Token", c.Token)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Recipient", c.Recipient)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Value", c.Value)
		case commands.PAY_PORTION:
			c := cmd.(commands.PayPortion)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Token", c.Token)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Recipient", c.Recipient)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "BIPs", c.BIPs)
		case commands.V2_SWAP_EXACT_IN:
			c := cmd.(commands.V2SwapExactIn)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Recipient", c.Recipient)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "AmountIn", c.AmountIn)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "AmountOutMin", c.AmountOutMin)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "PayerIsUser", strconv.FormatBool(c.PayerIsUser))
		case commands.V2_SWAP_EXACT_OUT:
			c := cmd.(commands.V2SwapExactOut)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Recipient", c.Recipient)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "AmountOut", c.AmountOut)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "AmountInMin", c.AmountInMin)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "PayerIsUser", strconv.FormatBool(c.PayerIsUser))
		case commands.PERMIT2_PERMIT:
			c := cmd.(commands.Permit2Permit)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Details", "")
			fmt.Fprintf(os.Stdout, actionArgSubSubFmt, "", "Token", c.Details.Token)
			fmt.Fprintf(os.Stdout, actionArgSubSubFmt, "", "Amount", c.Details.Amount)
			fmt.Fprintf(os.Stdout, actionArgSubSubFmt, "", "Expiration", c.Details.Expiration)
			fmt.Fprintf(os.Stdout, actionArgSubSubFmt, "", "Nonce", fmt.Sprintf("%d", c.Details.Nonce))
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Spender", c.Spender)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "SigDeadline", c.SigDeadline)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Sig", fmt.Sprintf("%x", c.Sig))
		case commands.WRAP_ETH:
			c := cmd.(commands.WrapWETH)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Recipient", c.Recipient)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "AmountMin", c.AmountMin)
		case commands.UNWRAP_WETH:
			c := cmd.(commands.UnwrapWETH)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Recipient", c.Recipient)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "AmountMin", c.AmountMin)
		case commands.PERMIT2_TRANSFER_FROM_BATCH:
			// TODO Implement
		case commands.BALANCE_CHECK_ERC20:
			c := cmd.(commands.BalanceCheckERC20)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Owner", c.Owner)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Token", c.Token)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "MinBalance", c.MinBalance)
		case commands.V3_POSITION_MANAGER_PERMIT:
			c := cmd.(commands.V3PositionManagerPermit)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Spender", c.Spender)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Amount", c.Amount)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Deadline", c.Deadline)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Sig", fmt.Sprintf("%x", c.Sig))
		case commands.V3_POSITION_MANAGER_CALL:
			c := cmd.(commands.V3PositionManagerCall)
			printActions(c)
		case commands.V4_INITIALIZE_POOL:
			c := cmd.(commands.V4InitializePool)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Key", c.Key.ID())
			fmt.Fprintf(os.Stdout, actionArgSubFmt, "", "Currency0", c.Key.Currency0)
			fmt.Fprintf(os.Stdout, actionArgSubFmt, "", "Currency1", c.Key.Currency1)
			fmt.Fprintf(os.Stdout, actionArgSubFmt, "", "Fee", c.Key.Fee)
			fmt.Fprintf(os.Stdout, actionArgSubFmt, "", "TickSpacing", c.Key.TickSpacing)
			fmt.Fprintf(os.Stdout, actionArgSubFmt, "", "Hooks", c.Key.Hooks)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "SqrtPrice", c.SqrtPrice)
		case commands.V4_POSITION_MANAGER_CALL:
			c := cmd.(commands.V4PositionManagerCall)
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "Deadline", c.Deadline)
			printActions(c)

		// Not Implemented
		case commands.EXECUTE_SUB_PLAN,
			commands.FLAG_ALLOW_REVERT,
			commands.COMMAND_TYPE_MASK:
			// TODO Not implemented
		default:
		}
	}
}

//func calldata_raw() {
//	if cmd.Actions() != nil {
//		fmt.Printf("\t[%s] actions: %d\n\t\t%+v\n", cmd.Type(), len(cmd.Actions()), cmd)
//	} else {
//		fmt.Printf("\t[%s] \n\t\t%+v\n", cmd.Type(), cmd)
//	}
//	//switch cmd.Type() {
//	//case commands.V4_SWAP:
//	//	swap := cmd.(commands.V4Swap)
//	//	for _, action := range swap.Actions {
//	//		//fmt.Printf("\t[%T] %+v\n", action, action)
//	//		fmt.Printf("\t\t[%s] %+v\n", action.Type(), action)
//	//	}
//	//case commands.V4_POSITION_MANAGER_CALL:
//	//	swap := cmd.(commands.V4PositionManagerCall)
//	//	for _, action := range swap.Actions {
//	//		//fmt.Printf("\t[%T] %+v\n", action, action)
//	//		fmt.Printf("\t\t[%s] %+v\n", action.Type(), action)
//	//	}
//	//}
//	switch cmd.Type() {
//	case commands.V3_POSITION_MANAGER_CALL:
//		call := cmd.(commands.V3PositionManagerCall)
//		fmt.Printf("\t\t[%s] %+v\n", call.Action.Type(), call.Action)
//	default:
//		for _, action := range cmd.Actions() {
//			fmt.Printf("\t\t[%s] %+v\n", action.Type(), action)
//		}
//	}
//}

var usageFmt = `Usage: %s [options...] command [args...]
`

func usage() {
	c := os.Args[0]
	fmt.Fprintf(os.Stdout, usageFmt, c)
}

var (
	jsonFlag       bool
	jsonPrettyFlag bool
)

func main() {
	mainFlags := flag.NewFlagSet("main", flag.ExitOnError)
	calldataFlags := flag.NewFlagSet("calldata", flag.ExitOnError)

	calldataFlags.BoolVar(&jsonFlag, "json", false, "outputs results in JSON")
	calldataFlags.BoolVar(&jsonPrettyFlag, "jsonpretty", false, "outputs results in pretty JSON")

	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]
	switch command {
	case "":
		err := mainFlags.Parse(args)
		checkErr("", err)

		mainFlags.Parse(mainFlags.Args())
	case "calldata":
		err := calldataFlags.Parse(args)
		checkErr("", err)

		args = calldataFlags.Args()
		if len(args) != 1 {
			usage()
			os.Exit(1)
		}

		if !jsonFlag && !jsonPrettyFlag {
			calldataCmd(args...)
		} else {
			calldataCmd_JSON(jsonPrettyFlag, args...)
		}
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", command)
		usage()
		os.Exit(1)
	}
}
