package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk("/home/charles/GOPATH/src/deployfromgo", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
