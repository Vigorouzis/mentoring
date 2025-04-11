package main

// Задание 1 - Отложили, нужно переписать на O(N) с помощью map[int]int, где map[first]index
func findSumNumbers(ii []int, sumToFind int) (found bool, first, second int) {
	for i := 0; i < len(ii)-1; i++ {
		for j := i + 1; j < len(ii); j++ {
			if ii[i]+ii[j] == sumToFind {
				return true, ii[i], ii[j]
			}
		}
	}

	return false, 0, 0
}

func findSumNumbersWithMap(ii []int, sumToFind int) (found bool, first, second int) {
    mapOfNumbers := make(map[int]int) 
    for i := 0; i < len(ii); i++ {
        value := sumToFind - ii[i]
        
		if idx, ok := mapOfNumbers[value]; ok {
            return true, ii[idx], ii[i] 
        }

        mapOfNumbers[ii[i]] = i
    }
    return false, 0, 0
}