/*

func main() {

	A go string is a read-only slice of bytes. The
	language and the standard library treat strings
	specially - as containers of text encoded in UTF-8.
	In other language, strings are made of
	"characters". In Go, the concept of a character is
	called a rune - it's an integer that represents a Unicode code point.

	s is a string assigned a literal value representing
	the word "hello" in the Thai language. Go string literals are UTF-8
	encoded text.

	const s = "Hello"

	Since strings are equivalent to []byte, this will produce the length
	of the raw bytes stored within.

	fmt.Println("Len:", len(s))

	Indexing into a string produces the raw byte values at each index.
	This loop generates the hex values of all the bytes that constitute
	the code points in s.

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	To cound how many runes are in a string, we can
	use the utf8 package. Note that the run-time of
	RuneCountInString depends on the size of the
	string, because it has to decode each UTF-8 rune
	sequentially. Some Thai characters are
	represented by multiple UTF-8 code points, so the
	result of this count may be surprising.

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	A range loop handles strings specially and decodes
	each rune along with its offset in the string.

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	We can achieve the same iteration by using the
	utf8.DecodeRuneInString function explicitly.

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		examineRune(runeValue)
	}

	This demonstrates passing a rune value to a function.

}

Values enclosed in single quotes are rune literals.
We can compare a rune value to a rune literal directly.

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == "h" {
		fmt.Println("fount eich")
	}
}

*/

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// A go string is a read-only slice of bytes. The
	// language and the standard library treat strings
	// specially - as containers of text encoded in UTF-8.
	// In other language, strings are made of
	// "characters". In Go, the concept of a character is
	// called a rune - it's an integer that represents a Unicode code point.

	// s is a string assigned a literal value representing
	// the word "hello" in the Thai language. Go string literals are UTF-8
	// encoded text.

	const s = "Hello"

	// Since strings are equivalent to []byte, this will produce the length
	// of the raw bytes stored within.

	fmt.Println("Len:", len(s))

	// Indexing into a string produces the raw byte values at each index.
	// This loop generates the hex values of all the bytes that constitute
	// the code points in s.

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
		if s[i] == 'l' {
			fmt.Println("This was treated as a rune", s[i])
		}
	}
	fmt.Println()

	// To cound how many runes are in a string, we can
	// use the utf8 package. Note that the run-time of
	// RuneCountInString depends on the size of the
	// string, because it has to decode each UTF-8 rune
	// sequentially. Some Thai characters are
	// represented by multiple UTF-8 code points, so the
	// result of this count may be surprising.

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// A range loop handles strings specially and decodes
	// each rune along with its offset in the string.

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// We can achieve the same iteration by using the
	// utf8.DecodeRuneInString function explicitly.

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		examineRune(runeValue)
	}

	// This demonstrates passing a rune value to a function.

}

// Values enclosed in single quotes are rune literals.
// We can compare a rune value to a rune literal directly.

func examineRune(r rune) {
	if r == 'l' {
		fmt.Println("found el")
	} else if r == 'H' {
		fmt.Println("fount eich")
	}
}
