package main

import (
    "fmt"
)

func main()        {
    fmt.Println("Hello, my dearest")
    fmt.Println("I am having a Go")
    // fmt.Println("I'm not having a Go")
    fmt.Printf("%d is %b in binary and %x in hexadecimal", 42, 42, 42)
    fmt.Printf("\n")
    fmt.Printf("If I wanted to add 0x I could add a # to make it %#x", 42)
}
