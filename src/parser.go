package entrego

import (
	"bufio"
	"log"
	"os"
	"strings"
)

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
	rules.Length = convertStringToNumber(splittedData[0])
	rules.Width = convertStringToNumber(splittedData[1])
}

func setNumberOfRounds(rules *GameRules, splittedData []string) {
	rules.Rounds = convertStringToNumber(splittedData[2])
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

func getWeightFromColor(color string) int {
	switch color {
	case "YELLOW", "yellow":
		return 100
	case "GREEN", "green":
		return 200
	case "BLUE", "blue":
		return 500
	default:
		log.Fatal("Invalid weight for parcel specified.")
		return 0
	}
}

func appendParcelToParcels(parcels *[]Parcel, splittedData []string) {
	position := getParcelPosition(splittedData)
	weight := getWeightFromColor(splittedData[3])
	newParcel := Parcel{Name: splittedData[0], Weight: weight, Position: position, Color: splittedData[3]}
	*parcels = append(*parcels, newParcel)
}

func getParcelPosition(splittedData []string) Position {
	position := Position{X: convertStringToNumber(splittedData[1]), Y: convertStringToNumber(splittedData[2])}
	return position
}

func addForkliftToForklifts(forklifts *[]Forklift, splittedData []string) {
	position := Position{X: convertStringToNumber(splittedData[1]), Y: convertStringToNumber(splittedData[2])}
	newForklift := Forklift{Name: splittedData[0], Position: position, StartPosition: position}
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
	position := Position{X: convertStringToNumber(splittedData[1]), Y: convertStringToNumber(splittedData[2])}
	loadTruck := make(chan LoadPackage)
	newTruck := Truck{Name: splittedData[0], Position: position, MaxWeight: convertStringToNumber(splittedData[3]), Delay: convertStringToNumber(splittedData[4]), Available: true, LoadTruck: loadTruck}
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

func setTicks(ticks *int, rules *GameRules) {
	*ticks = rules.Rounds
}

func setCoreDataFromFileLines(c *Core, lines []string) {
	setSizeAndNumberOfRound(&c.Rules, &lines)
	setTicks(&c.Ticks, &c.Rules)
	setParcels(&c.Parcels, &lines)
	setForklifts(&c.Forklifts, &lines)
	setTrucks(&c.Trucks, &lines)
}

func addTrucksToSpaceMap(c *Core) {
	for _, truck := range c.Trucks {
		if c.FindExistingSpaceMapIndex(truck.Position) {
			c.SpaceMap[truck.Position.Y][truck.Position.X] = TRUCK
		} else {
			log.Fatal("Error: Invalid position specified for Truck.")
		}
	}
}
func addForkliftsToSpaceMap(c *Core) {
	for _, forklift := range c.Forklifts {
		if c.FindExistingSpaceMapIndex(forklift.Position) {
			c.SpaceMap[forklift.Position.Y][forklift.Position.X] = FORKLIFT
		} else {
			log.Fatal("Error: Invalid position specified for forklift.")
		}
	}
}

func addParcelsToSpaceMap(c *Core) {
	for _, parcel := range c.Parcels {
		if c.FindExistingSpaceMapIndex(parcel.Position) {
			c.SpaceMap[parcel.Position.Y][parcel.Position.X] = PARCEL
		} else {
			log.Fatal("Error: Invalid position specified for parcel.")
		}
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
			tmp = append(tmp, EMPTY)
		}
		c.SpaceMap = append(c.SpaceMap, tmp)
	}
}

func Parser(c *Core, args []string) {
	if len(args) < 1 {
		return
	}

	lines := readFileIntoArray(args[1])
	setCoreDataFromFileLines(c, lines)
	populateSpaceMap(c)
}
