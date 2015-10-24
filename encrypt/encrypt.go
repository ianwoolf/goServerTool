package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func Base64EncodeByte(data []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(data))
}

func Base64DecodeByte(data []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(data))
}

func AESEncodeByte(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return data, err
	}
	var iv = key[:aes.BlockSize]
	blockMode := cipher.NewCFBEncrypter(block, iv)
	dest := make([]byte, len(string(data)))
	blockMode.XORKeyStream(dest, data)
	return dest, nil
}

func AESDecodeByte(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return data, err
	}
	var iv = key[:aes.BlockSize]
	blockMode := cipher.NewCFBDecrypter(block, iv)
	dest := make([]byte, len(string(data)))
	blockMode.XORKeyStream(dest, data)
	return dest, nil
}

//	body := SetOfflineStatus{User: "libaokun", Reason: "down", Ips: []string{"10.222.8.102"}}
//	txt, _ := json.Marshal(body)
//	const SALT = "m!6KkRDiHQjW!2Of2TZw!Si92BJdNT##"
//	aesResult, _ := AESEncodeByte(txt, []byte(SALT))
//	QueryApi("http://sb.cmdb.xiaojukeji.com/cmdb/v1/server/offline", string(aesResult))
