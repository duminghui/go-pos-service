// Package dstcmd provides ...
package dstcmd

import (
	"github.com/francoispqt/gojay"
)

// info is cmd getinfo
type info struct {
	Balance float64 `json:"balance"`
	Stake   float64 `json:"stake"`
	Blocks  int     `json:"blocks"`
}

func (i *info) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "balance":
		return dec.Float(&i.Balance)
	case "stake":
		return dec.Float(&i.Stake)
	case "blocks":
		return dec.Int(&i.Blocks)
	}
	return nil
}

func (i *info) NKeys() int {
	return 0
}

// GetInfo method
func GetInfo() (*info, error) {
	bytes, _ := execDstShell("getinfo")
	info_ := &info{}
	err := gojay.UnmarshalJSONObject(*bytes, info_)
	return info_, err
}
