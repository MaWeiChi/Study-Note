// package main

// import (
// 	"bytes"
// 	"crypto/aes"
// 	"crypto/cipher"
// 	"crypto/rand"
// 	"crypto/sha256"
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"os"
// )

// func createHash(key string) []byte {
// 	hash := sha256.Sum256([]byte(key))
// 	return hash[:]
// }

// func encrypt(data []byte, passphrase string) []byte {
// 	fmt.Println(len(data))
// 	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
// 	gcm, err := cipher.NewGCM(block)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	nonce := make([]byte, gcm.NonceSize())
// 	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
// 		panic(err.Error())
// 	}
// 	ciphertext := gcm.Seal(nonce, nonce, data, nil)
// 	fmt.Println(len(ciphertext))
// 	return ciphertext
// }

// func decrypt(data []byte, passphrase string) []byte {
// 	key := []byte(createHash(passphrase))
// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	gcm, err := cipher.NewGCM(block)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	nonceSize := gcm.NonceSize()
// 	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
// 	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return plaintext
// }

// func encryptfile(filename string, data []byte, passphrase string) {
// 	f, _ := os.Create(filename)
// 	defer f.Close()
// 	f.Write(encrypt(data, passphrase))
// }

// func decryptFile(filename string, passphrase string) []byte {
// 	data, _ := ioutil.ReadFile(filename)
// 	return decrypt(data, passphrase)
// }

// func encryptFile(fileInput, fileOutput, passphrase string) {
// 	// read content from your file
// 	plaintext, err := ioutil.ReadFile(fileInput)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	// this is a key
// 	key := []byte(createHash(passphrase))

// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// The IV needs to be unique, but not secure. Therefore it's common to
// 	// include it at the beginning of the ciphertext.
// 	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
// 	iv := ciphertext[:aes.BlockSize]
// 	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
// 		panic(err)
// 	}

// 	stream := cipher.NewCFBEncrypter(block, iv)
// 	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

// 	// create a new file for saving the encrypted data.
// 	f, err := os.Create(fileOutput)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	_, err = io.Copy(f, bytes.NewReader(ciphertext))
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	// done
// 	// data, _ := ioutil.ReadFile("a_aes.zip")
// 	// fmt.Println(string(decrypt(data, "example key 1234")))
// }
