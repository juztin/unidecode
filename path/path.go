package path

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/juztin/unidecode/hex"
	"github.com/juztin/unidecode/pool"
)

type Key struct {
	IntermediateCurrency common.Address `json:"intermediateCurrency"`
	Fee                  *big.Int       `json:"fee"`
	TickSpacing          *big.Int       `json:"tickSpacing"`
	Hooks                common.Address `json:"hooks"`
	HookData             []byte         `json:"hookData"`
}

func (k Key) MarshalJSON() ([]byte, error) {
	type Alias Key
	return json.Marshal(&struct {
		Alias
		HookData string `json:"hookData"`
	}{(Alias)(k), fmt.Sprintf("0x%x", k.HookData)})
}

func PoolAndSwapDirection(k Key, currencyIn common.Address) (pool.Key, bool) {
	poolKey := pool.NewKey(
		k.IntermediateCurrency,
		currencyIn,
		k.Fee,
		k.TickSpacing,
		k.Hooks,
	)
	zeroForOne := currencyIn.Cmp(poolKey.Currency0) == 0
	return poolKey, zeroForOne
}

func Decode(calldata []byte, offset int) (Key, error) {
	var k Key
	k.IntermediateCurrency = common.BytesToAddress(calldata[offset : offset+0x20])
	k.Fee = new(big.Int).SetBytes(calldata[offset+0x20 : offset+0x40])
	k.TickSpacing = new(big.Int).SetBytes(calldata[offset+0x40 : offset+0x60])
	k.Hooks = common.BytesToAddress(calldata[offset+0x60 : offset+0x80])
	hookDataStart, err := hex.Int(calldata[offset+0x80 : offset+0xa0])
	if err != nil {
		return k, fmt.Errorf("invalid hookData start value; %w", err)
	}
	k.HookData = calldata[offset+hookDataStart : offset+hookDataStart+0x20]
	return k, nil
}

func DecodeMany(calldata []byte, offset int) ([]Key, error) {
	count, err := hex.Int(calldata[offset : offset+0x20])
	if err != nil {
		return nil, fmt.Errorf("invalid path-key count; %w", err)
	}
	offset += 0x20

	var keys []Key
	for i := 0; i < count; i++ {
		keyOffset, err := hex.Int(calldata[offset+0x20*i : offset+0x20*i+0x20])
		if err != nil {
			return keys, fmt.Errorf("invalid path-key index start location for index %d; %w", i, err)
		}
		key, err := Decode(calldata, offset+keyOffset)
		if err != nil {
			return keys, fmt.Errorf("failed to decode path at index %d; %w", i, err)
		}
		keys = append(keys, key)
	}
	return keys, nil
}
