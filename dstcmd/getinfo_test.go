package dstcmd

import (
	"fmt"
	"testing"
)

func TestGetInfo(t *testing.T) {
	t.SkipNow()
	info, _ := GetInfo()
	fmt.Printf("%#v\n", info)
}
