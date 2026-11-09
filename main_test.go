package main

import "testing"

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Fatal("Add is wrong")
	}
}

func TestSub(t *testing.T) {
	if Sub(3, 1) != 2 {
		t.Fatal("Sub is wrong")
	}
}
