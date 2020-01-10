package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// go run main.go run <cmd> <args>
func main() {
	fmt.Printf("Running child %v \n", os.Args)

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	panicOnErr(syscall.Sethostname([]byte("host2")))
	panicOnErr(syscall.Chroot("/home/senthil/lab/host1-go1.12.7"))
	panicOnErr(syscall.Setenv("GOROOT", "/root/go"))
	panicOnErr(syscall.Setenv("PATH", "/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/snap/bin:/root/go/bin"))
	panicOnErr(os.Chdir("/"))
	panicOnErr(syscall.Mount("proc", "proc", "proc", 0, ""))

	panicOnErr(cmd.Run())
	panicOnErr(syscall.Unmount("proc", 0))
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
