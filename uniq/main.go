package main

import (
	"flag"
	"fmt"
	"os"
	"uniq/read_write"
)

func main() {

	flags, output, err := read_write.GetFlags(os.Args[0], os.Args[1:])
	if err == flag.ErrHelp {
		fmt.Println(output)
		os.Exit(2)
	} else if err != nil {
		fmt.Println("got error:", err)
		fmt.Println("output:\n", output)
		flag.CommandLine.PrintDefaults()
		os.Exit(1)
	}
	data, err := read_write.ReadFile(flags.FNameIn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = read_write.WriteFile(data, flags.FNameOut)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("flags", flags)

	os.Exit(0)
}
