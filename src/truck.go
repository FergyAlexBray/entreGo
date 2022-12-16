package entrego

type Truck struct {
	Name      string
	Position  [2]int
	MaxWeight int
	Delay     int
	Load      []Parcel
}
