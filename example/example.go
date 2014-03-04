package main

import (
	"bytes"

	"github.com/mdwhatcott/format"
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
	format.PrintLine("\nError functions...")
	format.PrintLine("==================")
	format.PrintLine(format.Error("This is a '%s, %s!' error.", "Hello", "World"))
}

func print() {
	format.PrintLine("\nPrint functions...")
	format.PrintLine("==================")
	format.Print("Hello, World!\n")
	format.PrintFormat("Hello, %s!\n", "World")
	format.PrintLine("Hello, World!")
}

func write() {
	format.PrintLine("\nWrite functions...")
	format.PrintLine("==================")
	writer := bytes.NewBufferString("")
	format.Write(writer, "1")
	format.WriteFormat(writer, "%d", 2)
	format.WriteLine(writer, "3")
	format.Print(writer.String())
}

func compose() {
	format.PrintLine("\nCompose functions...")
	format.PrintLine("====================")
	format.Print(format.Compose(1, 2, 3, "\n"))
	format.Print(format.ComposeFormat("%d %d %d\n", 1, 2, 3))
	format.Print(format.ComposeLine(1, 2, 3))
}

func read() {
	format.PrintLine("\nRead functions...")
	format.PrintLine("=================")
	format.Read(bytes.NewBufferString("Hello, World!"), &hello, &world)
	format.PrintLine(hello, world)
	format.ReadFormat(bytes.NewBufferString("Goodbye world"), "%s %s", &hello, &world)
	format.PrintLine(hello, world)
	format.ReadLine(bytes.NewBufferString("hi world\n"), &hello, &world)
	format.PrintLine(hello, world)
}

func scan() {
	format.PrintLine("\nScan functions... (enter two words, separated by white space, then <ENTER>)")
	format.PrintLine("=================")
	format.Print("> ")
	format.Scan(&hello, &world)
	format.PrintLine(hello, world)
	format.Print("> ")
	format.ScanLine(&hello, &world)
	format.PrintLine(hello, world)
	format.Print("> ")
	format.ScanFormat("%s %s\n", &hello, &world)
	format.PrintLine(hello, world)
}

func decompose() {
	format.PrintLine("\nDecompose functions...")
	format.PrintLine("======================")
	format.Decompose("blah blah", &hello, &world)
	format.PrintLine(hello, world)
	format.DecomposeFormat("BLAH BLAH", "%s %s", &hello, &world)
	format.PrintLine(hello, world)
	format.DecomposeLine("BlAh bLaH\n", &hello, &world)
	format.PrintLine(hello, world)
}
