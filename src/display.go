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

		if i < len(transpalette) && transpalette[i].Name != "" {
			fmt.Printf(transpalette[i].Name + " ")
		}

		if i < len(transpalette) && transpalette[i].State != "" {
			switch transpalette[i].State {
			case "GO":
				fmt.Printf(transpalette[i].State + " " + transpalette[i].Coordinates + "\n")
			case "WAIT":
				fmt.Printf(transpalette[i].State + "\n")
			case "TAKE":
				fmt.Printf(transpalette[i].State + " " + transpalette[i].Colis + " " + transpalette[i].Color + "\n")
			case "LEAVE":
				fmt.Printf(transpalette[i].State + " " + transpalette[i].Colis + " " + transpalette[i].Color + "\n")
			}
		}

		if i < len(camion) && camion[i].Name != "" {
			fmt.Printf(camion[i].Name + " " + camion[i].State + " " + camion[i].Weight + "\n\n")
		}
	}
}
