package main

import (
	"fmt"
	"uniq/read_write"
)

func main() {

	flags, fnameIn, fnameOut, err := read_write.GetFlags()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("flags", flags)
	fmt.Println("input file", fnameIn)
	fmt.Println("out file", fnameOut)

}