package guardian

import (
	"encoding/hex"
	"fmt"
	"math/big"

	ec "github.com/ethereum/go-ethereum/common"
)

func ToPadBytes(v interface{}, rlen ...int) []byte {
	var orig []byte
	var retlen int
	switch k := v.(type) {
	case uint8:
		retlen = 1
		orig = big.NewInt(int64(k)).Bytes()
	case uint32:
		retlen = 4
		orig = big.NewInt(int64(k)).Bytes()
	case int64:
		retlen = 8
		orig = big.NewInt(k).Bytes()
	case uint64:
		retlen = 8
		orig = new(big.Int).SetUint64(k).Bytes()
	case *big.Int:
		if len(rlen) == 1 {
			retlen = rlen[0]
		} else {
			retlen = 32
		}
		orig = k.Bytes()
	default:
		return nil
	}
	ret := make([]byte, retlen)
	copy(ret[retlen-len(orig):], orig)
	return ret
}

func Hex2Bytes(s string) (b []byte) {
	if len(s) >= 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
		s = s[2:]
	}
	// hex.DecodeString expects an even-length string
	if len(s)%2 == 1 {
		s = "0" + s
	}
	b, _ = hex.DecodeString(s)
	return b
}

func Hex2Hash(s string) ec.Hash {
	return ec.HexToHash(s)
}

// Bytes2Hash converts bytes to Hash
func Bytes2Hash(b []byte) ec.Hash {
	return ec.BytesToHash(b)
}

// Hex2Addr accepts hex string with or without 0x prefix and return Addr
func Hex2Addr(s string) ec.Address {
	return ec.HexToAddress(s)
}

// if in is 20 bytes, return directly. otherwise return a new []byte w/ len 20, left pad 0x00..[in]
// if len(in)>20, panic
func Pad20Bytes(in []byte) []byte {
	origLen := len(in)
	if origLen == 20 {
		return in
	}
	if origLen > 20 {
		panic(fmt.Sprintf("%x len %d > 20", in, origLen))
	}
	ret := make([]byte, 20)
	copy(ret[20-origLen:], in)
	return ret
}

func Pad32Bytes(in []byte) []byte {
	origLen := len(in)
	if origLen == 32 {
		return in
	}
	if origLen > 32 {
		panic(fmt.Sprintf("%x len %d > 32", in, origLen))
	}
	ret := make([]byte, 32)
	copy(ret[32-origLen:], in)
	return ret
}
