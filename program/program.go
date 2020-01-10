package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Printf("Running parent %v \n", os.Args)

	cmd := exec.Command("./setup", os.Args[1:]...)
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
