// A simple terminal program which connects to a named pipe (on windows).
// Handy for serial console on Hyper-V VMs.
// Build with:
// GOOS=windows GOARCH=amd64 go build
//
// Use:
// - npterm <pipe name>
//
package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"github.com/Azure/go-ansiterm/winterm"
	"github.com/Microsoft/go-winio"
)

var pipeName = "\\\\.\\pipe\\dockerMobyLinuxVM-com1"

// Some of the code below is copied and modified from:
// https://github.com/moby/moby/blob/master/pkg/term/term_windows.go
const (
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms683167(v=vs.85).aspx
	enableVirtualTerminalInput      = 0x0200
	enableVirtualTerminalProcessing = 0x0004
	disableNewlineAutoReturn        = 0x0008
)

func main() {
	filePtr := flag.String("file", "", "Optionally log the output to a file")
	flag.Parse()

	remArgs := flag.Args()
	if len(remArgs) != 0 {
		pipeName = remArgs[0]
	}
	fmt.Println("Connecting to:", string(pipeName))

	out := io.MultiWriter(os.Stdout)
	if *filePtr != "" {
		fmt.Println("Logging to:", *filePtr)
		if f, err := os.Create(*filePtr); err == nil {
			out = io.MultiWriter(os.Stdout, f)
		} else {
			panic(err)
		}
	}

	if err := hypervConfigureConsole(); err != nil {
		fmt.Printf("Configure Console: %v\n", err)

	}
	defer hypervRestoreConsole()

	var c net.Conn
	var err error
	for count := 1; count < 500; count++ {
		c, err = winio.DialPipe(pipeName, nil)
		if err != nil {
			// Argh, different Windows versions seem to
			// return different errors and we can't easily
			// catch the error. On some versions it is
			// winio.ErrTimeout...
			// Instead poll 100 times and then error out
			fmt.Printf("Connect to console: %v\n", err)
			time.Sleep(10 * 1000 * 1000 * time.Nanosecond)
			continue
		}
		break
	}
	if err != nil {
		fmt.Printf("Connect to console: %v\n", err)
		return
	}
	defer c.Close()

	fmt.Println("Connected")
	go io.Copy(c, os.Stdin)

	_, err = io.Copy(out, c)
	if err != nil {
		fmt.Printf("Copy to console: %v\n", err)
	}

}

var (
	hypervStdinMode  uint32
	hypervStdoutMode uint32
	hypervStderrMode uint32
)

func hypervConfigureConsole() error {
	// Turn on VT handling on all std handles, if possible. This might
	// fail on older windows version, but we'll ignore that for now
	// Also disable local echo

	fd := os.Stdin.Fd()
	if hypervStdinMode, err := winterm.GetConsoleMode(fd); err == nil {
		if err = winterm.SetConsoleMode(fd, hypervStdinMode|enableVirtualTerminalInput); err != nil {
			fmt.Println("VT Processing is not supported on stdin")

		}
	}

	fd = os.Stdout.Fd()
	if hypervStdoutMode, err := winterm.GetConsoleMode(fd); err == nil {
		if err = winterm.SetConsoleMode(fd, hypervStdoutMode|enableVirtualTerminalProcessing|disableNewlineAutoReturn); err != nil {
			fmt.Println("VT Processing is not supported on stdout")
		}
	}

	fd = os.Stderr.Fd()
	if hypervStderrMode, err := winterm.GetConsoleMode(fd); err == nil {
		if err = winterm.SetConsoleMode(fd, hypervStderrMode|enableVirtualTerminalProcessing|disableNewlineAutoReturn); err != nil {
			fmt.Println("VT Processing is not supported on stderr")
		}
	}
	return nil
}

func hypervRestoreConsole() {
	winterm.SetConsoleMode(os.Stdin.Fd(), hypervStdinMode)
	winterm.SetConsoleMode(os.Stdout.Fd(), hypervStdoutMode)
	winterm.SetConsoleMode(os.Stderr.Fd(), hypervStderrMode)
}
