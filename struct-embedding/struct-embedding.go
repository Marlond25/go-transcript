/*

Go supports embedding of structs and interfaces to express
a more seameless composition of types.
This is not to be confused with //go:embed which is a go
directive introduced in Go version 1.16+ to embed files and
folders into the application binary.

type base struct {
	number int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.number)
}

A container embeds a base. An embedding looks
like a field without a name.

type container struct {
	base
	str string
}

func main() {

	When creating structs with literals, we have to initialize
	the embedding explicitly; here the embedded type serves
	as the field name.

	co := container{
		base: base{number: 1},
		str: "some name",
	}

	We can access the base's fields directly on co, e.g.
	co.number.

	fmt.Printf("co={num: %v, str: %v}\n", co.number, co.str)

	Alternatively, we can spell out the full path using
	the embedded type name.

	fmt.Printf("also num:", co.base.num)

	Since container embeds base, the methods of base
	also become methods of a container. Here we
	invoke a method that was embedded from base
	directly on co.

	fmt.Printf("describe:", co.describe())

	type describer interface {
		describe() string
	}

	Embedding structs with methods may be used to
	bestow interface implementations onto other
	structs. Here we see that a container now
	implements the describer interface because it
	embeds base.

	var d describer = co
	fmt.Println("describer:", d.describe())

}

*/

package main

import "fmt"

// Go supports embedding of structs and interfaces to express
// a more seameless composition of types.
// This is not to be confused with //go:embed which is a go
// directive introduced in Go version 1.16+ to embed files and
// folders into the application binary.

type base struct {
	number int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.number)
}

// A container embeds a base. An embedding looks
// like a field without a name.

type container struct {
	base
	str string
}

func main() {

	// When creating structs with literals, we have to initialize
	// the embedding explicitly; here the embedded type serves
	// as the field name.

	co := container{
		base: base{number: 1},
		str:  "some name",
	}

	// We can access the base's fields directly on co, e.g.
	// co.number.

	fmt.Printf("co={num: %v, str: %v}\n", co.number, co.str)

	// Alternatively, we can spell out the full path using
	// the embedded type name.

	fmt.Println("also num:", co.base.number)

	// Since container embeds base, the methods of base
	// also become methods of a container. Here we
	// invoke a method that was embedded from base
	// directly on co.

	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// Embedding structs with methods may be used to
	// bestow interface implementations onto other
	// structs. Here we see that a container now
	// implements the describer interface because it
	// embeds base.

	var d describer = co
	fmt.Println("describer:", d.describe())

}
