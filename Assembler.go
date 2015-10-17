package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: Assembler foo.asm")
		os.Exit(1)
	}
	filename := os.Args[1]

	if !strings.HasSuffix(filename, ".asm") {
		fmt.Println("Usage: Assembler foo.asm")
		os.Exit(2)
	}

	fp, err := os.Open(filename)
	if err != nil {
		fmt.Println("Err! The file cannot open")
		os.Exit(3)
	}

	scanner := bufio.NewScanner(fp)
	hackfile := strings.TrimSuffix(filename, ".asm") + ".hack"
	writerfp, err := os.OpenFile(hackfile, os.O_WRONLY | os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Err! .hack file couldn't create.")
		os.Exit(4)
	}
	writer := bufio.NewWriter(writerfp)

	for p := NewParser(scanner); p.HasMoreCommands(); p.Advance() {
		var output string
		switch p.CommandType() {
		case A_COMMAND:
			output = "0"
			//symbol := p.Symbol()
			// if number-string then transform binary
			// if symbol-table then transform address and transform binary

			fmt.Fprintln(writer, output)
		case C_COMMAND:
			output = "111"
			comp, err := CodeComp(p.Comp())
			if err != nil {
				os.Exit(5)
			}
			output += comp
			dest, err := CodeDest(p.Dest())
			if err != nil {
				os.Exit(5)
			}
			output += dest
			jump, err := CodeJump(p.Jump())
			if err != nil {
				os.Exit(5)
			}
			output += jump

			fmt.Fprintln(writer, output)
		case L_COMMAND:
			// do nothing
		}
		writer.Flush()
	}
}

