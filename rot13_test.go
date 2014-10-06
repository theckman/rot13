package main

import (
	"bytes"
	"testing"
)

func TestRotter(t *testing.T) {
	content := bytes.NewReader([]byte("hello"))
	channel := make(chan byte)
	bytes := make([]byte, 0)

	go Rotter(content, channel)

	for {
		buffer, ok := <-channel
		if !ok {
			break
		}
		bytes = append(bytes, buffer)
	}

	if string(bytes) != "uryyb" {
		t.Error("")
	}
}

func TestRot13(t *testing.T) {
	r13 := []byte("abcdefghijklmnopqrstuvwxyz!5")
	R13 := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ*4")
	e13 := []byte("nopqrstuvwxyzabcdefghijklm!5")
	E13 := []byte("NOPQRSTUVWXYZABCDEFGHIJKLM*4")

	for i, v := range r13 {
		if Rot13(v) != e13[i] {
			t.Error("rot13ing", string(v), "did not equal", string(e13[i]))
		}
	}

	for i, v := range R13 {
		if Rot13(v) != E13[i] {
			t.Error("Oh yes")
		}
	}
}
