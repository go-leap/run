# urun
--
    import "github.com/go-leap/run"


## Usage

#### func  CmdExec

```go
func CmdExec(cmdName string, cmdArgs ...string) (stdout string, stderr string, err error)
```
CmdExecStdin runs the specified program, returning its full `stdout` response
and whitespace-trimmed `stderr` response.

#### func  CmdExecIn

```go
func CmdExecIn(dir string, cmdName string, cmdArgs ...string) (stdout string, stderr string, err error)
```
CmdExecStdin runs the specified program in the specified `dir` (if any),
returning its full `stdout` response and whitespace-trimmed `stderr` response.

#### func  CmdExecStdin

```go
func CmdExecStdin(stdin string, dir string, cmdName string, cmdArgs ...string) (stdout string, stderr string, err error)
```
CmdExecStdin runs the specified program with the specified `stdin` (if any),
returning its full `stdout` response and whitespace-trimmed `stderr` response.

#### func  CmdTryStart

```go
func CmdTryStart(cmdname string, cmdargs ...string) (err error)
```
CmdTryStart attempts to run the specified program, records whether this resulted
in an `error`, and immediately kills the program before returning.

#### func  CmdsTryStart

```go
func CmdsTryStart(cmds map[string]*CmdTry)
```
CmdsTryStart calls `CmdTryStart` for all specified `cmds` in parallel.

#### func  SetupIpcPipes

```go
func SetupIpcPipes(bufferCapacity int, ipcSplitProtocol bufio.SplitFunc, needJsonOut bool) (stdin *bufio.Scanner, rawOut *bufio.Writer, jsonOut *json.Encoder)
```
SetupIpcPipes sets up IPC pipes with the specified `bufferCapacity`. If
`ipcSplitProtocol` is `nil`, a readline-writeline ping-pong protocol is assumed.
Only if `needJsonOut`, is `jsonOut` allocated via `json.NewEncoder(rawOut)`.

#### type CmdTry

```go
type CmdTry struct {
	Args []string
	Ran  *bool
}
```

CmdTry is used for `CmdsTryStart` to probe in parallel the installation status
of multiple programs.
