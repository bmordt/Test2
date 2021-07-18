package main

import (
	"Test2/src/mathstuff"
	"Test2/src/objects"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func worker(id string, jobs <-chan []float64, results chan<- objects.FridgeResult) {
	numbers := <-jobs
	log.Printf("Working id:%s \n", id)
	results <- objects.FridgeResult{
		Id:      id,
		Average: mathstuff.Average(numbers),
		Median:  mathstuff.Median(numbers),
		Mode:    mathstuff.Mode(numbers),
	}

	//test commit 1
	//test commit blah 2
	//test commit blah 3
}

func main() {
	var fileName string
	flag.StringVar(&fileName, "f", "", "location and name of json data file")
	flag.Parse()

	if strings.Compare(fileName, "") == 0 {
		log.Fatalf("Please enter a value for -f. location and name of json data file")
	}
	log.Printf("Getting data from %s", fileName)

	jsonFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Cannot open the file %s", err)
	}

	log.Println("Successfully opened the file")
	defer jsonFile.Close()

	jsonBytes, readErr := ioutil.ReadAll(jsonFile)
	if readErr != nil {
		log.Fatalf("Cannot read the file %s", err)
	}

	var fridges []objects.Fridge
	errJson := json.Unmarshal(jsonBytes, &fridges)
	if errJson != nil {
		log.Fatalf("Error unmarshalling json into fridge objects. %s", errJson)
	}

	log.Printf("Fridges: %v", fridges)

	var fridgesA []objects.Fridge
	var fridgesB []objects.Fridge
	var fridgesC []objects.Fridge
	for _, fridge := range fridges {
		switch fridge.Id {
		case "a":
			fridgesA = append(fridgesA, fridge)
			break
		case "b":
			fridgesB = append(fridgesB, fridge)
			break
		case "c":
			fridgesC = append(fridgesC, fridge)
			break
		default:
			log.Printf("There is an unkown ID in the list %s ", fridge.Id)
			break
		}
	}

	var allFridges [][]objects.Fridge
	allFridges = append(allFridges, fridgesA)
	allFridges = append(allFridges, fridgesB)
	allFridges = append(allFridges, fridgesC)

	numFridges := len(allFridges)

	// In order to use our pool of workers we need to send
	// them work and collect their results. We make 2
	// channels for this.
	jobs := make(chan []float64, numFridges)
	results := make(chan objects.FridgeResult, numFridges)

	// This starts up 3 workers, initially blocked
	// because there are no jobs yet.
	for _, fridgeList := range allFridges {
		go worker(fridgeList[0].Id, jobs, results)
	}

	// Here we send 5 `jobs` and then `close` that
	// channel to indicate that's all the work we have.
	for _, fridgeList := range allFridges {
		jobs <- objects.GetTemperatureListForFridge(fridgeList)
	}
	close(jobs)

	// Finally we collect all the results of the work.
	// This also ensures that the worker goroutines have
	// finished. An alternative way to wait for multiple
	// goroutines is to use a [WaitGroup](waitgroups).
	var resultList []objects.FridgeResult
	for a := 1; a <= numFridges; a++ {
		resultList = append(resultList, <-results)
	}

	log.Printf("Results %+v", resultList)
}
