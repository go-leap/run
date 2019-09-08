package urun

import (
	"bufio"
	"encoding/json"
	"os"
)

// SetupIpcPipes sets up IPC pipes with the specified `bufferCapacity`. If `ipcSplitProtocol` is `nil`, a readline-writeline ping-pong protocol is assumed. Only if `needJsonOut`, is `jsonOut` allocated via `json.NewEncoder(rawOut)`.
func SetupIpcPipes(bufferCapacity int, ipcSplitProtocol bufio.SplitFunc, needJsonOut bool) (stdin *bufio.Scanner, rawOut *bufio.Writer, jsonOut *json.Encoder) {
	stdin = bufio.NewScanner(os.Stdin)
	stdin.Buffer(make([]byte, bufferCapacity), 8+bufferCapacity /* max needs to be greater than cap to allow support for later growth-if-and-when-needed */)
	if ipcSplitProtocol != nil {
		stdin.Split(ipcSplitProtocol)
	}

	rawOut = bufio.NewWriterSize(os.Stdout, bufferCapacity)
	if needJsonOut {
		jsonOut = json.NewEncoder(rawOut)
		jsonOut.SetEscapeHTML(false)
		jsonOut.SetIndent("", "")
	}
	return
}
