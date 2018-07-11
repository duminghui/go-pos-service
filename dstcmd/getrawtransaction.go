// Package dstcmd provides ...
package dstcmd

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/francoispqt/gojay"
)

type transaction struct {
	Txid          string `json:"txid"`
	Vin           vins   `json:"vin"`
	Vout          vouts  `json:"vout"`
	Blocktime     int    `json:"blocktime"`
	Confirmations int    `json:"confirmations"`
}

func (t *transaction) String() string {
	if t == nil {
		return "Transaction:<nil>"
	}
	buf := bytes.Buffer{}
	buf.WriteString("Transaction:\n  txid:")
	buf.WriteString(t.Txid)
	buf.WriteString(", Blocktime:")
	buf.WriteString(strconv.Itoa(t.Blocktime))
	buf.WriteString(", Confirmations:")
	buf.WriteString(strconv.Itoa(t.Confirmations))
	vinLen := len(t.Vin)
	buf.WriteString("\n  vin:[")
	buf.WriteString(strconv.Itoa(vinLen))
	buf.WriteString("]\n")
	for i, vin := range t.Vin {
		buf.WriteString(fmt.Sprintf("%*s%+v", 4, " ", vin))
		if i+1 != vinLen {
			buf.WriteString("\n")
		}
	}
	voutLen := len(t.Vout)
	buf.WriteString("\n  vout:[")
	buf.WriteString(strconv.Itoa(voutLen))
	buf.WriteString("]\n")
	if voutLen < 4 {
		for i, vout := range t.Vout {
			buf.WriteString(fmt.Sprintf("%*s%+v", 4, " ", vout))
			if i+1 != voutLen {
				buf.WriteString("\n")
			}
		}
	}
	return buf.String()

}

func (t *transaction) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "txid":
		return dec.String(&t.Txid)
	case "vin":
		t.Vin = make(vins, 0, 5)
		if err := dec.Array(&t.Vin); err != nil {
			return err
		}
	case "vout":
		t.Vout = make(vouts, 0, 5)
		if err := dec.Array(&t.Vout); err != nil {
			return err
		}
	case "blocktime":
		return dec.Int(&t.Blocktime)
	case "confirmations":
		return dec.Int(&t.Confirmations)
	}
	return nil

}

func (t *transaction) NKeys() int {
	return 0
}

type vin struct {
	Txid string `json:"txid"`
	Vout int    `json:"vout"`
}

// func (v *vin) String() string {
// 	if v == nil {
// 		return "Vin:<nil>"
// 	}
// 	buf := bytes.Buffer{}
// 	buf.WriteString("Vin:{Txid:")
// 	buf.WriteString(v.Txid)
// 	buf.WriteString(", Vout:")
// 	buf.WriteString(strconv.Itoa(v.Vout))
// 	buf.WriteString("}")
// 	return buf.String()
// }

func (v *vin) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "txid":
		return dec.String(&v.Txid)
	case "vout":
		return dec.Int(&v.Vout)
	}
	return nil
}

func (v *vin) NKeys() int {
	return 0
}

type vins []*vin

func (v *vins) UnmarshalJSONArray(dec *gojay.Decoder) error {
	vinTmp := new(vin)
	if err := dec.Object(vinTmp); err != nil {
		return err
	}
	*v = append(*v, vinTmp)
	return nil
}

type vout struct {
	Value        float64       `json:"value"`
	N            int           `json:"n"`
	ScriptPubKey *scriptPubKey `json:"scriptPubKey"`
}

func (v *vout) String() string {
	if v == nil {
		return "Vout:<nil>"
	}
	buf := bytes.Buffer{}
	buf.WriteString("Vout:{Value:")
	buf.WriteString(strconv.FormatFloat(v.Value, 'f', -1, 64))
	buf.WriteString(", N:")
	buf.WriteString(strconv.Itoa(v.N))
	buf.WriteString(",")
	buf.WriteString(fmt.Sprintf("%+v", v.ScriptPubKey))
	buf.WriteString("}")
	return buf.String()
}

func (v *vout) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "value":
		return dec.Float(&v.Value)
	case "n":
		return dec.Int(&v.N)
	case "scriptPubKey":
		v.ScriptPubKey = new(scriptPubKey)
		if err := dec.Object(v.ScriptPubKey); err != nil {
			return err
		}
	}
	return nil
}

func (v *vout) NKeys() int {
	return 0
}

type vouts []*vout

func (v *vouts) UnmarshalJSONArray(dec *gojay.Decoder) error {
	voutTmp := new(vout)
	if err := dec.Object(voutTmp); err != nil {
		return err
	}
	*v = append(*v, voutTmp)
	return nil
}

type scriptPubKey struct {
	Addresses addresses `json:"addresses"`
}

// func (s *scriptPubKey) String() string {
// 	if s == nil {
// 		return "ScriptPubKey:<nil>"
// 	}
// 	buf := bytes.Buffer{}
// 	buf.WriteString("ScriptPubKey:{")
// 	buf.WriteString(fmt.Sprintf("%s", s.Addresses))
// 	buf.WriteString("}")
// 	return buf.String()
// }

func (s *scriptPubKey) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "addresses":
		s.Addresses = make(addresses, 0, 1)
		if err := dec.Array(&s.Addresses); err != nil {
			return err
		}
	}
	return nil
}

func (s *scriptPubKey) NKeys() int {
	return 0
}

type addresses []string

// func (a *addresses) String() string {
// 	if a == nil {
// 		return "Addresses:<nil>"
// 	}
// 	return fmt.Sprintf("%#v", a)
// }

func (a *addresses) UnmarshalJSONArray(dec *gojay.Decoder) error {
	address := ""
	if err := dec.String(&address); err != nil {
		return err
	}
	*a = append(*a, address)
	return nil
}

func GetRawTransaction(txid string) (*transaction, error) {
	tx := new(transaction)
	cmd := fmt.Sprintf("getrawtransaction %s 1", txid)
	bytes, err := execDstShell(cmd)
	if err != nil {
		return nil, err
	}
	if err := gojay.UnmarshalJSONObject(*bytes, tx); err != nil {
		return nil, err
	}
	return tx, nil
}
