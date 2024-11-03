package gnc1431

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var max int
	for _, nr := range candies {
		if nr >= max {
			max = nr
		}
	}
	result := make([]bool, len(candies))
	for idx, nrCandies := range candies {
		result[idx] = nrCandies+extraCandies >= max
	}
	return result
}
