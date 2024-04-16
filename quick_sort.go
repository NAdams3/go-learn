package main

import "math"

func QuickSort(unsorted []int, high, low int) []int {

    if low >= high {
        return unsorted
    }
    
    pivotIndex := Partition(unsorted, high, low)

    QuickSort(unsorted, pivotIndex - 1, low)
    QuickSort(unsorted, high, pivotIndex + 1)

    return nil

}

func Partition(unsorted []int, high, low int) int {
    midIndex := int(math.Floor((float64(high) + float64(low))/2))
    pivot := unsorted[midIndex]

    unsorted[midIndex] = unsorted[high]
    unsorted[high] = pivot

    pivotIndex := low - 1

    for i := low; i < high; i++ {
       if unsorted[i] <= pivot {
           pivotIndex++
           temp := unsorted[i]
           unsorted[i] = unsorted[pivotIndex]
           unsorted[pivotIndex] = temp
       }
    }

    pivotIndex++
    unsorted[high] = unsorted[pivotIndex]
    unsorted[pivotIndex] = pivot

    return pivotIndex

}

