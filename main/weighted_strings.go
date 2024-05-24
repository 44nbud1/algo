package main

func WeightedStrings(s string, queries []int) []string {

	// init a map to store the weights
	weights := make(map[int]bool)

	// iterate through the string to calculate weights
	for i := 0; i < len(s); i++ {
		currentChar := s[i]
		currentWeight := int(currentChar - 'a' + 1)
		repeatCount := 1

		// add the weight of the single character
		weights[currentWeight] = true

		// check for consecutive characters
		for j := i + 1; j < len(s) && s[j] == currentChar; j++ {
			repeatCount++
			currentWeight += int(currentChar - 'a' + 1)
			weights[currentWeight] = true
			i = j
		}
	}

	// process each query to check if it exists in the weights map
	result := make([]string, len(queries))
	for i, query := range queries {
		if weights[query] {
			result[i] = "Yes"
		} else {
			result[i] = "No"
		}
	}

	// return result
	return result
}
