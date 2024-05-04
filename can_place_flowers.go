package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	oneZero := false
	canPlace := false
	for i, item := range flowerbed {

		if item == 0 {
			if oneZero || i-1 == -1 {
				if i+1 <= len(flowerbed) {
					if i+1 == len(flowerbed) || flowerbed[i+1] == 0 {
						n--
						oneZero = false

						if n <= 0 {
							canPlace = true
							break
						}

						continue
					}
				} else {
					break
				}
			} else {
				oneZero = true
			}
		} else {
			oneZero = false
		}

	}
	return canPlace
}
