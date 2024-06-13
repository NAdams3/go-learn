package main

import "strconv"

func IntegerToBinary(num int) (int, error) {
	temp := ""

	for num > 1 {
		remainder := num % 2
		if remainder == 0 {
			temp = "0" + temp
			num = num / 2
		} else {
			temp = "1" + temp
			num = (num - 1) / 2
		}
	}

	temp = "1" + temp

	binary, err := strconv.Atoi(temp)
	if err != nil {
		return 0, err
	}

	return binary, nil
}
