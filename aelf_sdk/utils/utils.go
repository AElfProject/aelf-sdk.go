package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"unsafe"

	pt "aelf_sdk.go/aelf_sdk/proto"
	"github.com/btcsuite/btcutil/base58"
)

//Base58StringToAddress address to  bytes
func Base58StringToAddress(addr string) (*pt.Address, error) {
	var address = new(pt.Address)
	addressBytes, err := DecodeCheck(addr)
	if err != nil {
		return nil, errors.New("Base58String To Address error")
	}
	address.Value = addressBytes
	return address, nil
}

//BytesToString  Bytes To String
func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

//StringTo32Bytes String To 32Bytes
func StringTo32Bytes(s string) []byte {
	var bytes32 [32]byte
	copy(bytes32[:], s)
	return bytes32[:]
}

//GetBytesSha256 Get Bytes Sha256
func GetBytesSha256(s string) []byte {
	sha := sha256.New()
	sha.Write([]byte(s))
	return sha.Sum(nil)
}

//BytesToInt Bytes To Int
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var data int32
	binary.Read(bytesBuffer, binary.BigEndian, &data)
	return int(data)
}

// EncodeCheck prepends appends a four byte checksum.
func EncodeCheck(input []byte) string {
	b := make([]byte, 0, 1+len(input)+4)
	b = append(b, input[:]...)
	cksum := checksum(b)
	b = append(b, cksum[:]...)
	return base58.Encode(b)
}

// checksum: first four bytes of sha256^2
func checksum(input []byte) (cksum [4]byte) {
	h := sha256.Sum256(input)
	h2 := sha256.Sum256(h[:])
	copy(cksum[:], h2[:4])
	return
}

// DecodeCheck decodes a string that was encoded with CheckEncode and verifies the checksum.
func DecodeCheck(input string) (result []byte, err error) {
	decoded := base58.Decode(input)
	if len(decoded) < 5 {
		return nil, errors.New("invalid format: version and/or checksum bytes missing")
	}
	var cksum [4]byte
	copy(cksum[:], decoded[len(decoded)-4:])
	if checksum(decoded[:len(decoded)-4]) != cksum {
		return nil, errors.New("checksum error")
	}
	payload := decoded[0 : len(decoded)-4]
	result = append(result, payload...)
	return
}

//ParamsToString Params To String like "AElf.ContractNames.Consensus"
func ParamsToString(params string) string {
	paramMap := map[string]interface{}{
		"value": base64.StdEncoding.EncodeToString(GetBytesSha256(params)),
	}
	paramsBytes, err := json.Marshal(paramMap)
	if err != nil {
		fmt.Println("json Marshal error")
	}
	return string(paramsBytes)
}

//GetAddressByBytes  sha256 twice Bytes to get address
func GetAddressByBytes(b []byte) string {
	firstBytes := sha256.Sum256(b)
	secondBytes := sha256.Sum256(firstBytes[:])
	address := EncodeCheck(secondBytes[:])
	return address
}
