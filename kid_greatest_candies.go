package main

func kidsWithCandies(candies []int, extra int) []bool {
	// you have to loop through twice, first get the largest candies and then make true false array
	greatest := 0
	for _, kid := range candies {
		if kid > greatest {
			greatest = kid
		}
	}

	isGreatest := make([]bool, 0)
	for _, kid := range candies {
		if kid == greatest {
			isGreatest = append(isGreatest, true)
		} else {
			isGreatest = append(isGreatest, false)
		}
	}
	return isGreatest
}
