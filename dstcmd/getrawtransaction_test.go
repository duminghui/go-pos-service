package dstcmd

import (
	"fmt"
	"log"
	"testing"
)

func TestGetRawTransaction(t *testing.T) {
	t.SkipNow()
	tx, _ := GetRawTransaction("cbf13d5451579f2a0111b8988419e268c7ec571acfbf5d6776f645283cc8f31b")
	fmt.Printf("TX:%#v\n", tx)
}

func BenchmarkGetRawTransaction(b *testing.B) {
	b.SkipNow()
	for i := 0; i < b.N; i++ {
		tx, _ := GetRawTransaction("cbf13d5451579f2a0111b8988419e268c7ec571acfbf5d6776f645283cc8f31b")
		log.Printf("%s", tx)
	}
}
