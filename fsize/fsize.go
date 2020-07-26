package main

import (
	"fmt"
	"io/ioutil"
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

	d, err := ioutil.ReadAll(f)

	if err != nil {
		fmt.Fprintln(os.Stderr, "can't read: ", err)
		os.Exit(-1)
	}

	fmt.Printf("The file has %d bytes\n", len(d))
}
