package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fname := "" // no file name -> open error

	if len(os.Args) > 1 {
		fname = os.Args[1]
	}

	f, err := os.Open(fname)

	if err != nil {
		fmt.Fprintln(os.Stderr, "bad file: ", err)
		os.Exit(-1)
	}

	data := make([]byte, 100)
	total := 0

	for {
		n, err := f.Read(data)

		if err == io.EOF {
			break
		}

		total += n
		fmt.Println("read", n, "chars")
	}

	fmt.Printf("The file has %d bytes\n", total)
}
