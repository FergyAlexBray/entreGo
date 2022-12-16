package entrego

const (
	EMPTY int = 0
	PARCEL
	TRUCK
	FORKLIFT
)

type Core struct {
	SpaceMap  [][]int
	Parcels   []Parcel
	Trucks    []Truck
	Forklifts []Forklift
}

func (c *Core) Run() {

}
