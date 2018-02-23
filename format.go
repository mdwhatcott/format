// package format is basically the fmt package, but spelled correctly, without
// the AWOL characters, the ambiguities, and the archaic ambiance.
package format

import (
	"fmt"
	"io"
)

/////////////////////////////////////////////////////////////////////////////////

// Error -> fmt.Errorf
func Error(format string, items ...interface{}) error {
	return fmt.Errorf(format, items...)
}

/////////////////////////////////////////////////////////////////////////////////

// Write -> fmt.Fprint
func Write(writer io.Writer, items ...interface{}) (written int, err error) {
	return fmt.Fprint(writer, items...)
}

// WriteFormat -> fmt.Fprintf
func WriteFormat(writer io.Writer, format string, items ...interface{}) (written int, err error) {
	return fmt.Fprintf(writer, format, items...)
}

// WriteLine -> fmt.Fprintln
func WriteLine(writer io.Writer, items ...interface{}) (written int, err error) {
	return fmt.Fprintln(writer, items...)
}

/////////////////////////////////////////////////////////////////////////////////

// Read -> fmt.Fscan
func Read(reader io.Reader, items ...interface{}) (elements int, err error) {
	return fmt.Fscan(reader, items...)
}

// ReadFormat -> fmt.Fscanf
func ReadFormat(reader io.Reader, format string, items ...interface{}) (elements int, err error) {
	return fmt.Fscanf(reader, format, items...)
}

// ReadLine -> Fscanln
func ReadLine(reader io.Reader, items ...interface{}) (elements int, err error) {
	return fmt.Fscanln(reader, items...)
}

/////////////////////////////////////////////////////////////////////////////////

// Print -> fmt.Print (Hey, what a coincidence!)
func Print(items ...interface{}) (written int, err error) {
	return fmt.Print(items...)
}

// PrintFormat -> fmt.Printf
func PrintFormat(format string, items ...interface{}) (written int, err error) {
	return fmt.Printf(format, items...)
}

// PrintLine -> fmt.Println
func PrintLine(items ...interface{}) (written int, err error) {
	return fmt.Println(items...)
}

/////////////////////////////////////////////////////////////////////////////////

// Scan -> fmt.Scan (Hey, what a coincidence!)
func Scan(items ...interface{}) (elements int, err error) {
	return fmt.Scan(items...)
}

// ScanFormat -> fmt.Scanf
func ScanFormat(format string, items ...interface{}) (elements int, err error) {
	return fmt.Scanf(format, items...)
}

// ScanLine -> fmt.Scanln
func ScanLine(items ...interface{}) (elements int, err error) {
	return fmt.Scanln(items...)
}

/////////////////////////////////////////////////////////////////////////////////

// Compose -> fmt.Sprint
func Compose(items ...interface{}) string {
	return fmt.Sprint(items...)
}

// ComposeFormat -> fmt.Sprintf
func ComposeFormat(format string, items ...interface{}) string {
	return fmt.Sprintf(format, items...)
}

// ComposeLine -> fmt.Sprintln
func ComposeLine(items ...interface{}) string {
	return fmt.Sprintln(items...)
}

/////////////////////////////////////////////////////////////////////////////////

// Decompose -> fmt.Sscan
func Decompose(value string, items ...interface{}) (elements int, err error) {
	return fmt.Sscan(value, items...)
}

// DecomposeFormat -> fmt.Sscanf
func DecomposeFormat(value string, format string, items ...interface{}) (elements int, err error) {
	return fmt.Sscanf(value, format, items...)
}

// DecomposeLine -> fmt.Sscanln
func DecomposeLine(value string, items ...interface{}) (elements int, err error) {
	return fmt.Sscanln(value, items...)
}

/////////////////////////////////////////////////////////////////////////////////
