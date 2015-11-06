package main

import (
	"testing"
)

func TestSymbolTable(t *testing.T) {
	s := NewSymbolTable()

	s.AddEntry("LOOP", 16)
	if !s.Contains("LOOP") {
		t.Fatalf("s must contains LOOP")
	}
	actual := s.GetAddress("LOOP")
	if actual != 16 {
		t.Fatalf("Expecting LOOP address is 16, but %v", actual)
	}
}
