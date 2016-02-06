package main

// inspired by http://blog.jonathanoliver.com/golang-has-generics/

import "fmt"

type Integer int32
type Long int64

type GenericCollection interface {
    Something()
}

func (i Integer) Something() {
    fmt.Println("Integer = ", i)
}

func (l Long) Something() {
    fmt.Println("Long = ", l)
}

func main() {
    collection := [] GenericCollection{Integer(1), Long(2) }
    fmt.Println(collection)

    for _, element := range collection {
        element.Something()
    }
}