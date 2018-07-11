// Package dstcmd provides ...
package dstcmd

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type TxNotExistError string

func (t TxNotExistError) Error() string {
	return string(t)
}

func execDstShell(s string) (*[]byte, error) {
	cmdStr := "/root/dstrad " + s
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	var out bytes.Buffer
	var errOut bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errOut
	err := cmd.Run()
	if err != nil {
		errInfo := errOut.String()
		if strings.Contains(errInfo, `"code":-5,"message":"No information available about transaction"`) {
			return nil, TxNotExistError(errInfo)
		}
		panic(fmt.Sprintf("ExecDstShell Error:[%s],[%s],[%s]", cmdStr, err, errInfo))
	}
	bytes := out.Bytes()
	return &bytes, err
}
