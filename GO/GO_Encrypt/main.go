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

	"github.com/sirupsen/logrus"
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

func decrypt(data []byte, passphrase string) ([]byte, error) {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, err
}

func decryptFile(filename string, passphrase string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return decrypt(data, passphrase)
}

func main() {
	// fmt.Println("Starting the application...")
	// ciphertext := encrypt([]byte("Hello World"), "password")
	// fmt.Printf("Encrypted: %x\n", ciphertext)
	// plaintext := decrypt(ciphertext, "password")
	// fmt.Printf("Decrypted: %s\n", plaintext)
	// encryptfile("sample.txt", []byte("Hello Word"), "password1")
	// encryptFile("2020-06-05-10-28-backup.zip", "a_aes.tar.gz", "example key 1234")
	// data, _ := ioutil.ReadFile("2020-06-18-17-14-backup.tar.gz")
	// encryptfile("2020-06-18-17-14-backup.tar.gz.encryption", data, "password1")
	data, err := decryptFile("/home/moxa/Erik/Study-Note/GO/GO_Encrypt/backup.tar.gz", "!2345678")
	if err != nil {
		fmt.Println("decryptFile err")
		fmt.Println(err.Error())
	}
	ioutil.WriteFile(filepath.Join(extractDir, "2020-06-18-17-14-backup.tar.gz"), data, 0755)

	data, _ = ioutil.ReadFile("2020-06-18-17-14-backup.tar.gz")
	// encryptfile("2020-06-18-17-14-backup.tar.gz.encryption", data, "password1")
	filename := "2020-06-18-17-14-backup.tar.gz"

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
		if err != nil {
			fmt.Println(err.Error())
		}
		if f.FileInfo().IsDir() {
			continue
		}
		logrus.Debugf("extract %s", f.Name)
		os.MkdirAll(filepath.Join(extractDir, filepath.Dir(f.Name)), 0755)
		file, err := os.Create(filepath.Join(extractDir, f.Name))
		if err != nil {
			logrus.Debugf("create file %s failed", f.Name)
			continue
		}
		io.Copy(file, tr)

	}
	gr.Close()
	fmt.Println("extractDir: " + extractDir)
}
