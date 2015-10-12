package main

import (
	"errors"
	"reflect"
	"testing"
)

type expectData struct {
	in  string
	out string
	err error
}

var expectsDest = []expectData{
	{"", "000", nil},
	{"M", "001", nil},
	{"D", "010", nil},
	{"MD", "011", nil},
	{"A", "100", nil},
	{"AM", "101", nil},
	{"AD", "110", nil},
	{"AMD", "111", nil},
	{"MA", "", errors.New("undefined dest mnemonic: MA")},
}

func TestCodeDest(t *testing.T) {
	for i := range expectsDest {
		test := &expectsDest[i]
		out, err := CodeDest(test.in)
		if test.out != out || !reflect.DeepEqual(test.err, err) {
			t.Errorf("CodeDest(%v) = %v, %v want %v, %v",
				test.in, out, err, test.out, test.err)
		}
	}
}

var expectsComp = []expectData{
	{"0", "0101010", nil},
	{"1", "0111111", nil},
	{"-1", "0111010", nil},
	{"D", "0001100", nil},
	{"A", "0110000", nil},
	{"!D", "0001101", nil},
	{"!A", "0110001", nil},
	{"-D", "0001111", nil},
	{"-A", "0110011", nil},
	{"D+1", "0011111", nil},
	{"A+1", "0110111", nil},
	{"D-1", "0001110", nil},
	{"A-1", "0110010", nil},
	{"D+A", "0000010", nil},
	{"D-A", "0010011", nil},
	{"A-D", "0000111", nil},
	{"D&A", "0000000", nil},
	{"D|A", "0010101", nil},
	{"M", "1110000", nil},
	{"!M", "1110001", nil},
	{"-M", "1110011", nil},
	{"M+1", "1110111", nil},
	{"M-1", "1110010", nil},
	{"D+M", "1000010", nil},
	{"D-M", "1010011", nil},
	{"M-D", "1000111", nil},
	{"D&M", "1000000", nil},
	{"D|M", "1010101", nil},
	{"M&D", "", errors.New("undefined comp mnemonic: M&D")},
}

func TestCodeComp(t *testing.T) {
	for i := range expectsComp {
		test := &expectsComp[i]
		out, err := CodeComp(test.in)
		if test.out != out || !reflect.DeepEqual(test.err, err) {
			t.Errorf("CodeComp(%v) = %v, %v want %v, %v",
				test.in, out, err, test.out, test.err)
		}
	}
}

var expectsJump = []expectData{
	{"", "000", nil},
	{"JGT", "001", nil},
	{"JEQ", "010", nil},
	{"JGE", "011", nil},
	{"JLT", "100", nil},
	{"JNE", "101", nil},
	{"JLE", "110", nil},
	{"JMP", "111", nil},
	{"JUM", "", errors.New("undefined jump mnemonic: JUM")},
}

func TestCodeJump(t *testing.T) {
	for i := range expectsJump {
		test := &expectsJump[i]
		out, err := CodeJump(test.in)
		if test.out != out || !reflect.DeepEqual(test.err, err) {
			t.Errorf("CodeJump(%v) = %v, %v want %v, %v",
				test.in, out, err, test.out, test.err)
		}
	}
}
