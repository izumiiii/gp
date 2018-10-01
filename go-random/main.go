package main

import (
	"crypto/rand"
	"encoding/binary"
	"strconv"
)

func main() {
	println(random())
	println(random())
	println(random())
}

func random() string {
	var n uint64
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return strconv.FormatUint(n, 36)
}
