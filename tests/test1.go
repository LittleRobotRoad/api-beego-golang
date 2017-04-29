package test

import (
	"crypto/aes"
	"crypto/cipher"
)


func main() {

}

func Decrypt(ciphertext, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCEncrypter(block, iv[:aes.BlockSize])
	mode.CryptBlocks(ciphertext, ciphertext)
	return ciphertext, nil
}
