package rot13

import (
	"bytes"
	"fmt"
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

func neq_err(v1, v2, v3 byte) string {
	return fmt.Sprintf("rot13ing '%v' did not eql '%v', it equaled '%v'", string(v1), string(v2), string(v3))
}

func TestRot13(t *testing.T) {
	r13 := []byte("abcdefghijklmnopqrstuvwxyz!5")
	R13 := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ*4")
	e13 := []byte("nopqrstuvwxyzabcdefghijklm!5")
	E13 := []byte("NOPQRSTUVWXYZABCDEFGHIJKLM*4")

	for i, v := range r13 {
		if val := Rot13(v); val != e13[i] {
			t.Error(neq_err(v, val, e13[i]))
		}
	}

	for i, v := range R13 {
		if val := Rot13(v); val != E13[i] {
			t.Error(neq_err(v, val, E13[i]))
		}
	}
}
