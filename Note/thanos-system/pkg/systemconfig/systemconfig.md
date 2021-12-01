# systemconfig

The tar.gz is  encrypted with our until after compressed and packed.

Decrypt with the following code:

```[go]
package main

import (
	"archive/tar"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

var extractDir = "."

func createHash(key string) []byte {
	hash := sha256.Sum256([]byte(key))
	return hash[:]
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

	data, e := decryptFile("backup.tar.gz", "!2345678")
	if e != nil {
		fmt.Println(e)
	}

	ioutil.WriteFile(filepath.Join(extractDir, "backup.tar"), data, 0755)

	filename := "backup.tar"

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
		if f.FileInfo().IsDir() {
			continue
		}
		fmt.Printf("extract %s", f.Name)
		os.MkdirAll(filepath.Join(extractDir, filepath.Dir(f.Name)), 0755)
		file, err := os.Create(filepath.Join(extractDir, f.Name))
		if err != nil {
			fmt.Printf("create file %s failed", f.Name)
			continue
		}
		io.Copy(file, tr)

	}
	gr.Close()
	fmt.Println("extractDir: " + extractDir)
}
```