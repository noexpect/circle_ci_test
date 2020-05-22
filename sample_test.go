package main

import (
	"testing"
)

func TestExistArgs(t *testing.T) {
	input := []string{"a", "b"}
	result := check_args(input)
	if result != "a" {
		t.Fatal("failed test")
	}
}

func TestNoArgs(t *testing.T) {
	input := []string{}
	result := check_args(input)
	if result != "no args" {
		t.Fatal("failed test")
	}
}
