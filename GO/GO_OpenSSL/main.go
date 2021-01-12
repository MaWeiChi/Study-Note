package main

import (
	"bytes"
	"context"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strconv"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
)

var Logger *log.Logger

type certificateInfo struct {
	deviceID     string
	Org          string `json:"organization"`
	IssuerCN     string `json:"issuerCN"`
	NotBefore    string `json:"notBefore"`
	NotAfter     string `json:"notAfter"`
	ModelName    string `json:"modelName"`
	MacAddress   string `json:"macAddress"`
	SerialNumber string `json:"serialNumber"`
}

type personalInfo struct {
	ModelName    string `json:"model_name"`
	MacAddress   string `json:"mac"`
	SerialNumber string `json:"serial_number"`
}

var timeFormat = "Jan _2, 2006 15:04:05"

func tFormat(value string) string {
	if t, err := time.Parse("Jan 2 15:04:05 2006 MST", value); err == nil {
		return t.Format(timeFormat)
	}
	return value
}

func main() {
	var c certificateInfo
	certPath := "C:/dsc/enroll/dev.crt"
	out, rc := RunWithTimeout(60, []string{"openssl", "x509", "-in", certPath, "-noout", "-text"}...)
	if rc != 0 {
		fmt.Println(errors.New(out))
	}

	fileContent, _ := ioutil.ReadFile(certPath)
	block, _ := pem.Decode(fileContent)
	cert, _ := x509.ParseCertificate(block.Bytes)
	c.IssuerCN = cert.Issuer.CommonName
	c.NotBefore = cert.NotBefore.Format(timeFormat)
	c.NotAfter = cert.NotAfter.Format(timeFormat)
	c.deviceID = cert.Subject.CommonName
	c.Org = cert.Subject.Organization[0]
	regex := *regexp.MustCompile(`personalSignature = "([ \S]*)"`)
	if res := regex.FindAllStringSubmatch(out, 1); res != nil {
		s, _ := strconv.Unquote("\"" + res[0][1] + "\"")
		var p personalInfo
		json.Unmarshal([]byte(s), &p)
		c.MacAddress = p.MacAddress
		c.ModelName = p.ModelName
		c.SerialNumber = p.SerialNumber
	}

	fmt.Println(c.IssuerCN)
	fmt.Println(c.NotBefore)
	fmt.Println(c.NotAfter)
	fmt.Println(c.deviceID)
	fmt.Println(c.Org)
	fmt.Println(c.MacAddress)
	fmt.Println(c.ModelName)
	fmt.Println(c.SerialNumber)

}

// RunWithTimeout : run system command with timeout.
func RunWithTimeout(timeout int, command ...string) (string, int) {
	if timeout <= 0 {
		timeout = 30
	}
	ctxt, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	head := command[0]
	parts := command[1:]
	execmd := exec.CommandContext(ctxt, head, parts...)

	// Use a bytes.Buffer to get the output
	var outbuf bytes.Buffer
	var errbuf bytes.Buffer
	execmd.Stdout = &outbuf
	execmd.Stderr = &errbuf

	execmd.Start()

	if err := execmd.Wait(); err != nil {
		exitCode := -1
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			exitCode = ws.ExitStatus()
		}
		return fmt.Sprintf("command: %v\nNon-zero exit code: %s\nstderr: %s\nstdout: %s",
			command, err.Error(), errbuf.String(), outbuf.String()), exitCode
	}
	return outbuf.String(), 0
}
