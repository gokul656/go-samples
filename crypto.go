package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

var (
	key       = []byte("1234567812345678") // AES-128 key (16 bytes)
	plaintext = []byte("Hello, Block1234") // Must be exactly 16 bytes!
)

func plainEncryption() {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize)
	block.Encrypt(ciphertext, plaintext)

	fmt.Printf("Plain Encrypted : %x\n", ciphertext)
}

func CBCEncryption() {
	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	iv := make([]byte, cipherBlock.BlockSize())
	if _, err := io.ReadFull(rand.Reader, iv); err != nil { // provides randomness
		panic(err)
	}

	cipherText := make([]byte, len(plaintext))
	blockMode := cipher.NewCBCEncrypter(cipherBlock, iv)
	blockMode.CryptBlocks(cipherText, plaintext)

	hexCipherText := hex.EncodeToString(cipherText) // Can also use base64 for string format
	fmt.Printf("Cipher Encrypted: %v\n", hexCipherText)
}
