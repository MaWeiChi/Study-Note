package main

import (
	"os"
	"path/filepath"
)

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
