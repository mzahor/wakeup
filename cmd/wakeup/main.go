// 192.168.0.2 E0:D5:5E:A6:56:74
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mzahor/wakeup"
)

func main() {
	flag.Parse()
	ip := flag.Arg(0)
	mac := flag.Arg(1)
	if ip == "" || mac == "" {
		exitWithErr("Usage: wakeup <ip> <mac>\n", 1)
	}
	err := wakeup.WakeUp(ip, mac)
	if err != nil {
		if wErr, ok := err.(*wakeup.WakeupError); ok {
			exitWithErr(err.Error(), wErr.Code)
		} else {
			exitWithErr("Unknown error", 100)
		}
	}
}

func exitWithErr(msg string, code int) {
	fmt.Printf("Error: %s\n", msg)
	os.Exit(code)
}
