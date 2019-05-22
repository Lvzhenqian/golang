package main

import (
	"gopkg.in/cheggaaa/pb.v1"
	"io"
)

func main() {
	// create and start bar
	bar := pb.New(myDataLen).SetUnits(pb.U_BYTES)
	bar.Start()

	// my io.Reader
	r := myReader

	// my io.Writer
	w := myWriter

	// create proxy reader
	reader := bar.NewProxyReader(r)

	// and copy from pb reader
	io.Copy(w, reader)

	// create and start bar
	bar := pb.New(myDataLen).SetUnits(pb.U_BYTES)
	bar.Start()

	// my io.Reader
	r := myReader

	// my io.Writer
	w := myWriter

	// create multi writer
	writer := io.MultiWriter(w, bar)

	// and copy
	io.Copy(writer, r)

	bar.Finish()
}
