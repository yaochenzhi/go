package main 

import "fmt"

func main() {
    a := []int{1, 2, 3}
    fmt.Println(a[0])

    b := a[1:]
    fmt.Println(b)
}