package entrego

const (
	EMPTY int = 0
	PARCEL
	TRUCK
	FORKLIFT
)

type Core struct {
	Rules       GameRules
	SpaceMap    [][]int
	Parcels     []Parcel
	Trucks      []Truck
	Forklifts   []Forklift
	Identifiers SpaceMapIdentifiers
}

func (c *Core) Run() {

}
