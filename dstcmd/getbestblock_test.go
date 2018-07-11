package dstcmd

import (
	"fmt"
	"testing"
)

func TestGetBestBlock(t *testing.T) {
	t.SkipNow()
	block, _ := GetBestBlock()
	fmt.Printf("%#v\n", block)
}
