package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	cmd := exec.Command("/bin/sh", "/home/moxa/Erik/Study-Note/GO/GO_ScriptEnc/tt1")
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "inputEnv=true")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(out))
}
