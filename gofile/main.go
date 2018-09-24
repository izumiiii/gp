package main

import (
	"fmt"
	"os"
)

func main() {
	// defaultUmask := syscall.Umask(0)
	// os.Mkdir("/tmp/go-dir1", 0777)
	// syscall.Umask(defaultUmask)

	for i := 1; i <= 10; i++ {
		os.MkdirAll(fmt.Sprintf("/tmp/go-dirs/%02d", i), os.ModePerm)
	}
}
