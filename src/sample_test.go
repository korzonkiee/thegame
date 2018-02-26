package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	res := add(2, 2)
	if res != 4 {
		t.Error("Error")
	}
}
