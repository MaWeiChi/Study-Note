package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func compress() {

	// dir, err := ioutil.TempDir("", "appman-configuration")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// defer os.RemoveAll(dir)
	dir := "./"
	fw, err := os.Create(filepath.Join(dir, "test.tar.gz"))
	if err != nil {
		fmt.Println(err.Error())
	}

	gzw := gzip.NewWriter(fw)
	defer gzw.Close()
	tw := tar.NewWriter(gzw)
	defer tw.Close()
	type Folder struct {
		Name    string
		Ignores []string
	}
	folders := []Folder{
		Folder{
			Name:    filepath.Join(DataDir, "env"),
			Ignores: []string{},
		},
		Folder{
			Name:    filepath.Join(DataDir, "sdcard"),
			Ignores: []string{},
		},
		Folder{
			filepath.Join(DataDir, "certs"),
			[]string{},
		},
		{
			filepath.Join(DataDir, "nginx"),
			[]string{},
		},
		{
			filepath.Join(BaseDir, "deviceinfo"),
			[]string{},
		},
		{
			JoinBaseDir("device"),
			[]string{},
		},
		{
			filepath.Join(DataDir, "redis"),
			[]string{},
		},
	}

	if dirs, err := ioutil.ReadDir(AppsDir); err == nil {
		for _, d := range dirs {
			if d.IsDir() {
				folders = append(
					folders,
					Folder{
						filepath.Join(AppsDir, d.Name(), "data"),
						[]string{"moxaenv"},
					},
				)
			}
		}
	} else {
		fmt.Println(err.Error())
	}
	fmt.Println(folders)
	// TODO ignore folder
	backupFunc := func(path string, finfo os.FileInfo, ignores []string) error {

		fout, err := os.Open(path)
		if err != nil {
			return nil
		}
		zipFilename := path[len(BaseDir)+1:]

		fh, _ := tar.FileInfoHeader(finfo, path)
		fh.Name = zipFilename
		if err := tw.WriteHeader(fh); err == nil {
			io.Copy(tw, fout)
		}

		fout.Close()
		return nil
	}
	for _, f := range folders {

		if stat, err := os.Stat(f.Name); err == nil {
			if stat.IsDir() {
				filepath.Walk(f.Name, func(path string, finfo os.FileInfo, err error) error {
					if finfo.IsDir() {
						return nil
					}
					return backupFunc(path, finfo, f.Ignores)
				})
			} else {
				backupFunc(f.Name, stat, f.Ignores)
			}
		}
	}
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
