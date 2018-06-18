package main 

import (
        "fmt"
        "time"
        // "strings"
        // "unicode/utf8"
        "sync"
        )

var wg sync.WaitGroup

func say(s string) {
    for i:=0; i<3; i++ {
        fmt.Println(s)
        time.Sleep(time.Millisecond*100)
    }
    wg.Done()
}

func main() {
    /*
    // c1
    say("Hey")
    say("There")

    // c2
    go say("Hey")
    say("There")

    // c3
    say("Hey")
    go say("There")
    */

    // go say("Hey")
    // go say("There")

    // time.Sleep(time.Second)

    // x := "1234567890"
    // fmt.Println(strings.Repeat("-", utf8.RuneCountInString(x)))
    // fmt.Println(strings.Repeat("-", len(x)))
    // fmt.Println(strings.Repeat("-", 10))

    wg.Add(1)
    go say("Hey")
    wg.Add(1)
    go say("There")
    wg.Wait()

    wg.Add(1)
    go say("Hi")
    wg.Wait()
}