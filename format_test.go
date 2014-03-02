package format

import (
	"bytes"
	"fmt"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFormatFunctions(t *testing.T) {
	var (
		formatWriter *bytes.Buffer
		fmtWriter    *bytes.Buffer
	)

	Convey("Output:", t, func() {

		Convey("Writing...", func() {
			formatWriter = bytes.NewBufferString("")
			fmtWriter = bytes.NewBufferString("")

			Convey("Write: items to a writer", func() {
				actualWritten, actualError := Write(formatWriter, 1, 2, 3)
				expectedWritten, expectedError := fmt.Fprint(fmtWriter, 1, 2, 3)

				So(actualWritten, ShouldEqual, expectedWritten)
				So(actualError, ShouldEqual, expectedError)
				So(formatWriter.String(), ShouldEqual, fmtWriter.String())
			})

			Convey("WriteFormat: formatted items to a writer", func() {
				actualWritten, actualError := WriteFormat(formatWriter, "%d %d %d", 1, 2, 3)
				expectedWritten, expectedError := fmt.Fprintf(fmtWriter, "%d %d %d", 1, 2, 3)

				So(actualWritten, ShouldEqual, expectedWritten)
				So(actualError, ShouldEqual, expectedError)
				So(formatWriter.String(), ShouldEqual, fmtWriter.String())
			})

			Convey("WriteLine: items as a line", func() {
				actualWritten, actualError := WriteLine(formatWriter, 1, 2, 3)
				expectedWritten, expectedError := fmt.Fprintln(fmtWriter, 1, 2, 3)

				So(actualWritten, ShouldEqual, expectedWritten)
				So(actualError, ShouldEqual, expectedError)
				So(formatWriter.String(), ShouldEqual, fmtWriter.String())
			})
		})

		Convey("Printing...", func() {
			// TODO: swap os.Stdout with os.Pipe:
			// http://stackoverflow.com/questions/10473800/in-go-how-do-i-capture-stdout-of-a-function-into-a-string?rq=1

			Convey("items to standard out", func() {
				actualWritten, actualError := Print(1, 2, 3, "\n")
				expectedWritten, expectedError := fmt.Print(1, 2, 3, "\n")

				So(actualWritten, ShouldEqual, expectedWritten)
				So(actualError, ShouldEqual, expectedError)
				// TODO: assert what was written is equal
			})

			Convey("formatted items to standard out", nil)

			Convey("items as a line to standard out", nil)

			Reset(func() {
				// TODO
			})
		})

		Convey("Composing...", func() {

			Convey("items as a string", nil)

			Convey("formatted items as a string", nil)

			Convey("items as a newline-terminated string", nil)
		})
	})

	Convey("Input:", t, func() {

		Convey("Reading...", func() {

			Convey("space-delimited items from a reader", nil)

			Convey("formatted, space-delimited items from a reader", nil)

			Convey("space-delimited items in a line from a reader", nil)
		})

		Convey("Scanning...", func() {

			Convey("space-delimited items from standard in", nil)

			Convey("formatted, space-delimited items from standard in", nil)

			Convey("space-delimited items in a line from standard in", nil)
		})

		Convey("Decomposing...", func() {

			Convey("space-delimited items from a string", nil)

			Convey("formatted, space-delimited items from a string", nil)

			Convey("space-delimited items in a newline-terminated string", nil)
		})
	})
}
