package unidecode

import (
	"bytes"
	"errors"
	"fmt"
	"time"

	"github.com/juztin/unidecode/commands"
	"github.com/juztin/unidecode/hex"
)

var (
	hexPrefix              = []byte{0x30, 0x78}
	executeSig             = []byte{0x24, 0x85, 0x6b, 0xc3}
	executeWithDeadlineSig = []byte{0x35, 0x93, 0x56, 0x4c}

	executeSigStr             = fmt.Sprintf("%x", executeSig)
	executeWithDeadlineSigStr = fmt.Sprintf("%x", executeWithDeadlineSig)
)

var (
	ErrInvalidCallData    = errors.New("invalid calldata")
	ErrIncorrectMethodSig = errors.New("invalid method signature")
)

type messageType int

const (
	UnknownMessage messageType = iota
	ExecuteMessage
)

func (t messageType) String() string {
	switch t {
	case ExecuteMessage:
		return "EXECUTE"
	default:
		return "UNKNOWN"
	}
}

type Execute struct {
	Commands []commands.Command `json:"commands"`
	Deadline *time.Time         `json:"deadline"`
}

//func (e Execute) MarshalJSON() ([]byte, error) {
//	return json.Marshal(&struct {
//		//Execute struct {
//		Commands []commands.Command `json:"commands"`
//		Deadline *time.Time         `json:"deadline"`
//		//} `json:"execute"`
//	}{
//		Commands: e.Commands,
//		Deadline: e.Deadline,
//	})
//}

func (e Execute) Swap() *commands.V4Swap {
	for i := range e.Commands {
		if e.Commands[i].Type() == commands.V4_SWAP {
			cmd := e.Commands[i]
			swap := cmd.(commands.V4Swap)
			return &swap
		}
	}
	return nil
}

func MessageType(calldata []byte) messageType {
	if len(calldata) < 8 {
		return UnknownMessage
	}

	sig := fmt.Sprintf("%x", hex.MethodSig(calldata))
	switch sig {
	case executeSigStr, executeWithDeadlineSigStr:
		return ExecuteMessage
	}
	return UnknownMessage
}

func DecodeExecute(calldata []byte) (Execute, error) {
	var e Execute
	if bytes.HasPrefix(calldata, hexPrefix) {
		calldata = calldata[0x02:]
	}

	sig := hex.MethodSig(calldata)
	if sig == nil {
		return e, ErrInvalidCallData
	} else if bytes.Compare(sig, executeSig) != 0 && bytes.Compare(sig, executeWithDeadlineSig) != 0 {
		return e, fmt.Errorf("%w; expected one of [%x, %x] but got %x",
			ErrIncorrectMethodSig, executeSig, executeWithDeadlineSig, sig)
	}

	// Remove method signature
	calldata = calldata[4:]

	// Get command start location
	commandStart, err := hex.Int(calldata[:0x20])
	if err != nil {
		return e, fmt.Errorf("invalid command start memory location; %w", err)
	}

	commandLen, err := hex.Int(calldata[commandStart : commandStart+0x20])
	if err != nil {
		return e, fmt.Errorf("invalid command length value; %w", err)
	}

	e = Execute{}
	if commandStart == 0x60 {
		epoch, err := hex.Int64(calldata[0x40:0x60])
		if err != nil {
			return e, fmt.Errorf("invalid deadline value; %w", err)
		}
		deadline := time.Unix(epoch, 0)
		e.Deadline = &deadline
	}

	offset := commandStart + 0x20
	commandTypes, err := commands.DecodeType(calldata[offset : offset+commandLen])
	if err != nil {
		return e, fmt.Errorf("invalid commands; %w", err)
	}

	if commandLen > 0x20 {
		// TODO
	} else {
		offset += 0x20
	}

	inputsLen, err := hex.Int(calldata[offset : offset+0x20])
	if err != nil {
		return e, fmt.Errorf("invalid inputs length; %w", err)
	} else if inputsLen != commandLen {
		return e, fmt.Errorf("inputs length mismatch; got %d but expected %d", inputsLen, commandLen)
	}
	offset += 0x20

	var inputsLoc []int
	for i := 0; i < inputsLen; i++ {
		inputOffset := offset + i*0x20
		loc, err := hex.Int(calldata[inputOffset : inputOffset+0x20])
		if err != nil {
			return e, fmt.Errorf("invalid input location for intput %d; %w", i, err)
		}
		inputsLoc = append(inputsLoc, loc)
	}

	for i, t := range commandTypes {
		c, err := commands.Decode(t, calldata, offset+inputsLoc[i])
		if err != nil {
			return e, fmt.Errorf("invalid command data at index %d for %s; %w", i, t, err)
		}
		e.Commands = append(e.Commands, c)
	}
	return e, nil
}
