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

var configurationDir = "/home/moxa/study-note/GO/GO_Tar&Gz/exact"

func untargz() {

	os.RemoveAll(configurationDir)
	os.MkdirAll(configurationDir, 0755)
	// extract
	fname := "/home/moxa/study-note/GO/GO_Tar&Gz/2020-06-18-14-45-backup.tar.gz"
	r, err := os.Open(fname)
	if err != nil {

	}
	gr, err := gzip.NewReader(r)
	if err != nil {

	}
	tr := tar.NewReader(gr)

	extractDir := filepath.Join(configurationDir, filepath.Base(fname)[:len(filepath.Base(fname))-7])
	os.RemoveAll(extractDir)
	os.MkdirAll(extractDir, 0755)
	for {
		f, err := tr.Next()
		if err == io.EOF {
			fmt.Println("exact end")
			break
		}

		os.MkdirAll(filepath.Join(extractDir, filepath.Dir(f.Name)), 0755)
		data := make([]byte, f.Size)
		_, err = tr.Read(data)
		if err != nil && err != io.EOF {
			fmt.Println(err.Error())

		}
		ioutil.WriteFile(filepath.Join(extractDir, f.Name), data, 0755)

	}
	r.Close()
}
