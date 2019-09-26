package mathstuff

import (
	"log"
	"sort"
)

func Average(numbers []float64) float64 {
	log.Printf("Calculating avg")
	var sum float64
	for _, num := range numbers {
		sum += num
	}
	return ((sum) / float64(len(numbers)))
}

func Median(numbers []float64) float64 {
	log.Printf("Calculating median")
	sort.Float64s(numbers)
	middleNumber := len(numbers) / 2

	if isOdd(len(numbers)) {
		return numbers[middleNumber]
	}

	return (numbers[middleNumber-1] + numbers[middleNumber]) / 2
}

func Mode(numbers []float64) []float64 {
	log.Printf("Calculating mode")
	var valueMap = make(map[float64]int)

	for _, num := range numbers {
		valueMap[num]++
	}

	currCount := valueMap[numbers[0]]
	var mode []float64
	for value, count := range valueMap {
		if currCount < count {
			currCount = count
			mode = []float64{value}
		} else if currCount == count {
			mode = append(mode, value)
		}
	}
	sort.Float64s(mode)
	return mode
}

func isOdd(i int) bool {
	return !(i%2 == 0)
}
