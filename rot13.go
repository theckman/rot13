// rot13 utility
// this was used as a short task to help learn Golang -- really shouldn't be taken seriously
//
// released under the MIT license
// Copyright 2014 Tim Heckman <t@heckman.io>

package main

import (
	"fmt"
	"github.com/theckman/rot13/rot13"
	"os"
)

func main() {
	var b byte
	var ok bool

	// create a channel for goroutine communication
	c := make(chan byte)

	// crate the goroutine for rotter()
	go rot13.Rotter(os.Stdin, c)

	// forever loop and print any values sent through the channel
	// break the loop on cahnnel close
	for {
		b, ok = <-c

		if !ok {
			break
		}

		fmt.Print(string(b))
	}
}
