package encrypt_test

import (
	"goToolLab/encrypt"
	"testing"
)

func TestBase64EncodeByte(t *testing.T) {
	test := []byte(`test to encode and decode, 
			if decode equal to this text, it works`)
	encodeText := encrypt.Base64EncodeByte(test)
	decodeText, err := encrypt.Base64DecodeByte(encodeText)
	if err != nil || string(decodeText) != string(test) {
		t.Error("decode fail")
	}
}

func TestAESEncodeByte(t *testing.T) {
	test := []byte(`test to encode and decode, 
			if decode equal to this text, it works`)
	key := []byte("m@6OxRDiPSjW!32fUSMS!PLASCUdNT##")

	encodeText, errEn := encrypt.AESEncodeByte(test, key)
	if errEn != nil {
		t.Error("encode fail:", errEn.Error())
	}
	decodeText, errDe := encrypt.AESDecodeByte(encodeText, key)
	if errDe != nil || string(decodeText) != string(test) {
		t.Error("decode fail:", errDe.Error())
	}
}
