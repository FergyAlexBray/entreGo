package entrego

import "sort"

var Colors = map[string]int{
	"BLUE":   500,
	"GREEN":  200,
	"YELLOW": 100,
}

type Parcel struct {
	Name     string
	Weight   int
	Color    string
	Position Position
}

func (c *Core) OrderParcels() {
	sort.SliceStable(c.Parcels, func(i, j int) bool {
		return c.Parcels[i].Weight < c.Parcels[j].Weight
	})
}
