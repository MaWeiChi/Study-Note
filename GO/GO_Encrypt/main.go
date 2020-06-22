package main

import (
	"archive/tar"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

var extractDir = "."

// turn password into 32 byte to use aes-256
// func createHash(key string) string {
// 	hasher := md5.New()
// 	hasher.Write([]byte(key))
// 	return hex.EncodeToString(hasher.Sum(nil))
// }

// turn password into 32 byte to use aes-256
func createHash(key string) []byte {
	hash := sha256.Sum256([]byte(key))
	return hash[:]
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
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

func encryptfile(filename string, data []byte, passphrase string) {
	f, _ := os.Create(filename)
	defer f.Close()
	f.Write(encrypt(data, passphrase))
}

func decryptFile(filename string, passphrase string) []byte {
	data, _ := ioutil.ReadFile(filename)
	return decrypt(data, passphrase)
}

func main() {
	fmt.Println("Starting the application...")
	ciphertext := encrypt([]byte("Hello World"), "password")
	fmt.Printf("Encrypted: %x\n", ciphertext)
	plaintext := decrypt(ciphertext, "password")
	fmt.Printf("Decrypted: %s\n", plaintext)
	encryptfile("sample.txt", []byte("Hello Word"), "password1")
	encryptFile("2020-06-05-10-28-backup.zip", "a_aes.tar.gz", "example key 1234")
	data, _ := ioutil.ReadFile("2020-06-18-17-14-backup.tar.gz")
	encryptfile("2020-06-18-17-14-backup.tar.gz.encryption", data, "password1")
	data = decryptFile("sample.tar.gz", "password1")

	ioutil.WriteFile(filepath.Join(extractDir, "2020-06-18-17-14-backup.tar."), data, 0755)

	data, _ = ioutil.ReadFile("2020-06-18-17-14-backup.tar.gz")
	encryptfile("2020-06-18-17-14-backup.tar.gz.encryption", data, "password1")

	filename := "2020-06-19-17-42-backup.tar.gz"

	br, _ := os.Open(filename)
	gr, err := gzip.NewReader(br)
	if err != nil {
		fmt.Println(err.Error())
	}
	tr := tar.NewReader(gr)
	if len(filepath.Base(filename)) <= 7 {
		fmt.Println(err.Error())
	}
	extractDir := filepath.Join(filepath.Dir(filename), filepath.Base(filename)[:len(filepath.Base(filename))-7])
	os.RemoveAll(extractDir)
	os.MkdirAll(extractDir, 0755)
	for {
		f, err := tr.Next()
		if err == io.EOF {
			break
		}
		// if f.FileInfo().IsDir() {
		// 	continue
		// }

		os.MkdirAll(filepath.Join(extractDir, filepath.Dir(f.Name)), 0755)
		data := make([]byte, f.Size)
		_, err = tr.Read(data)
		if err != nil && err != io.EOF {
		}

		ioutil.WriteFile(filepath.Join(extractDir, f.Name), data, 0755)

	}
	gr.Close()
	fmt.Println("extractDir: " + extractDir)
}
