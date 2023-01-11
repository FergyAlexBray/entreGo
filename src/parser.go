package entrego

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Rules struct {
	width  int
	length int
	rounds int
}

func readFileIntoArray(filePath string) []string {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	sc := bufio.NewScanner(file)
	return fileToArrayByLines(sc)
}

func fileToArrayByLines(sc *bufio.Scanner) []string {
	lines := make([]string, 0)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func setSizeOfMap(rules *GameRules, splittedData []string) {
	rules.Length, _ = strconv.Atoi(splittedData[0])
	rules.Width, _ = strconv.Atoi(splittedData[1])
}

func setNumberOfRounds(rules *GameRules, splittedData []string) {
	rules.Rounds, _ = strconv.Atoi(splittedData[2])
}

func setSizeAndNumberOfRound(rules *GameRules, lines *[]string) {
	linesTemp := *lines
	splittedData := strings.Split(linesTemp[0], " ")
	if len(splittedData) > 2 {
		setSizeOfMap(rules, splittedData)
		setNumberOfRounds(rules, splittedData)
		*lines = linesTemp[1:]
	} else {
		log.Fatal("invalid rules")
	}
}

func validateSpecifiedData(lines []string) {
	for index, element := range lines {
		fmt.Println("At index", index, "value is", element)
	}
}

func setParcels(parcels *[]Parcel, lines *[]string) {
	linesTemp := *lines
	for index, element := range linesTemp {
		*lines = linesTemp[index:]
		splittedData := strings.Split(element, " ")
		if len(splittedData) == 4 {
			appendParcelToParcels(parcels, splittedData)
		} else {
			return
		}
	}
}

func appendParcelToParcels(parcels *[]Parcel, splittedData []string) {
	index := [2]int{0, 1}
	index[0], _ = strconv.Atoi(splittedData[2])
	index[1], _ = strconv.Atoi(splittedData[3])
	weight, _ := strconv.Atoi(splittedData[1])
	newParcel := Parcel{Name: splittedData[0], Weight: weight, Position: index}
	*parcels = append(*parcels, newParcel)
}

func addForkliftToForklifts(forklifts *[]Forklift, splittedData []string) {
	index := [2]int{0, 1}
	index[0], _ = strconv.Atoi(splittedData[1])
	index[1], _ = strconv.Atoi(splittedData[2])
	newForklift := Forklift{Name: splittedData[0], Position: index}
	*forklifts = append(*forklifts, newForklift)
}

func setForklifts(forklifts *[]Forklift, lines *[]string) {
	linesTemp := *lines
	for index, element := range linesTemp {
		*lines = linesTemp[index:]
		splittedData := strings.Split(element, " ")
		if len(splittedData) == 3 {
			addForkliftToForklifts(forklifts, splittedData)
		} else {
			return
		}
	}
}

func addTruckToTrucks(trucks *[]Truck, splittedData []string) {
	index := [2]int{0, 1}
	index[0], _ = strconv.Atoi(splittedData[1])
	index[1], _ = strconv.Atoi(splittedData[2])
	maxWeight, _ := strconv.Atoi(splittedData[3])
	delay, _ := strconv.Atoi(splittedData[4])
	newTruck := Truck{Name: splittedData[0], Position: index, MaxWeight: maxWeight, Delay: delay}
	*trucks = append(*trucks, newTruck)
}

func setTrucks(trucks *[]Truck, lines *[]string) {
	linesTemp := *lines
	for index, element := range linesTemp {
		*lines = linesTemp[index:]
		splittedData := strings.Split(element, " ")
		if len(splittedData) == 5 {
			addTruckToTrucks(trucks, splittedData)
		} else {
			return
		}
	}
}

func setCoreDataFromFileLines(c *Core, lines []string) {
	setSizeAndNumberOfRound(&c.Rules, &lines)
	setParcels(&c.Parcels, &lines)
	setForklifts(&c.Forklifts, &lines)
	setTrucks(&c.Trucks, &lines)
}

func Parser(c *Core, args []string) {

	fmt.Println(args[1])

	if len(args) < 1 {
		return
	}

	lines := readFileIntoArray(args[1])
	validateSpecifiedData(lines)
	setCoreDataFromFileLines(c, lines)
}
