package main

import (
	"fmt"
	"os"
	"uniq/read_write"
)

func main() {

	flags, fnameIn, fnameOut, err := read_write.GetFlags()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	data, err := read_write.ReadFile(fnameIn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = read_write.WriteFile(data, fnameOut)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("flags", flags)

	os.Exit(0)
}
