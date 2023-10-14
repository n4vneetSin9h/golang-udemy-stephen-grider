package main

import (
	"fmt"
	"io"
	"os"
)

// type fileWriter struct {}

// func (fileWriter) Write(bs []byte) (int, error) {
// 	return len(bs), nil
// }

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: printfile filename")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	
	if err != nil {
		fmt.Println("Something went wrong: ", err)
		os.Exit(1)
	}

	// fW := fileWriter{}

	// FIXME
	io.Copy(os.Stdout, file)
}