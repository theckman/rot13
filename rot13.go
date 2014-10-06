package main

import (
	"bufio"
	"io"
)

// rot13 converter for a byte value
func Rot13(b byte) byte {
	// variables used to define the upper and lower bounds of the alphabet range
	var a, z byte

	// determine whether it's an upper or lowercase letter
	// if it's not a-zA-Z, return the byte as-is
	switch {
	case 'a' <= b && b <= 'z':
		a, z = 'a', 'z'
	case 'A' <= b && b <= 'Z':
		a, z = 'A', 'Z'
	default:
		return b
	}

	// convert the byte to its rot13 counterpart and return it
	return (b-a+13)%(z-a+1) + a
}

// go routine runner for sending stdin through ror13 conversion
func Rotter(io_r io.Reader, ch chan byte) {
	// buf is a single byte used for reading a single byte from stdin
	// err is whether or not we've hit an error reading the next byte (EOF)
	var buf byte
	var err error

	// create a new buffered IO reader for stdin
	r := bufio.NewReader(io_r)

	// loop until ReadByte() returns an error
	for {
		// read stdin one byte at a time
		buf, err = r.ReadByte()

		// if it errored, bail out
		if err != nil {
			break
		}

		// send the rot13 value of the last byte through the channel
		ch <- Rot13(buf)
	}

	// close the channel when finished to signal that we should exit
	close(ch)
}
