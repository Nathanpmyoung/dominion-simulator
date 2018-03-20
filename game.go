package main

import "fmt"
import "math/rand"

func shuffle(a int) int {
    return a
}

func main() {

    var a, b [10]int

    for i := 0; i<10 ; i++ {
        a [i] = i
    }

    for i := 0; i<10 ; i++ {
        a [i]

    }

    fmt.Println(a)

    for i := 0;i < 10;i++ {
        a [i] = rand.Intn(10)

    }
    fmt.Println(a)
}
