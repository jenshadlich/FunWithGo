package main

import "fmt"

func main() {
    array := [...]int{1, 2}

    for i, a := range array {
        fmt.Printf("a=%d, array[i]=%d\n", a, array[i])
    }
}