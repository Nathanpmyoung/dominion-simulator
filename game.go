package main

import "fmt"
import "math/rand"

func main() {

    var a [10]int

    for i := 0; i<10 ; i++ {
        a [i] = i
    }

    fmt.Println(a)

    for i := 0;i < 10;i++ {
        a [i] = rand.Intn(10)

    }
    fmt.Println(a)
}
