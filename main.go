package main

import (
	"fmt"
	"os"

	"langmak/cmd"
	//"Users/user/Documents/univeritmo/lang/micro-lang-master"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
