package main

import "fmt"

/*

Branching with if and else in Go is straightforward.

Here's a basic example.

if 7%2 == 0 {
    fmt.Println("7 is even")
} else {
    fmt.Println("7 is odd")
}

You can have an if statement whithout an else.

if 8%4 == 0 {
    fmt.Println("8 is divisible by 4")
}

A statement can precede conditionals; any variables declared in this statement are
available in the current and all subsequent branches.

Note that you don't need parentheses around conditions in Go, but that the braces are
required.

There is no ternary if in Go, so you'll need to use a full if statement even for basic
conditions.

if num := 9; num < 0 {
    fmt.Println("is negative")
} else if {
    fmt.Println("has 1 digit")
} else {
    fmt.Println("has multiple digits")
}

*/

func main() {

    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    for j := 0; j < 15; j++ {
    	num := j
        if num < 0 {
    	   fmt.Println("is negative")
    	} else if num < 10 {
    	   fmt.Println(num, "has 1 digit")
    	} else {
    	   fmt.Println(num, "has multiple digits")
    	}
    }

}
