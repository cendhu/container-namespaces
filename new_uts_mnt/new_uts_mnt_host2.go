package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	fmt.Printf("Running parent %v \n", os.Args)

	cmd := exec.Command("./setup_host2", os.Args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWNS,
	}

	panicOnErr(cmd.Run())
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
