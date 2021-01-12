package main

import (
	"bytes"
	"context"
	"os/exec"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
)

// Commander ...
type Commander struct {
	Logger *log.Logger
}

type CommanderInterface interface {
	Log() *log.Logger
	// RunAgentCommand(command ...string) (string, int, error)
	Run(logEnable bool, command ...string) (string, int, error)
	RunWithCancel(ctx context.Context, timeout int, logEnable bool, command ...string) (string, int)
	RunWithTimeout(timeout int, logEnable bool, command ...string) (string, int)
}

func CommanderNew(log *log.Logger) *Commander {
	return &Commander{
		log,
	}
}

func (cmd *Commander) Log() *log.Logger {
	return cmd.Logger
}

// Run : run system command.
func (cmd *Commander) Run(logEnable bool, command ...string) (string, int, error) {
	cmd.Logger.Println("exec:", command)

	head := command[0]
	parts := command[1:]
	out, err := exec.Command(head, parts...).CombinedOutput()
	if err != nil {
		exitCode := -1
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			exitCode = ws.ExitStatus()
		}
		cmd.Logger.Printf("exec return code : %d,\nerror : %s\noutput : %s\n", exitCode, err, string(out))
		return string(out), exitCode, err
	}

	if logEnable && len(out) > 0 {
		cmd.Logger.Debug("exec output : \n%s\n", string(out))
	}
	return string(out), 0, nil
}

// RunWithCancel : run system command with exit cancel context.
func (cmd *Commander) RunWithCancel(ctx context.Context, timeout int, logEnable bool, command ...string) (string, int) {
	cmd.Logger.Println("exec:", command)

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

	done := make(chan error)
	go func() { done <- execmd.Wait() }()

	select {
	case <-ctx.Done():
		execmd.Process.Kill()
		cmd.Logger.Println("Command exit force.")
	case err := <-done:
		if logEnable {
			cmd.Logger.Debug("stdout:", outbuf.String())
			cmd.Logger.Debug("stderr:", errbuf.String())
		}
		if err != nil {
			exitCode := -1
			if exitError, ok := err.(*exec.ExitError); ok {
				ws := exitError.Sys().(syscall.WaitStatus)
				exitCode = ws.ExitStatus()
			}
			cmd.Logger.Println("Non-zero exit code:", exitCode, err.Error())
			cmd.Logger.Println("stderr:", errbuf.String())
			return outbuf.String(), exitCode
		}
	}
	return outbuf.String(), 0
}

// RunWithTimeout : run system command with timeout.
func (cmd *Commander) RunWithTimeout(timeout int, logEnable bool, command ...string) (string, int) {
	cmd.Logger.Println("exec:", command)

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

	err := execmd.Wait()
	if logEnable {
		cmd.Logger.Debug("stdout:", outbuf.String())
		cmd.Logger.Debug("stderr:", errbuf.String())
	}
	if err != nil {
		exitCode := -1
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			exitCode = ws.ExitStatus()
		}
		cmd.Logger.Println("Non-zero exit code:", err.Error())
		cmd.Logger.Println("stderr:", errbuf.String())
		return outbuf.String(), exitCode
	}
	return outbuf.String(), 0
}

// Run : run system command.
func (cmd *Commander) RunWithoutPrintln(logEnable bool, command ...string) (string, int, error) {
	//cmd.Logger.Println("exec:", command)

	head := command[0]
	parts := command[1:]
	out, err := exec.Command(head, parts...).CombinedOutput()
	if err != nil {
		exitCode := -1
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			exitCode = ws.ExitStatus()
		}
		cmd.Logger.Printf("exec return code : %d,\nerror : %s\noutput : %s\n", exitCode, err, string(out))
		return string(out), exitCode, err
	}

	if logEnable && len(out) > 0 {
		cmd.Logger.Debug("exec output : \n%s\n", string(out))
	}
	return string(out), 0, nil
}

// func (cmd *Commander) RunAgentCommand(command ...string) (string, int, error) {
// 	httpc := http.Client{
// 		Transport: &http.Transport{
// 			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
// 				return net.Dial("unix", "/host/var/run/system-agent.sock")
// 			},
// 		},
// 	}
// 	var response *http.Response
// 	var err error

// 	payloag, _ := sjson.Set("{}", "data", command)
// 	response, err = httpc.Post("http://unix/api/v1/system/command", "charset=utf-8", strings.NewReader(payloag))
// 	if err != nil {
// 		cmd.Logger.Println("system-agent command run failed, error:", err.Error())
// 		return "", 1, err
// 	}
// 	defer httpc.CloseIdleConnections()
// 	defer response.Body.Close()
// 	responseOutput, resReadErr := ioutil.ReadAll(response.Body)
// 	responseStr := string(responseOutput)
// 	if resReadErr != nil {
// 		cmd.Logger.Printf("system-agent command run failed, status code: %d, \noutput:%s, \nerror:%v\n", response.StatusCode, responseStr, resReadErr)
// 		return responseStr, response.StatusCode, resReadErr
// 	}
// 	if response.StatusCode < 200 || response.StatusCode >= 300 {
// 		ret := 1
// 		msg := responseStr
// 		if retJson := gjson.Get(responseStr, "return"); retJson.Exists() {
// 			ret = int(retJson.Int())
// 		}
// 		if msgJson := gjson.Get(responseStr, "message"); msgJson.Exists() {
// 			msg = msgJson.String()
// 		}
// 		return responseStr, ret, errors.New(msg)
// 	}
// 	cmd.Logger.Println("status code:", response.StatusCode)
// 	cmd.Logger.Println("output:", responseStr)
// 	return responseStr, 0, nil
// }
