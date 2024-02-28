/*

String Formatting

Go offers excellent support for string formatting in the printf tradition. Here are some examples of common string formatting tasks.

type point struct {
	x, y int
}

var f = fmt.Printf

func main() {

	Go offers several printing "verbs" designed to format general Go values. For example, this prints an instance of our point struct.

	p := point{1, 2}
	f("struct1: %v\n", p)

	If the value is a struct, the %+v variant will include the struct's field names.

	f("struct2: %+v\n", p)

	The %#v variant prints a Go syntax representation of the value, i.e. the source code snippet that would produce that value.

	f("struct3: %#v\n", p)

	To print the type of a value, use %T.

	f("type: %T\n", p)

	Formatting booleans is straight-forward.

	f("bool: %t\n", true)

	There are many options for formatting integers. Use %d for standard, base-10 formatting.

	f("int: %d\n", 123)

	This prints a binary representation.

	f("binary: %b\n", 14)

	This prints the character corresponding to the given integer.

	f("char %c\n", 33)

	%x provides hex encoding.

	f("hex: %x\n", 456)

	There are also several formatting options for floats. For basic decimal formatting use %f.

	f("float1: %f\n", 78.9)

	%e and %E format the float in (slightly different versions of) scientific notation.

	f("format2: %e\n", 123400000.0)
	f("format3: %E\n", 123400000.0)

	For basic string printing use %s.

	f("str1: %s\n", "\"string\"")

	To double-quote strings as in Go source, use %q.

	f("str1: %q\n", "\"string\"")

	As with integers seen earlier, %x renders the string in base-16, with two output characters per byte of input.

	f("str3: %x\n", "hex this")

	To print a representation of a pointer, use %p.

	f("pointer: %p\n", &p)

	When formatting numbers you will often want to control the width and precision of the resulting figure. To specify the width of an integer, use number after the in the verb. By default the result will be right-justified and padded with spaces.

	f("width1: |%6d|%6d|\n", 12, 345)

	You can also specify the width of printed floats, though usually you'll also want to restrict the decimal precision at the same time with the width.precision syntax.

	f("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	To left-justify, use the - flag.

	f("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	You may also want to control width when formatting strings, especially to ensure that they align in table-like output. For basic right-justified width.

	f("width4: |%6s|%6s|\n", "foo", "b")

	To left-justify, use the - flag as with numbers.

	f("width5: |%-6s|%-6s|\n", "foo", "b")

	So far we've seen Printf, which prints the formatted string to os.Stdout. Sprintf formats and returns a string without printing it anywhere.

	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	You can format+print to io.Writers other than os.Stdout using Fprintf.

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")

}

*/

package main

import (
	"fmt"
	"os"
)

// String Formatting

// Go offers excellent support for string formatting in the printf tradition. Here are some examples of common string formatting tasks.

type point struct {
	x, y int
}

var f = fmt.Printf

func main() {

	// Go offers several printing "verbs" designed to format general Go values. For example, this prints an instance of our point struct.

	p := point{1, 2}
	f("struct1: %v\n", p)

	// If the value is a struct, the %+v variant will include the struct's field names.

	f("struct2: %+v\n", p)

	// The %#v variant prints a Go syntax representation of the value, i.e. the source code snippet that would produce that value.

	f("struct3: %#v\n", p)

	// To print the type of a value, use %T.

	f("type: %T\n", p)

	// Formatting booleans is straight-forward.

	f("bool: %t\n", true)

	// There are many options for formatting integers. Use %d for standard, base-10 formatting.

	f("int: %d\n", 123)

	// This prints a binary representation.

	f("binary: %b\n", 14)

	// This prints the character corresponding to the given integer.

	f("char %c\n", 33)

	// %x provides hex encoding.

	f("hex: %x\n", 456)

	// There are also several formatting options for floats. For basic decimal formatting use %f.

	f("float1: %f\n", 78.9)

	// %e and %E format the float in (slightly different versions of) scientific notation.

	f("format2: %e\n", 123400000.0)
	f("format3: %E\n", 123400000.0)

	// For basic string printing use %s.

	f("str1: %s\n", "\"string\"")

	// To double-quote strings as in Go source, use %q.

	f("str1: %q\n", "\"string\"")

	// As with integers seen earlier, %x renders the string in base-16, with two output characters per byte of input.

	f("str3: %x\n", "hex this")

	// To print a representation of a pointer, use %p.

	f("pointer: %p\n", &p)

	// When formatting numbers you will often want to control the width and precision of the resulting figure. To specify the width of an integer, use number after the in the verb. By default the result will be right-justified and padded with spaces.

	f("width1: |%6d|%6d|\n", 12, 345)

	// You can also specify the width of printed floats, though usually you'll also want to restrict the decimal precision at the same time with the width.precision syntax.

	f("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	// To left-justify, use the - flag.

	f("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// You may also want to control width when formatting strings, especially to ensure that they align in table-like output. For basic right-justified width.

	f("width4: |%6s|%6s|\n", "foo", "b")

	// To left-justify, use the - flag as with numbers.

	f("width5: |%-6s|%-6s|\n", "foo", "b")

	// So far we've seen Printf, which prints the formatted string to os.Stdout. Sprintf formats and returns a string without printing it anywhere.

	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	// You can format+print to io.Writers other than os.Stdout using Fprintf.

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")

}
