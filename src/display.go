package entrego

import (
	"fmt"
)

type Transpalette struct {
	Name        string
	State       string
	Coordinates string
	Colis       string
	Color       string
}

type Camion struct {
	Name   string
	State  string
	Weight string
}

func EndMessage(c Core) {
	fmt.Println("EntreGo !")
}

func Display(tour int, transpalette []Transpalette, camion []Camion) {
	if tour < 10 || tour > 10000 {
		fmt.Printf("Number must be between 10 and 10000")
		return
	}

	for i := 0; i < tour; i++ {
		fmt.Printf("tour %d\n", i+1)
		for _, t := range transpalette {
			fmt.Printf(t.Name + " ")
			switch t.State {
			case "GO":
				fmt.Printf(t.State + " " + t.Coordinates + "\n")
			case "WAIT":
				fmt.Printf(t.State + "\n")
			case "TAKE":
				fmt.Printf(t.State + " " + t.Colis + " " + t.Color + "\n")
			case "LEAVE":
				fmt.Printf(t.State + " " + t.Colis + " " + t.Color + "\n")
			}
		}
		for _, c := range camion {
			fmt.Printf(c.Name + " " + c.State + " " + c.Weight + "\n")
		}
		fmt.Printf("\n")
	}
}
