package entrego

import "sort"

type Parcel struct {
	Name     string
	Weight   int
	Position Position
	Color    string
}

func (c *Core) OrderParcels() {
	sort.SliceStable(c.Parcels, func(i, j int) bool {
		return c.Parcels[i].Weight < c.Parcels[j].Weight
	})
}
