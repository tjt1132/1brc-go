package aggregator

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var allStations StationStatsMap = StationStatsMap{}
var sortedStationNames []string

func Do(filepath string) {
	processFile(filepath)
	sortStationNames()
	printResult()
}

func processFile(filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(FileOpenErr{Err: err})
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		NameTemperature := strings.Split(scanner.Text(), ";")
		Name := NameTemperature[0]
		Temperature, err := strconv.ParseFloat(NameTemperature[1], 64)
		if err != nil {
			log.Fatal(ParceFloatErr{Err: err})
		}
		if station, ok := allStations[Name]; ok {
			station.AddNewValue(Temperature)
		} else {
			allStations[Name] = &StationStats{
				Name:  Name,
				Min:   Temperature,
				Max:   Temperature,
				Total: Temperature,
				Count: 1,
			}
		}
	}
}

func sortStationNames() {
	sortedStationNames = make([]string, 0, len(allStations))
	for k := range allStations {
		sortedStationNames = append(sortedStationNames, k)
	}
	sort.Strings(sortedStationNames)
}

func printResult() {
	fmt.Print("{")
	fmt.Print(allStations[sortedStationNames[0]].Report())
	for _, name := range sortedStationNames[1:] {
		fmt.Print(", " + allStations[name].Report())
	}
	fmt.Print("}")
}
