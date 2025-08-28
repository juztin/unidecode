package commands

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/juztin/unidecode/actions"
	"github.com/juztin/unidecode/hex"
)

type PermitDetails struct {
	Token      common.Address
	Amount     *big.Int
	Expiration time.Time
	Nonce      uint64
}

type Permit2Permit struct {
	Details     PermitDetails
	Spender     common.Address
	SigDeadline *big.Int
	Sig         []byte
}

func (Permit2Permit) Type() Type {
	return PERMIT2_PERMIT
}

func (Permit2Permit) Actions() []actions.Action {
	return nil
}

func DecodePermit2Permit(calldata []byte, offset int) (Permit2Permit, error) {
	// Dispatcher.sol:195
	//
	// permit(address owner, PermitSingle memory permitSingle, bytes calldata signature)
	//
	//                ┌────── PermitDetails ─────────┐
	//        owner   │ token   amount  expire nonce   spender deadline sig
	// permit(address,((address,uint160,uint48,uint48),address,uint256),bytes)
	// -------------------------------------------------------------------------------------------------------
	//    struct PermitDetails {
	//        address token;
	//        uint160 amount;
	//        uint48 expiration;
	//        uint48 nonce;
	//    }
	//    struct PermitSingle {
	//        PermitDetails details;
	//        address spender;
	//        uint256 sigDeadline;
	//    }

	var p Permit2Permit

	dataLen, err := hex.Int(calldata[offset : offset+0x20])
	if err != nil {
		return p, fmt.Errorf("invalid %s input count; %w", err)
	} else if dataLen > len(calldata)-offset {
		return p, fmt.Errorf("invalid %s data length", PERMIT2_PERMIT)
	}
	offset += 0x20

	p.Details = PermitDetails{
		Token:  common.BytesToAddress(calldata[offset : offset+0x20]),
		Amount: new(big.Int).SetBytes(calldata[offset+0x20 : offset+0x40]),
	}

	epoch, err := hex.Int64(calldata[offset+0x40 : offset+0x60])
	if err != nil {
		return p, fmt.Errorf("invalid PermitDetails expiration; %w", err)
	}
	p.Details.Expiration = time.Unix(epoch, 0)

	nonce, err := hex.Int64(calldata[offset+0x60 : offset+0x80])
	if err != nil {
		return p, fmt.Errorf("invalid PermitDetails nonce; %w", err)
	}
	p.Details.Nonce = uint64(nonce)

	p.Spender = common.BytesToAddress(calldata[offset+0x80 : offset+0xa0])
	p.SigDeadline = new(big.Int).SetBytes(calldata[offset+0xa0 : offset+0xc0])

	sigLen, err := hex.Int(calldata[offset+0xc0 : offset+0xe0])
	if err != nil {
		return p, fmt.Errorf("invalid Permit2 signature length; %w", err)
	}
	p.Sig = calldata[offset+0x100 : offset+0x100+sigLen]

	return p, nil
}

type Permit2PermitBatch struct {
	Details     []PermitDetails
	Spender     common.Address
	SigDeadline *big.Int
	Sig         []byte
}

func (Permit2PermitBatch) Type() Type {
	return PERMIT2_PERMIT_BATCH
}

func (Permit2PermitBatch) Actions() []actions.Action {
	return nil
}

func DecodePermit2PermitBatch(calldata []byte, offset int) (Permit2PermitBatch, error) {
	// Dispatcher.sol:109
	//
	//                ┌────── PermitDetails ─────────┐
	//        owner   │ token   amount  expire nonce     spender deadline sig
	// permit(address,((address,uint160,uint48,uint48)[],address,uint256),bytes)
	var p Permit2PermitBatch
	// TODO Implement; should be the same as permit2, with an array
	return p, errNotImplemented
}

type Permit2TransferFrom struct {
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
}

func (Permit2TransferFrom) Type() Type {
	return PERMIT2_TRANSFER_FROM
}

func (Permit2TransferFrom) Actions() []actions.Action {
	return nil
}

func DecodePermit2TransferFrom(calldata []byte, offset int) (Permit2TransferFrom, error) {
	return Permit2TransferFrom{
		Token:     common.BytesToAddress(calldata[offset : offset+0x20]),
		Recipient: common.BytesToAddress(calldata[offset+0x20 : offset+0x40]),
		Amount:    new(big.Int).SetBytes(calldata[offset+0x40 : offset+0x60]),
	}, nil
}

type Permit2TransferFromBatch struct{}

func (Permit2TransferFromBatch) Type() Type {
	return PERMIT2_TRANSFER_FROM_BATCH
}

func (Permit2TransferFromBatch) Actions() []actions.Action {
	return nil
}

func DecodePermit2TransferFromBatch(calldata []byte, offset int) (Permit2TransferFromBatch, error) {
	// permit(address,((address,uint160,uint48,uint48)[],address,uint256),bytes)
	return Permit2TransferFromBatch{}, errNotImplemented
}
