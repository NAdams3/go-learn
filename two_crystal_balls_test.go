package main

import (
    "testing"
    "math/rand"
)

func TestTwoCrystalBalls(t *testing.T) {
    index := rand.Intn(10000)
    floors := make([]int, 10000)

    got := TwoCrystalBalls(floors)
    if got != -1 {
        t.Errorf("expected -1, got %d", got)
    }

    for i := index; i < len(floors); i++ {
        floors[i]++
    }

    got = TwoCrystalBalls(floors)
    if got != index {
        t.Errorf("exptected %d, got %d", index, got)
    }

}

