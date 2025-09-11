package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/juztin/unidecode"
	"github.com/juztin/unidecode/actions"
	"github.com/juztin/unidecode/commands"
)

const (
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
			fmt.Fprintf(os.Stdout, actionArgFmt, "", "AmountOutMinimum", a.AmountOutMinimum)
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

func printExecute(execute unidecode.Execute) {
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

func process(isJSON, isJSONPretty bool, calldata []byte) {
	t := unidecode.MessageType(calldata)
	if t != unidecode.ExecuteMessage {
		checkErr("currently only supports EXECUTE messages; %s", fmt.Errorf("got %x but expected %s", calldata, unidecode.ExecuteMessage))
	}
	execute, err := unidecode.DecodeExecute(calldata)
	checkErr("", err)

	if isJSON || isJSONPretty {
		var b []byte
		if isJSONPretty {
			b, err = json.MarshalIndent(execute, "", "  ")
		} else {
			b, err = json.Marshal(execute)
		}
		checkErr("", err)

		fmt.Fprintf(os.Stdout, "%s\n", b)
	} else {
		printExecute(execute)
	}
}

func calldataCmd(isJSON, isPretty bool, calldata []byte) {
	if len(calldata) == 0 {
		checkErr("", fmt.Errorf("empty/missing CALLDATA"))
	}
	if bytes.HasPrefix(calldata, []byte("0x")) {
		calldata = calldata[2:]
	}

	b := make([]byte, len(calldata))
	_, err := hex.Decode(b, calldata)
	checkErr("", err)

	process(isJSON, isPretty, b)
}

func transactionCmd(ctx context.Context, rpcURL string, isJSON, isPretty bool, hash common.Hash) {
	client, err := ethclient.DialContext(ctx, rpcURL)
	checkErr("", err)

	tx, _, err := client.TransactionByHash(ctx, hash)
	checkErr("", err)

	process(isJSON, isPretty, tx.Data())
}

func pipedOrArg(args []string) ([]byte, error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}

	hasPipedData := (info.Mode() & os.ModeCharDevice) == 0
	if len(args) > 0 && hasPipedData {
		return nil, fmt.Errorf("can't accept both piped data and argument")
	} else if len(args) == 0 && !hasPipedData {
		return nil, fmt.Errorf("missing stdin and argument")
	}

	var b []byte
	if len(args) > 0 {
		b = []byte(args[0])
	} else {
		in, err := io.ReadAll(os.Stdin)
		checkErr("", err)

		s := string(in)
		s = strings.TrimSpace(s)
		s = strings.Trim(s, "\"")
		s = strings.TrimPrefix(s, "0x")
		b = []byte(s)
	}
	return b, err
}

var usageFmt = `Usage: %s [options...] command [args...]

COMMANDS:

  calldata [FLAGS] CALLDATA       decodes raw execute calldata
  tx [FLAGS] HASH                 decodes an execute transactions calldata

FLAGS

  -json                           outputs compressed JSON
  -jsonpretty                     outputs pretty JSON

  tx
    -rpc                          URL of an RPC API endpoint (DEFAULT: http://localhost:8545)

EXAMPLES

  unidecode calldata "3593564c000000000000000000000..."
  unidecode calldata -json "3593564c000000000000000000000..."
  unidecode calldata -jsonpretty "3593564c000000000000000000000..."

  unidecode tx 0x061c3b1c472fb8727648e106e002b620b413e46d2b3e50d70dd7cefac4e110a3
  unidecode tx -json 0x061c3b1c472fb8727648e106e002b620b413e46d2b3e50d70dd7cefac4e110a3
  unidecode tx -jsonpretty 0x061c3b1c472fb8727648e106e002b620b413e46d2b3e50d70dd7cefac4e110a3

  unidecode tx \
    -rpc "https://mainnet.infura.io/v3/00000000000000000000000000000000" \
    0x061c3b1c472fb8727648e106e002b620b413e46d2b3e50d70dd7cefac4e110a3

  unidecode tx \
    -jsonpretty \
    -rpc "https://mainnet.infura.io/v3/00000000000000000000000000000000" \
    0x061c3b1c472fb8727648e106e002b620b413e46d2b3e50d70dd7cefac4e110a3
`

func usage() {
	c := os.Args[0]
	fmt.Fprintf(os.Stdout, usageFmt, c)
}

var (
	jsonFlag       bool
	jsonPrettyFlag bool
	rpcURL         string
)

func main() {
	mainFlags := flag.NewFlagSet("main", flag.ExitOnError)
	calldataFlags := flag.NewFlagSet("calldata", flag.ExitOnError)
	txFlags := flag.NewFlagSet("tx", flag.ExitOnError)

	calldataFlags.BoolVar(&jsonFlag, "json", false, "outputs results in JSON")
	calldataFlags.BoolVar(&jsonPrettyFlag, "jsonpretty", false, "outputs results in pretty JSON")

	txFlags.BoolVar(&jsonFlag, "json", false, "outputs results in JSON")
	txFlags.BoolVar(&jsonPrettyFlag, "jsonpretty", false, "outputs results in pretty JSON")
	txFlags.StringVar(&rpcURL, "rpc", "http://localhost:8545", "JSON RPC URL")

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

		b, err := pipedOrArg(args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			usage()
			os.Exit(1)
		}
		calldataCmd(jsonFlag, jsonPrettyFlag, b)
	case "tx":
		err := txFlags.Parse(args)
		checkErr("", err)

		args = txFlags.Args()
		b, err := pipedOrArg(args)
		if err != nil {
			fmt.Fprint(os.Stderr, "%s\n", err)
			usage()
			os.Exit(1)
		}
		hash := common.BytesToHash(b)
		ctx := context.Background()
		transactionCmd(ctx, rpcURL, jsonFlag, jsonPrettyFlag, hash)
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", command)
		usage()
		os.Exit(1)
	}
}
