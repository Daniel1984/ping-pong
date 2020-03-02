package utils

// GetUniqueInts accepts slice of ints as input and returns slice of unique ints from input
func GetUniqueInts(input []int) (output []int) {
	found := map[int]bool{}

	for _, item := range input {
		if _, ok := found[item]; ok {
			continue
		}

		output = append(output, item)
		found[item] = true
	}

	return
}
