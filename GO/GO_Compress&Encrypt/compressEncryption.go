// package main

// import (
// 	"archive/tar"
// 	"compress/gzip"
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"os"
// 	"path/filepath"
// )

// func compressAndEncrypt() {
// 	fw, err := os.Create("testReadytoEncrypt.tar.gz")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	gzw := gzip.NewWriter(fw)
// 	tw := tar.NewWriter(gzw)

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
// 	//fmt.Println(folders)
// 	// TODO ignore folder
// 	backupFunc := func(path string, finfo os.FileInfo, ignores []string) error {

// 		fout, err := os.Open(path)
// 		if err != nil {
// 			fmt.Println(err.Error())

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
// 	tw.Close()
// 	gzw.Close()
// 	fw.Close()
// 	a, _ := ioutil.ReadFile("2020-06-19-17-42-backup.tar.gz")
// 	encryptfile("testEncrypt.tar.gz", a, "12345678")
// }

// // func unCompressEncrypt() {
// // 	data := decryptFile("2020-08-04-11-05-backup.tar.gz", "fdaf1@123")
// // 	// if err != nil {
// // 	// 	fmt.Println(err.Error())
// // 	// 	return
// // 	// }

// // 	ioutil.WriteFile("Encrypt.tar.gz", data, 0755)
// // 	r := bytes.NewReader(data)
// // 	os.RemoveAll(ExtractEnDir)
// // 	os.MkdirAll(ExtractEnDir, 0755)
// // 	gr, err := gzip.NewReader(r)
// // 	if err != nil {
// // 		fmt.Println(err.Error())
// // 	}
// // 	tr := tar.NewReader(gr)

// // 	for {
// // 		f, err := tr.Next()
// // 		if err == io.EOF {
// // 			break
// // 		}
// // 		if f.FileInfo().IsDir() {
// // 			continue
// // 		}

// // 		os.MkdirAll(filepath.Join(ExtractEnDir, filepath.Dir(f.Name)), 0755)
// // 		data := make([]byte, f.Size)
// // 		_, err = tr.Read(data)
// // 		if err != nil && err != io.EOF {
// // 			fmt.Println(err.Error())
// // 		}

// // 		ioutil.WriteFile(filepath.Join(ExtractEnDir, f.Name), data, 0755)

// // 	}
// // 	gr.Close()
// // 	fmt.Println("extractDir: " + ExtractEnDir)

// // }
