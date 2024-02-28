/*

Interfaces are named collections of method signatures.

Here's a basic interface for geometric shapes.

type geometry interface {
	area() float64
	perimeter() float64
}

For our example we'll implement this interface on
rectangle and circle types.

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

To implement an interface in Go, we just need to
implement all the methods in the interface. Here
we implement geometry on rectangles.

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

The implementation for circles.

func (c circle) area() float64 {
	return math.Pi * c.radius * c. radius
}

func (c circle) perimeter() float64 {
	return 2*math.Pi * c.radius
}

If a variable has an interface type, then we can call methods that
are in the named interface.
Here's a generic measure function taking advantage of this to work
on any geometry.

func measure(g geometry) {
	fmt.Printlon(g)
	fmt.Printlon(g.area())
	fmt.Printlon(g.perimeter())
}

func main() {
	r := rectangle{width: 3, height: 4}
	var c = circle{radius: 5}

	measure(r)
	measure(c)
}

*/

package main

import (
	"fmt"
	"math"
)

// Interfaces are named collections of method signatures.

// Here's a basic interface for geometric shapes.

type geometry interface {
	area() float64
	perimeter() float64
}

// For our example we'll implement this interface on
// rectangle and circle types.

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

// To implement an interface in Go, we just need to
// implement all the methods in the interface. Here
// we implement geometry on rectangles.

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}

// The implementation for circles.

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// If a variable has an interface type, then we can call methods that
// are in the named interface.
// Here's a generic measure function taking advantage of this to work
// on any geometry.

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	r := rectangle{width: 3, height: 4}
	var c = circle{radius: 5}

	measure(r)
	measure(c)
}
