package address

import (
	"testing"
)

func TestGetPublicKeyFromPrivateKey1(t *testing.T) {
	privateKey := "0xf5737b0a37b87451a52863e0bbec15ab06983540f41b5b06b322ad1a1f8eef72"
	address := GetFromPrivateKey(privateKey)
	expectedAddressHex := "0x41a053cb629b2c281824381782d2311cbfb706ddc0"
	expectedBase58Address := "TQawVtoUz23osgePwjui9Fci8hFM9a3med"
	if address.AddressHex != expectedAddressHex {
		t.Errorf("GetPublicKeyFromPrivateKey(\"%s\")'s hexAddress = \"%s\"; want \"%s\"", privateKey, address.AddressHex, expectedAddressHex)
	}
	if address.Base58Address != expectedBase58Address {
		t.Errorf("GetPublicKeyFromPrivateKey(\"%s\")'s base58Address = \"%s\"; want \"%s\"", privateKey, address.Base58Address, expectedBase58Address)
	}
}

func TestGetPublicKeyFromPrivateKey2(t *testing.T) {
	privateKey := "0x22ccba98bc5a03a0743a00666fc91e866a46381f35802b39a66c9fe80f69600f"
	address := GetFromPrivateKey(privateKey)
	expectedAddressHex := "0x4157a2d1c863d858cdb3680843956d16d8f0682232"
	expectedBase58Address := "THxaqeF53iHB8PRaKAuNzuJy7pEikMGiUn"
	if address.AddressHex != expectedAddressHex {
		t.Errorf("GetPublicKeyFromPrivateKey(\"%s\")'s hexAddress = \"%s\"; want \"%s\"", privateKey, address.AddressHex, expectedAddressHex)
	}
	if address.Base58Address != expectedBase58Address {
		t.Errorf("GetPublicKeyFromPrivateKey(\"%s\")'s base58Address = \"%s\"; want \"%s\"", privateKey, address.Base58Address, expectedBase58Address)
	}
}
