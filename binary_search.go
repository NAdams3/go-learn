package main

import (
    "math"
)

func BinarySearch(haystack []int, needle int) bool {
    found := false

    low := float64(0)
    high := float64(len(haystack))
    checkIndex := 0

    for low < high {

        checkIndex = int(math.Floor((low + high) / 2))    

        if needle > haystack[checkIndex] {
            low = float64(checkIndex + 1)
        } else if needle < haystack[checkIndex] {
            high = float64(checkIndex)
        } else {
            found = true
            break
        }

    }

    return found
}

func BinarySearchRecursive(haystack []int, needle int) bool {

    low := float64(0)
    high := float64(len(haystack) - 1)
   
    return BSearchRecursive(haystack, needle, low, high) 

}


func BSearchRecursive(haystack []int, needle int, low, high float64) bool {
    checkIndex := int(math.Floor((low + high) / 2))

    if needle == haystack[checkIndex] {
        return true
    }
    
    if low >= high {
        return false
    }

    if needle < haystack[checkIndex] {
        high = float64(checkIndex)
    } else {
        low = float64(checkIndex + 1) 
    }

    return BSearchRecursive(haystack, needle, low, high)
}

