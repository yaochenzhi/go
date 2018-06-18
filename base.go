package main 

import ("fmt"
        "math/rand")

func foo() {
    fmt.Println("A rand number from 1-100: ", rand.Intn(200))
}

// func add(x float64,y float64) float64 {
//     return x+y
// }

func add(x,y float64) float64 {
    return x+y
}

func multiple(a,b string) (string, string) {
    return a,b
}

func main() {
    // var num1 float64 = 5.6
    // var num2 float64 = 6.6

    // var num1,num2 float64 = 5.6, 6.6

    /*    
    num1, num2  := 5.6, 6.6     //default float64
    fmt.Println(add(num1,num2))
    */

    w1, w2 := "Hey", "there"
    fmt.Println(multiple(w1,w2))

    // var a int = 62
    // var b float64 = float64(a)

    // x := a // x will be type int


    x := 15
    a := &x  // memory address
    fmt.Println("Memory address of var x is", a)
    fmt.Println("Val of memory address", a, "is", *a)

    *a = 5
    fmt.Println(x)
    *a = *a**a
    fmt.Println(x)
    fmt.Println(*a)
}