// Package dstcmd provides ...
package dstcmd

import (
	"github.com/francoispqt/gojay"
)

// block getblock
type block struct {
	Time int `json:"time"`
}

func (b *block) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "time":
		return dec.Int(&b.Time)
	}
	return nil
}

func (b *block) NKeys() int {
	return 8
}

func GetBestBlock() (*block, error) {
	bestblockhash, _ := execDstShell("getbestblockhash")
	// fmt.Printf("TTTT:%s\n", *bestblockhash)
	blockInfo, _ := execDstShell("getblock " + string(*bestblockhash))
	block_ := &block{}
	err := gojay.UnmarshalJSONObject(*blockInfo, block_)
	return block_, err
}
