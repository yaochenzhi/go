package main 

import "fmt"

func main() {
    for i:=0; i<10; i++ {
        fmt.Println(i)
    }

    i:=0
    for i<10 {
        fmt.Println(i)
        i++
    }

    x:=1
    for {
        fmt.Println("Do stuff", x)
        x++
        if x >25 {
            break
        }
    }

    a:=1
    for x:=5; a<5; a+=2 {
        fmt.Println("x is", x, "; a is ", a)
        x+=1
        if x == 6 {
            break
        }
    }
}