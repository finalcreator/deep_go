package main

import (
	"fmt"
	"hash/crc64"
	"math/rand"

	"deepgou.com/google_deep_go/example/loadBalancer/smartbalancer"
)

type HashBalance struct {
}

func init() {
	smartbalancer.RegisterBalancer("hash", &HashBalance{})
}

func (p *HashBalance) DoBalance(insts []*smartbalancer.Instance, key ...string) (inst *smartbalancer.Instance, err error) {
	var defKey string = fmt.Sprintf("%d", rand.Int())
	if len(key) > 0 {
		defKey = key[0]
	}

	lens := len(insts)
	if lens == 0 {
		err = fmt.Errorf("No backend instance")
		return
	}

	// crcTable := crc32.MakeTable(crc32.IEEE)
	// hashVal := crc32.Checksum([]byte(defKey), crcTable)
	crcTable := crc64.MakeTable(crc64.ECMA)
	hashVal := crc64.Checksum([]byte(defKey), crcTable)
	index := hashVal % uint64(lens)
	fmt.Printf("hash:=%d, lens=%d, index=%d\n", hashVal, lens, index)
	inst = insts[index]

	return
}
