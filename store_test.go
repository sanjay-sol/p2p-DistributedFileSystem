package main

import (
	"bytes"
	"testing"
)

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: DefaultPathTransformFunc,
	}

	s := NewStore(opts)

	data := bytes.NewReader(([]byte)("hello world"))
	if err := s.writeStream("key", data); err != nil {
		t.Error(err)
	}
}

