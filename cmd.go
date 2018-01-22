package urun

import (
	"bytes"
	"os/exec"
	"strings"
)

func CmdExec(cmdName string, cmdArgs ...string) (stdout string, stderr string, err error) {
	return CmdExecStdin("", "", cmdName, cmdArgs...)
}

func CmdExecIn(dir string, cmdName string, cmdArgs ...string) (stdout string, stderr string, err error) {
	return CmdExecStdin("", dir, cmdName, cmdArgs...)
}

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
