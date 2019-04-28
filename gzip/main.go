package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"time"
)

var zipoffset map[string]int64

func Zip(Src string, Dst string, SrcOffset int64) error {
	var DstFile *os.File

	if err := os.Chdir(path.Dir(Src)); err != nil {
		panic(err)
	}

	SrcName := path.Base(Src)
	SrcFile, err := os.Open(SrcName)
	defer SrcFile.Close()
	if err != nil {
		return err
	}

	if _, err := SrcFile.Seek(SrcOffset, 0); err != nil {
		return err
	}

	if SrcOffset == 0 {
		DstFile, _ = os.Create(Dst)
	} else {
		DstFile, _ = os.OpenFile(path.Base(Dst), os.O_APPEND|os.O_RDWR, 0)
	}
	defer DstFile.Close()
	zip := gzip.NewWriter(DstFile)
	defer zip.Close()

	fstat, _ := SrcFile.Stat()
	zip.Name = fstat.Name()
	zip.ModTime = fstat.ModTime()

	if _, err := io.Copy(zip, SrcFile); err != nil {
		return err
	} else {
		zip.Flush()
		cur, _ := SrcFile.Seek(0, os.SEEK_CUR)
		fmt.Println(cur)
		zipoffset[SrcName] = cur
		return nil
	}
	return nil
}

func encode(file string) {
	b, _ := json.Marshal(&zipoffset)
	f, _ := os.Create(file)
	defer f.Close()
	f.Write(b)
}

func decode(file string) {
	f, _ := os.Open(file)
	buf := make([]byte, 1024)
	n,_ := f.Read(buf)
	fmt.Println(buf)
	err:=json.Unmarshal(buf[:n], &zipoffset)
	if err != nil {
		panic(err)
	}
}

func main() {
	zipoffset = make(map[string]int64)
	SrcName := "D:\\err.log"
	err := Zip(SrcName, "D:\\err.gz", 0)
	encode("d:\\key.json")
	fmt.Println(err)
	time.Sleep(time.Millisecond * 15000)
	decode("d:\\key.json")
	err2 := Zip(SrcName, "D:\\err.gz", zipoffset[SrcName])
	fmt.Println(err2)
	//err := Zip()
}
