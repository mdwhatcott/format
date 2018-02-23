package main

import (
	"bytes"

	. "github.com/mdwhatcott/format"
)

var hello, world string

func main() {
	errors()
	print()
	write()
	compose()
	read()
	scan()
	decompose()
}

func errors() {
	PrintLine("\nError functions...")
	PrintLine("==================")
	PrintLine(Error("This is a '%s, %s!' error.", "Hello", "World"))
}

func print() {
	PrintLine("\nPrint functions...")
	PrintLine("==================")
	Print("Hello, World!\n")
	PrintFormat("Hello, %s!\n", "World")
	PrintLine("Hello, World!")
}

func write() {
	PrintLine("\nWrite functions...")
	PrintLine("==================")
	writer := bytes.NewBufferString("")
	Write(writer, "1")
	WriteFormat(writer, "%d", 2)
	WriteLine(writer, "3")
	Print(writer.String())
}

func compose() {
	PrintLine("\nCompose functions...")
	PrintLine("====================")
	Print(Compose(1, 2, 3, "\n"))
	Print(ComposeFormat("%d %d %d\n", 1, 2, 3))
	Print(ComposeLine(1, 2, 3))
}

func read() {
	PrintLine("\nRead functions...")
	PrintLine("=================")
	Read(bytes.NewBufferString("Hello, World!"), &hello, &world)
	PrintLine(hello, world)
	ReadFormat(bytes.NewBufferString("Goodbye world"), "%s %s", &hello, &world)
	PrintLine(hello, world)
	ReadLine(bytes.NewBufferString("hi world\n"), &hello, &world)
	PrintLine(hello, world)
}

func scan() {
	PrintLine("\nScan functions... (enter two words, separated by white space, then <ENTER>)")
	PrintLine("=================")
	Print("> ")
	Scan(&hello, &world)
	PrintLine(hello, world)
	Print("> ")
	ScanLine(&hello, &world)
	PrintLine(hello, world)
	Print("> ")
	ScanFormat("%s %s\n", &hello, &world)
	PrintLine(hello, world)
}

func decompose() {
	PrintLine("\nDecompose functions...")
	PrintLine("======================")
	Decompose("blah blah", &hello, &world)
	PrintLine(hello, world)
	DecomposeFormat("BLAH BLAH", "%s %s", &hello, &world)
	PrintLine(hello, world)
	DecomposeLine("BlAh bLaH\n", &hello, &world)
	PrintLine(hello, world)
}
