package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// Create a zip under pwd dir

	// contents := []byte("Hello World")
	// fzip, err := os.Create(`./test.zip`)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// zipw := zip.NewWriter(fzip)
	// defer zipw.Close()
	// w, err := zipw.Encrypt(`test.txt`, `golang`)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _, err = io.Copy(w, bytes.NewReader(contents))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// zipw.Flush()

	type Folder struct {
		Name    string
		Ignores []string
	}

	folders := []Folder{
		Folder{
			Name:    filepath.Join(config_DataDir, "env"),
			Ignores: []string{},
		},
		Folder{
			Name:    filepath.Join(config_DataDir, "sdcard"),
			Ignores: []string{},
		},
		Folder{
			filepath.Join(config_DataDir, "certs"),
			[]string{},
		},
		{
			filepath.Join(config_DataDir, "nginx"),
			[]string{},
		},
		{
			filepath.Join(config, "deviceinfo"),
			[]string{},
		},
		{
			JoinBaseDir("device"),
			[]string{},
		},
		{
			filepath.Join(config_DataDir, "redis"),
			[]string{},
		},
	}

	fmt.Println(folders)
}

const (
	config         string = "/host/var/thingspro"
	config_DataDir string = "/host/var/thingspro/data"
	config_AppsDir string = "/host/var/thingspro/app"
)

func JoinBaseDir(dirs ...string) string {
	return filepath.Join(config, filepath.Join(dirs...))
}
