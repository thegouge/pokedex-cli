package main

import (
	"fmt"
	"os"
)

func commandExit(config *config, args []string) error {
	fmt.Println("goodbye!")
	os.Exit(0)
	return nil
}
