package entrego

import (
	"fmt"
	"strconv"
)

func EndMessage() {
	fmt.Println("EntreGo !")
}

func CoordinatesToString(pos Position) string {
	return "[" + strconv.Itoa(pos.X) + "," + strconv.Itoa(pos.Y) + "]"
}

func DisplayTruckStates(c Core) {
	for _, truck := range c.Trucks {
		truckState := "WAITING"
		truckLoad := strconv.Itoa(truck.totalWeight()) + "/" + strconv.Itoa(truck.MaxWeight)

		if !truck.Available {
			truckState = "GONE"
		}

		fmt.Println(truck.Name, truckState, truckLoad)
	}
}

func DisplayMap(c Core) {
	for _, line := range c.SpaceMap {
		fmt.Println(line)
	}
}
