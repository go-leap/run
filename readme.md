# urun
--
    import "github.com/go-leap/run"


## Usage

#### func  CmdExec

```go
func CmdExec(cmdName string, cmdArgs ...string) (stdout string, stderr string, err error)
```

#### func  CmdExecIn

```go
func CmdExecIn(dir string, cmdName string, cmdArgs ...string) (stdout string, stderr string, err error)
```

#### func  CmdExecStdin

```go
func CmdExecStdin(stdin string, dir string, cmdName string, cmdArgs ...string) (stdout string, stderr string, err error)
```

#### func  SetupIpcPipes

```go
func SetupIpcPipes(bufferCapacity int, ipcSplitProtocol bufio.SplitFunc, needJsonOut bool) (stdin *bufio.Scanner, rawOut *bufio.Writer, jsonOut *json.Encoder)
```
SetupIpcPipes sets up IPC pipes with the specified `bufferCapacity`. If
`ipcSplitProtocol` is `nil`, a readline-writeline ping-pong protocol is assumed.
Only if `needJsonOut`, is `jsonOut` allocated via `json.NewEncoder(rawOut)`.
