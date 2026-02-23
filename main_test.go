package main

import (
	"bytes"
	"os"
	"testing"
)

func TestHello(t *testing.T) {
	got := hello()
	want := "Hello go"
	if got != want {
		t.Fatalf("hello() = %q; want %q", got, want)
	}
}

func TestMainOutput(t *testing.T) {
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("pipe: %v", err)
	}
	os.Stdout = w

	main()

	w.Close()
	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)
	os.Stdout = old
	if err != nil {
		t.Fatalf("read: %v", err)
	}

	got := buf.String()
	want := "Hello go\n"
	if got != want {
		t.Fatalf("main output = %q; want %q", got, want)
	}
}
