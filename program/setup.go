package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Printf("Running child %v \n", os.Args)

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	panicOnErr(cmd.Run())
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
