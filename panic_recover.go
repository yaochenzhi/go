package main 

import (
        "fmt"
        "sync"
        "time"
)

var wg sync.WaitGroup

func cleanup() {
    if r := recover(); r != nil {
        fmt.Println("Recovered in cleanup", r)
    }
}

func say(s string) {
    defer wg.Done()
    defer cleanup()
    for i:=0; i<3; i++ {
        time.Sleep(time.Millisecond*100)
        fmt.Println(s)
        if i == 2 {
            panic("Oh dear, a 2 accurs")
        }
    }
}

func main() {
    wg.Add(1)
    go say("Hello world!")
    wg.Add(1)
    go say("Hello testing!")
    wg.Wait()
}