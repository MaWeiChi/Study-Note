package main

import (
	"archive/tar"
	"bytes"
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

func main() {
	unCompressEncrypt()
}

func unCompressEncrypt() {
	data := decryptFile("2020-08-04-11-05-backup.tar.gz", "fdaf1@123")

	ioutil.WriteFile("Encrypt.tar.gz", data, 0755)
	r := bytes.NewReader(data)
	os.RemoveAll(ExtractEnDir)
	os.MkdirAll(ExtractEnDir, 0755)
	gr, err := gzip.NewReader(r)
	if err != nil {
		fmt.Println(err.Error())
	}
	tr := tar.NewReader(gr)

	for {
		f, err := tr.Next()
		if err == io.EOF {
			break
		}
		if f.FileInfo().IsDir() {
			continue
		}

		os.MkdirAll(filepath.Join(ExtractEnDir, filepath.Dir(f.Name)), 0755)
		data := make([]byte, f.Size)
		_, err = tr.Read(data)
		if err != nil && err != io.EOF {
			fmt.Println(err.Error())
		}

		ioutil.WriteFile(filepath.Join(ExtractEnDir, f.Name), data, 0755)

	}
	gr.Close()
	fmt.Println("extractDir: " + ExtractEnDir)

}

var (

	// BaseDir is base path of whole project
	BaseDir = "/home/moxa/study-note/GO/GO_Compress&Encrypt/config"

	// TempDir is temp folder
	TempDir = "/tmp"

	// AppsDir is apps folder
	AppsDir = "/apps"

	// DataDir is data folder
	DataDir = "/data"

	ExtractDir   = "/home/moxa/study-note/GO/GO_Compress&Encrypt/extra"
	ExtractEnDir = "/home/moxa/study-note/GO/GO_Compress&Encrypt/extraEN"
)

func init() {

	if basedir, ok := os.LookupEnv("APPMAN_BASEDIR"); ok {
		BaseDir = basedir
	}
	TempDir = filepath.Join(BaseDir, TempDir)
	AppsDir = filepath.Join(BaseDir, AppsDir)
	DataDir = filepath.Join(BaseDir, DataDir)

}

// JoinBaseDir will join filepath based on BaseDir
func JoinBaseDir(dirs ...string) string {
	return filepath.Join(BaseDir, filepath.Join(dirs...))
}

func createHash(key string) []byte {
	hash := sha256.Sum256([]byte(key))
	return hash[:]
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

func decryptFile(filename string, passphrase string) []byte {
	data, _ := ioutil.ReadFile(filename)
	return decrypt(data, passphrase)
}

func uncompress() {
	os.RemoveAll(ExtractDir)
	os.MkdirAll(ExtractDir, 0755)
	r, _ := os.Open("testReadytoEncrypt.tar.gz")
	fi, err := os.Stat("/home/moxa/study-note/GO/GO_Compress&Encrypt/testReadytoEncrypt.tar.gz")
	// get the size
	size := fi.Size()
	fmt.Print("compress")
	fmt.Println(size)
	gr, err := gzip.NewReader(r)
	if err != nil {
		fmt.Println(err.Error())
	}
	tr := tar.NewReader(gr)

	for {
		f, err := tr.Next()
		if err == io.EOF {
			break
		}
		if f.FileInfo().IsDir() {
			continue
		}

		os.MkdirAll(filepath.Join(ExtractDir, filepath.Dir(f.Name)), 0755)
		data := make([]byte, f.Size)
		_, err = tr.Read(data)
		if err != nil && err != io.EOF {
			fmt.Println(err.Error())
		}

		ioutil.WriteFile(filepath.Join(ExtractDir, f.Name), data, 0755)

	}
	gr.Close()
	fmt.Println("extractDir: " + ExtractDir)

}

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

// func encryptfile(filename string, data []byte, passphrase string) {
// 	f, _ := os.Create(filename)
// 	defer f.Close()
// 	f.Write(encrypt(data, passphrase))
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

// func compress() {

// 	// dir, err := ioutil.TempDir("", "appman-configuration")
// 	// if err != nil {
// 	// 	fmt.Println(err.Error())
// 	// }
// 	// defer os.RemoveAll(dir)
// 	dir := "./"
// 	fw, err := os.Create(filepath.Join(dir, "test.tar.gz"))
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	gzw := gzip.NewWriter(fw)
// 	defer gzw.Close()
// 	tw := tar.NewWriter(gzw)
// 	defer tw.Close()
// 	type Folder struct {
// 		Name    string
// 		Ignores []string
// 	}
// 	folders := []Folder{
// 		Folder{
// 			Name:    filepath.Join(DataDir, "env"),
// 			Ignores: []string{},
// 		},
// 		Folder{
// 			Name:    filepath.Join(DataDir, "sdcard"),
// 			Ignores: []string{},
// 		},
// 		Folder{
// 			filepath.Join(DataDir, "certs"),
// 			[]string{},
// 		},
// 		{
// 			filepath.Join(DataDir, "nginx"),
// 			[]string{},
// 		},
// 		{
// 			filepath.Join(BaseDir, "deviceinfo"),
// 			[]string{},
// 		},
// 		{
// 			JoinBaseDir("device"),
// 			[]string{},
// 		},
// 		{
// 			filepath.Join(DataDir, "redis"),
// 			[]string{},
// 		},
// 	}

// 	if dirs, err := ioutil.ReadDir(AppsDir); err == nil {
// 		for _, d := range dirs {
// 			if d.IsDir() {
// 				folders = append(
// 					folders,
// 					Folder{
// 						filepath.Join(AppsDir, d.Name(), "data"),
// 						[]string{"moxaenv"},
// 					},
// 				)
// 			}
// 		}
// 	} else {
// 		fmt.Println(err.Error())
// 	}
// 	fmt.Println(folders)
// 	// TODO ignore folder
// 	backupFunc := func(path string, finfo os.FileInfo, ignores []string) error {

// 		fout, err := os.Open(path)
// 		if err != nil {
// 			return nil
// 		}
// 		zipFilename := path[len(BaseDir)+1:]

// 		fh, _ := tar.FileInfoHeader(finfo, path)
// 		fh.Name = zipFilename
// 		if err := tw.WriteHeader(fh); err == nil {
// 			io.Copy(tw, fout)
// 		}

// 		fout.Close()
// 		return nil
// 	}
// 	for _, f := range folders {

// 		if stat, err := os.Stat(f.Name); err == nil {
// 			if stat.IsDir() {
// 				filepath.Walk(f.Name, func(path string, finfo os.FileInfo, err error) error {
// 					if finfo.IsDir() {
// 						return nil
// 					}
// 					return backupFunc(path, finfo, f.Ignores)
// 				})
// 			} else {
// 				backupFunc(f.Name, stat, f.Ignores)
// 			}
// 		}
// 	}
// }
