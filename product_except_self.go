package main

func productExceptSelf(nums []int) []int {
	product := 1
	nonZeroProduct := 1
	products := make([]int, 0)
	for _, item := range nums {
		if item != 0 {
			nonZeroProduct *= item
		}
		product = product * item
	}

	for _, item := range nums {
		if item == 0 {
			products = append(products, nonZeroProduct)
			continue
		}
		products = append(products, product/item)
	}

	return products
}
