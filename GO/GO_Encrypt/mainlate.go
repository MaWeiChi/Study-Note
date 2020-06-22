package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
	"os"
)

func encryptFile(fileInput, fileOutput, passphrase string) {
	// read content from your file
	plaintext, err := ioutil.ReadFile(fileInput)
	if err != nil {
		panic(err.Error())
	}

	// this is a key
	key := []byte(createHash(passphrase))

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// create a new file for saving the encrypted data.
	f, err := os.Create(fileOutput)
	if err != nil {
		panic(err.Error())
	}
	_, err = io.Copy(f, bytes.NewReader(ciphertext))
	if err != nil {
		panic(err.Error())
	}

	// done
	// data, _ := ioutil.ReadFile("a_aes.zip")
	// fmt.Println(string(decrypt(data, "example key 1234")))
}

// func decrypt() {
// 	key := []byte("example key 1234")
// 	ciphertext, err := ioutil.ReadFile("a_aes.zip")
// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	iv := make([]byte, aes.BlockSize)
// 	if _, err := rand.Read(iv); err != nil {
// 		log.Fatal(err)
// 	}
// 	newtext := make([]byte, len(ciphertext))
// 	dec := cipher.NewCBCDecrypter(block, iv)
// 	dec.CryptBlocks(newtext, ciphertext)
// }

func decrypt2(data []byte, passphrase string) []byte {
	key := []byte((passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}
