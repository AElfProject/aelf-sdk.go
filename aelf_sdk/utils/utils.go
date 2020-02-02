package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"unsafe"

	pt "aelf_sdk.go/aelf_sdk/proto"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

//Base58StringToAddress address to  bytes
func Base58StringToAddress(addr string) (*pt.Address, error) {
	var address = new(pt.Address)
	addressBytes, _, err := base58.CheckDecode(addr)
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

//StringToBytes String To Bytes
func StringToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{sh.Data, sh.Len, 0}
	return *(*[]byte)(unsafe.Pointer(&bh))
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

//HashKey Hash publicKey
func HashKey(key []byte) []byte {
	hash256 := sha256.Sum256(key)
	ripemd160Hasher := ripemd160.New()
	_, err := ripemd160Hasher.Write(hash256[:])
	if err != nil {
		fmt.Println("hash key error")
	}
	ripemdHash := ripemd160Hasher.Sum(nil)
	return ripemdHash
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

// SerializeToBytes Serialize To Bytes
func SerializeToBytes(tx *pt.Transaction) []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

// DeserializeToInter Deserialize To Interface
func DeserializeToInter(b []byte) interface{} {
	var data interface{}
	decoder := gob.NewDecoder(bytes.NewReader(b))
	err := decoder.Decode(&data)
	if err != nil {
		log.Panic(err)
	}
	return &data
}
