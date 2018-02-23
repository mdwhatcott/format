// package format is basically the fmt package, but spelled correctly, without
// the AWOL characters, the ambiguities, and the archaic ambiance.
package format

import "fmt"

var (
	Error = fmt.Errorf

	Write       = fmt.Fprint
	WriteFormat = fmt.Fprintf
	WriteLine   = fmt.Fprintln

	Read       = fmt.Fscan
	ReadFormat = fmt.Fscanf
	ReadLine   = fmt.Fscanln

	Print       = fmt.Print
	PrintFormat = fmt.Printf
	PrintLine   = fmt.Println

	Scan       = fmt.Scan
	ScanFormat = fmt.Scanf
	ScanLine   = fmt.Scanln

	Compose       = fmt.Sprint
	ComposeFormat = fmt.Sprintf
	ComposeLine   = fmt.Sprintln

	Decompose       = fmt.Sscan
	DecomposeFormat = fmt.Sscanf
	DecomposeLine   = fmt.Sscanln
)
