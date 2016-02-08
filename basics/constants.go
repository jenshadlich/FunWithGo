package main
import "fmt"

type Base string

const (
    A Base = "A"
    C Base = "C"
    T Base = "T"
    G Base = "G"
)

func main() {
    fmt.Println(A)
}
