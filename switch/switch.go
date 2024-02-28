package main

/*

Switch statements express conditionals across many branches

Here's a basic switch.

i := 2
fmt.Print("Write ", i, " as ")
switch i {
case 1:
    fmt.Println("one")
case 2:
    fmt.Println("two")
case 3:
    fmt.Println("three")
}

Commas can be used to separate multiple expressions in the same case statement. The
optional default case is used in this example as well.

switch time.Now().Weekday() {
case time.Saturday, time.Sunday:
    fmt.Println("It's the weekend")
default:
    fmt.Println("It's a weekday")
}

switch without an expression is an alternate way to express if/else logic.
Here the use of the case expression is shown to be non-constant.

t := time.Now()
switch {
case t.Hour() < 12:
    fmt.Println("It's before noon")
default:
    fmt.Println("It's after noon")
}

A type switch compares types instead of values.
You can use this to discover the type of an interface value. In this example,
the variable t will have the type corresponding to its clause.

The most important pattern to extrace from the following example is the assignment
of a switch to a function in order to make it reusable.

whatAmI := func(i interface) {
    switch t := i.(type) {
    case bool:
        fmt.Println("I'm a bool")
    case int:
        fmt.Println("I'm an int %d", t)
    default:
        fmt.Printf("Don't know type %T/n", t)
    }
}


*/


import (
    "fmt"
    "time"
)


func main() {

    i := 2
    fmt.Print("Write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }

    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }

    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int %d", t)
        default:
            fmt.Printf("Don't know type %T/n", t)
        }
    }
    
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}

