package mathstuff

import (
	"log"
	"math"
	"sort"
)

func Average(numbers []float64) float64 {
	log.Printf("Calculating avg")
	var sum float64
	for _, num := range numbers {
		sum += num
	}
	avg := ((sum) / float64(len(numbers)))
	return math.Round(avg*100) / 100
}

func Median(numbers []float64) float64 {
	log.Printf("Calculating median")
	sort.Float64s(numbers)
	middleNumber := len(numbers) / 2

	if isOdd(len(numbers)) {
		return numbers[middleNumber]
	}

	median := (numbers[middleNumber-1] + numbers[middleNumber]) / 2
	return math.Round(median*100) / 100
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
			mode = []float64{math.Round(value*100) / 100}
		} else if currCount == count {
			mode = append(mode, math.Round(value*100)/100)
		}
	}
	sort.Float64s(mode)
	return mode
}

func isOdd(i int) bool {
	return !(i%2 == 0)
}
