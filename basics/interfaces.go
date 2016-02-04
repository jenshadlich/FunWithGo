package main

// inspired by http://blog.jonathanoliver.com/golang-has-generics/

import "fmt"

type Integer int32
type Long int64

type Calculator interface {
    Calculate()
}

func (i Integer) Calculate() {}
func (l Long) Calculate() {}

func main() {
    collection := [] Calculator{ Integer(1), Long(2) }
    fmt.Println(collection)
}