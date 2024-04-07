package main

func LinearSearch(haystack []int, needle int) bool {

    found := false

    for _, n := range haystack {
        if n == needle {
            found = true
            break
        }
    }

    return found

}

