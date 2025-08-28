package hex

import (
	"fmt"
	"strconv"
)

func MethodSig(calldata []byte) []byte {
	if len(calldata) <= 4 {
		return nil
	}
	return calldata[:0x04]
}

func Int64(b []byte) (int64, error) {
	return strconv.ParseInt(fmt.Sprintf("%x", b), 16, 64)
}

func Int(b []byte) (int, error) {
	i, err := Int64(b)
	if err != nil {
		return 0, err
	}
	return int(i), err
}

func Bool(b []byte) (bool, error) {
	i, err := Int(b)
	if err != nil {
		return false, err
	}
	return i == 1, nil
}
