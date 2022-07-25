package address

import (
	"crypto/ecdsa"
	"strings"

	"github.com/anaskhan96/base58check"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

const tronPrefix = "0x41"
const base58Version = "41"

func GetBase58CheckFromAddressHex(addressHex string) string {
	base58, _ := base58check.Encode(base58Version, addressHex[4:])
	return base58
}

func GetAddressHexFromPublicKeyHex(publicKeyHex string) string {
	if strings.ToLower(publicKeyHex[0:2]) != "0x" {
		publicKeyHex = "0x" + publicKeyHex
	}
	publicKeyBytes, _ := hexutil.Decode(publicKeyHex)
	if len(publicKeyBytes) == 65 {
		publicKeyBytes = publicKeyBytes[1:]
	}
	hash := crypto.Keccak256(publicKeyBytes)
	hashString := hexutil.Encode(hash)
	addressHex := tronPrefix + hashString[len(hashString)-40:]
	return addressHex
}

func GetPublicKeyFromPrivateKey(privateKey string) string {
	if strings.ToLower(privateKey[0:2]) == "0x" {
		privateKey = privateKey[2:]
	}
	privateKeyECDSA, _ := crypto.HexToECDSA(privateKey)

	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	return hexutil.Encode(publicKeyBytes)[2:]
}

func GetFromPrivateKey(privateKey string) (string, string, string) {
	publicKeyHex := GetPublicKeyFromPrivateKey(privateKey)
	addressHex := GetAddressHexFromPublicKeyHex(publicKeyHex)
	base58 := GetBase58CheckFromAddressHex(addressHex)
	return privateKey, addressHex, base58
}

func Generate() (string, string, string) {
	privateKey, _ := crypto.GenerateKey()
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := hexutil.Encode(privateKeyBytes)
	return GetFromPrivateKey(privateKeyHex)
}
