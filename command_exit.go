package main

import (
	"fmt"
	"os"
)

func commandExit(config *config) error {
	fmt.Println("goodbye!")
	os.Exit(0)
	return nil
}
