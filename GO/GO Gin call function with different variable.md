### GO: Gin call function with different variable

#### Original:

There are five similar functions that each one is different at deviceInfo. 

Detail code see the [Link](#sample)

First :

```go
func WriteFakeBothPversionToJson(c *gin.Context) {
	...
	var deviceInfo ModelAndTPEVersionInfo
	deviceInfo.Model = "uc-1100-lx"
	deviceInfo.TPEVersion = "2.0.0-1266"
	deviceInfoJSON, _ := json.MarshalIndent(deviceInfo, "", " ")
	...
}
```

Second:

```go
func WriteFakeDevicePversionToJson(c *gin.Context) {
	...
	deviceInfo := readPversion()
	deviceInfo.Model = "uc-1100-lx"
	deviceInfoJSON, _ := json.MarshalIndent(deviceInfo, "", " ")
	...
}
```

Third:

```go
func WriteFakeTPEVersionPversionToJson(c *gin.Context) {
	...
	deviceInfo := readPversion()
	deviceInfo.TPEVersion = "1.8.0-1266"
	deviceInfoJSON, _ := json.MarshalIndent(deviceInfo, "", " ")
	...
}
```

Fourth:

```go
func WriteTPELastDiffPversionToJson(c *gin.Context) {
	...
	deviceInfo := readPversion()
	deviceInfo.TPEVersion = "2.0.0-1266"
	deviceInfoJSON, _ := json.MarshalIndent(deviceInfo, "", " ")
	...
}
```

Fifth:

```go
func WriteBothUnknownPversionToJson(c *gin.Context) {
	...
	var deviceInfo ModelAndTPEVersionInfo
	deviceInfo.Model = "unknown"
	deviceInfo.TPEVersion = "unknown"
	deviceInfoJSON, _ := json.MarshalIndent(deviceInfo, "", " ")
	...
}
```



---



### Refine:

First:

```go
func WriteFakeBothPversionToJson(c *gin.Context) {
	writeFakePversionToJson(c, "uc-1100-lx", "2.0.0-1266")
}
```

Second:

```go
func WriteFakeDevicePversionToJson(c *gin.Context) {
    writeFakePversionToJson(c, "uc-1100-lx", "")
}
```

Third:

```go
func WriteFakeTPEVersionPversionToJson(c *gin.Context) {
    writeFakePversionToJson(c, "", "1.8.0-1266")
}
```

Fourth:

```go
func WriteTPELastDiffPversionToJson(c *gin.Context) {
	writeFakePversionToJson(c, "", "2.0.0-1266")
}
```

Fifth:

```go
func WriteBothUnknownPversionToJson(c *gin.Context) {
	writeFakePversionToJson(c, "unknown", "unknown")
}
```



#### Sample:

```go
func writeFakePversionToJson(c *gin.Context, Model string, TPEVersion string) {
	logrus.Infof("WriteTPELastDiffPversionToJson")
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	// Prepare chunk data
	t := time.Now()
	parameter := &SystemConfigurationParameter{
		Encryption: true,
		Password:   "123456",
		Filename:   fmt.Sprintf("%d-%02d-%02d-%02d-%02d-backup.zip", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute()),
	}
	if err := c.BindQuery(&parameter); err != nil {
		logrus.Warn(err)
	}
	c.Header("Content-Type", "application/zip")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", parameter.Filename))
	c.Writer.WriteHeader(http.StatusOK)

	dir, err := ioutil.TempDir("", "appman-configuration")
	if err != nil {
		logrus.Error(err)
	}
	defer os.RemoveAll(dir)
    
	// ... difference begin
    
	deviceInfo := readPversion()
	if Model != "" {
		deviceInfo.Model = Model
	}
	if TPEVersion != "" {
		deviceInfo.TPEVersion = TPEVersion
	}
	deviceInfoJSON, _ := json.MarshalIndent(deviceInfo, "", " ")
    
    // ... difference end
    
	err = ioutil.WriteFile(config.BaseDir+"/deviceinfo", deviceInfoJSON, 0644)
	if err != nil {
		logrus.Error(err)
	}
	defer os.Remove(config.BaseDir + "/deviceinfo")
	zipw := zip.NewWriter(c.Writer)
	defer zipw.Close()

	folders := []Folder{
		{
			filepath.Join(config.BaseDir, "deviceinfo"),
			[]string{},
		},
	}

	// TODO ignore folder
	backupFunc := func(path string, finfo os.FileInfo, ignores []string) error {
		logrus.Debugf("backup %s", path)
		if ctx.Err() != nil {
			return nil
		}
		fout, err := os.Open(path)
		if err != nil {
			logrus.Warnf("ignore %s: %s", path, err.Error())
			return nil
		}
		zipFilename := path[len(config.BaseDir)+1:]
		if parameter.Encryption {
			if w, err := zipw.EncryptWithFileInfoHeader(zipFilename, finfo, parameter.Password); err == nil {
				io.Copy(w, fout)
			}
		} else {
			fh, _ := zip.FileInfoHeader(finfo)
			fh.Name = zipFilename
			if w, err := zipw.CreateHeader(fh); err == nil {
				io.Copy(w, fout)
			}
		}
		fout.Close()
		return nil
	}

	for _, f := range folders {
		if ctx.Err() != nil {
			break
		}
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
```



