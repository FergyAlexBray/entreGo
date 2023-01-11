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
	x, _ := strconv.Atoi(splittedData[1])
	y, _ := strconv.Atoi(splittedData[2])
	position := Position{X: x, Y: y}
	weight, _ := strconv.Atoi(splittedData[3])
	newParcel := Parcel{Name: splittedData[0], Weight: weight, Position: position}
	*parcels = append(*parcels, newParcel)
}

func addForkliftToForklifts(forklifts *[]Forklift, splittedData []string) {
	x, _ := strconv.Atoi(splittedData[1])
	y, _ := strconv.Atoi(splittedData[2])
	position := Position{X: x, Y: y}
	newForklift := Forklift{Name: splittedData[0], Position: position}
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
	x, _ := strconv.Atoi(splittedData[1])
	y, _ := strconv.Atoi(splittedData[2])
	position := Position{X: x, Y: y}
	maxWeight, _ := strconv.Atoi(splittedData[3])
	delay, _ := strconv.Atoi(splittedData[4])
	newTruck := Truck{Name: splittedData[0], Position: position, MaxWeight: maxWeight, Delay: delay}
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

func addTrucksToSpaceMap(c *Core) {
	for _, truck := range c.Trucks {
		c.SpaceMap[truck.Position.Y][truck.Position.X] = c.Identifiers.Truck
	}
}
func addForkliftsToSpaceMap(c *Core) {
	for _, forklift := range c.Forklifts {
		c.SpaceMap[forklift.Position.Y][forklift.Position.X] = c.Identifiers.Forklift
	}
}

func addParcelsToSpaceMap(c *Core) {
	for _, parcel := range c.Parcels {
		c.SpaceMap[parcel.Position.Y][parcel.Position.X] = c.Identifiers.Parcel
	}
}

func populateSpaceMap(c *Core) {
	createSpaceMapBase(c)
	addTrucksToSpaceMap(c)
	addParcelsToSpaceMap(c)
	addForkliftsToSpaceMap(c)
}

func createSpaceMapBase(c *Core) {
	c.SpaceMap = make([][]int, 0)
	for i := 0; i < c.Rules.Length; i++ {
		tmp := make([]int, 0)
		for j := 0; j < c.Rules.Width; j++ {
			tmp = append(tmp, c.Identifiers.Space)
		}
		c.SpaceMap = append(c.SpaceMap, tmp)
	}
}

func setIdentifiers(c *Core) {
	c.Identifiers.Truck = 3
	c.Identifiers.Parcel = 1
	c.Identifiers.Forklift = 2
	c.Identifiers.Space = 0
}

func Parser(c *Core, args []string) {
	setIdentifiers(c)

	if len(args) < 1 {
		return
	}

	lines := readFileIntoArray(args[1])
	validateSpecifiedData(lines)
	setCoreDataFromFileLines(c, lines)
	populateSpaceMap(c)
	fmt.Println(c.SpaceMap)
}
