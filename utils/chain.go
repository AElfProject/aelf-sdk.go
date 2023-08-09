package utils

import (
	"encoding/binary"
	"github.com/btcsuite/btcutil/base58"
)

func ConvertBase58ToChainId(base58String string) int32 {
	bytes := base58.Decode(base58String)
	if len(bytes) < 4 {
		bytes = append(make([]byte, 0), bytes...)
		bytes = bytes[:4]
	}
	return int32(binary.LittleEndian.Uint32(bytes))
}
