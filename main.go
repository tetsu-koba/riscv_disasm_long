package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func rewriteDisasm(file string, debug bool) {
	var instream io.Reader
	instream = os.Stdin
	if file != "-" {
		cmd := exec.Command("llvm-objdump", "-d", file)
		in, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}
		defer in.Close()
		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}
		instream = in
	}

	scanner := bufio.NewScanner(instream)
	scanner.Split(bufio.ScanLines)

	// Assume the lines have the following format
	// <address and hex dump>\t<opcode>\t<operand>...
	for scanner.Scan() {
		line := scanner.Text()
		if debug {
			fmt.Println(line)
		}
		groups := strings.Split(line, "\t")
		if len(groups) < 2 {
			fmt.Println(line)
			continue
		}
		t := opcodeMap[groups[1]]
		if len(t) > 0 {
			groups[1] = t
		}
		if (len(groups[1]) < 8) {
			groups[1] += "\t"
		}
		output := groups[0]
		for _, s := range groups[1:] {
			output += ("\t" + s)
		}
		fmt.Println(output)
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Println("Usage:")
		log.Printf("\t%s objfile  (\"llvm-objdump -d objfile\" is called internally)\n", os.Args[0])
		log.Printf("\t%s - < objdump_output\n", os.Args[0])
		os.Exit(1)
	}
	debug := false
	rewriteDisasm(os.Args[1], debug)
}
