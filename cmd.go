package urun

import (
	"bytes"
	"os/exec"
	"strings"
	"sync"
)

// CmdTry is used for `CmdsTryStart` to probe in parallel the installation status of multiple programs.
type CmdTry struct {
	Args []string
	Ran  *bool
}

// CmdTryStart attempts to run the specified program, records whether this resulted in an `error`, and immediately kills the program before returning.
func CmdTryStart(cmdname string, cmdargs ...string) (err error) {
	cmd := exec.Command(cmdname, cmdargs...)
	err = cmd.Start()
	defer cmd.Wait()
	if cmd.Process != nil {
		_ = cmd.Process.Kill()
	}
	return
}

// CmdsTryStart calls `CmdTryStart` for all specified `cmds` in parallel.
func CmdsTryStart(cmds map[string]*CmdTry) {
	var w sync.WaitGroup
	run := func(cmd string, try *CmdTry) {
		defer w.Done()
		*try.Ran = (nil == CmdTryStart(cmd, try.Args...))
	}
	w.Add(len(cmds))
	for cmdname, cmdmore := range cmds {
		go run(cmdname, cmdmore)
	}
	w.Wait()
}

// CmdExecStdin runs the specified program,
// returning its full `stdout` response and whitespace-trimmed `stderr` response.
func CmdExec(cmdName string, cmdArgs ...string) (stdout string, stderr string, err error) {
	return CmdExecStdin("", "", cmdName, cmdArgs...)
}

// CmdExecStdin runs the specified program in the specified `dir` (if any),
// returning its full `stdout` response and whitespace-trimmed `stderr` response.
func CmdExecIn(dir string, cmdName string, cmdArgs ...string) (stdout string, stderr string, err error) {
	return CmdExecStdin("", dir, cmdName, cmdArgs...)
}

// CmdExecStdin runs the specified program with the specified `stdin` (if any),
// returning its full `stdout` response and whitespace-trimmed `stderr` response.
func CmdExecStdin(stdin string, dir string, cmdName string, cmdArgs ...string) (stdout string, stderr string, err error) {
	if cmdName != "" && strings.Contains(cmdName, " ") && len(cmdArgs) == 0 {
		cmdArgs = strings.Split(cmdName, " ")
		cmdName, cmdArgs = cmdArgs[0], cmdArgs[1:]
	}
	cmd := exec.Command(cmdName, cmdArgs...)
	if cmd.Dir = dir; stdin != "" {
		cmd.Stdin = strings.NewReader(stdin)
	}
	var bufout, buferr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &bufout, &buferr
	if err = cmd.Run(); err != nil {
		if _, isexiterr := err.(*exec.ExitError); isexiterr || strings.Contains(err.Error(), "pipe has been ended") || strings.Contains(err.Error(), "pipe has been closed") {
			err = nil
		}
	}
	stdout, stderr = bufout.String(), strings.TrimSpace(buferr.String())
	return
}
