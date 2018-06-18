package main 

import "fmt"

func foo() {
    // run whatever in the end.
    defer fmt.Println("Done!")
    defer fmt.Println("Are we done?")
    fmt.Println("Do some stuff.")
    for i:=0; i<5; i++ {
        defer fmt.Println(i)
    }
}

func main() {
    foo()   
}