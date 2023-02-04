package main

import (
	"errors"
	"testing"
)

func TestVar(t *testing.T) {
	var a int64
	btn := true
	if btn {
		a, err := 2, errors.New("13sai")
		t.Log("btn", a, err)
	}

	t.Log(a)
}
