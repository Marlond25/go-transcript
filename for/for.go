package main

import "fmt"

/*

for is Go's only looping construct. Here are some basic types of for loops.

The most basic type, with a single condition.

    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

A classic initial/condition/after for loop.

    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

for without a condition will loop repeatedly until you break out of the loop or return
from the enclosing function.

    for {
        fmt.Println("loop")
        break
    }

You can also continue to the next iteration of the loop.

    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }


*/

func main() {

    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    for {
        fmt.Println("loop")
        break
    }

    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }


}
