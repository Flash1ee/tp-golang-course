package main

import (
	"flag"
	"fmt"
	"os"
	"uniq/io"
	"uniq/uniq"
)

func main() {
	flags, output, err := io.GetFlags(os.Args[0], os.Args[1:])
	if err == flag.ErrHelp {
		fmt.Println(output)
		os.Exit(2)
	} else if err != nil {
		fmt.Println("got error:", err)
		fmt.Println("output:\n", output)
		flag.CommandLine.PrintDefaults()
		os.Exit(1)
	}
	data, err := io.ReadFile(flags.FNameIn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	res, err := uniq.Uniq(data, flags)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = io.WriteFile(res, flags)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// fmt.Println("flags", flags)

	os.Exit(0)
}
