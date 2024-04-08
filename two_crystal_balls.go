package main

import "math"

func TwoCrystalBalls(tower []int) int {
    index := -1

    low := 0
    high := len(tower)

    increment := int(math.Ceil(math.Sqrt(float64(high))))

    for i := 0; i < high; i += increment {
        if tower[i] == 1 {
            break
        }
        low = i
    }
    
    for i := low; i < high; i++ {
        if tower[i] == 1 {
            index = i
            break
        }
    }

    return index

}

