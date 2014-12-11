package main

import fmt "fmt"

func main() {
    for i:=0; i < 100; i++ {
        switch {
            case i % 3 == 0 && i % 5 == 0:
                fmt.Print("Fizz Buzz\n")
            case i % 3 == 0:
                fmt.Print("Fizz\n")
            case i % 5 == 0:
                fmt.Print("Buzz\n")
            default:
                fmt.Print(i , "\n")
        } 
    }
}
