package dstcmd

import (
	"fmt"
	"testing"
)

func TestExecShell(t *testing.T) {
	t.SkipNow()
	str, err := execDstShell("getinfo")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(*str))
}
