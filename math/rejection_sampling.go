package main

import (
    "fmt"
    "math/rand"
    "time"
)

type Element struct {
    value   int
    penalty uint
}

type weightedRand struct {
    elements []Element
    seed     int64
    rand     *rand.Rand
}

func NewWR(seed int64, elements []Element) weightedRand {
    return weightedRand{
        elements: elements,
        seed:     seed,
        rand:     rand.New(rand.NewSource(seed)),
    }
}

func (w weightedRand) RandElem() int {
    len := len(w.elements)
    idx := w.rand.Intn(len)
    randPenalty := w.rand.Intn(int(w.maxPenalty() + 1))
    for w.elements[idx].penalty > uint(randPenalty) {
        idx = w.rand.Intn(len)
    }
    return w.elements[idx].value
}
func (w weightedRand) maxPenalty() uint {
    max := uint(0)
    for _, e := range w.elements {
        if e.penalty > max {
            max = e.penalty
        }
    }
    return max
}

func main() {
    elements := []Element{
        Element{1, 0},
        Element{2, 10},
        Element{3, 5},
    }
    wr := NewWR(time.Now().UnixNano(), elements)

    d := map[int]float64{1: 0.0, 2: 0.0, 3: 0.0}
    sampleCnt := 10000
    for i := 0; i < sampleCnt; i++ {
        val := wr.RandElem()
        d[val] = d[val] + 1.0
    }
    for idx, val := range d {
        d[idx] = val / float64(sampleCnt)
    }
    for idx, prob := range d {
        fmt.Printf("val: %d, p(val): %f\n", idx, prob)
    }
}