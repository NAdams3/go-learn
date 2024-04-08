package main

func BubbleSort(unsorted []int) []int {
    for j := len(unsorted)-2; j >= 0; j-- {
        for i := 0; i <= j; i++ {
            current := unsorted[i]
            next := unsorted[i+1]
            if current > next {
                unsorted[i] = next  
                unsorted[i+1] = current 
            }
        }
    }

    return unsorted

}

