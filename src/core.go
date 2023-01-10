package entrego

import (
	"fmt"
	"sort"
)

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
	Ticks     int
}

func (c *Core) OrderParcels() {
	sort.SliceStable(c.Parcels, func(i, j int) bool {
		return c.Parcels[i].Weight < c.Parcels[j].Weight
	})
}

func (c Core) FindEmptyForklift() (*Forklift, bool) {
	for _, forklift := range c.Forklifts {
		if forklift.Content == nil && forklift.TargetParcel == nil {
			return &forklift, true
		}
	}

	return nil, false
}

func (c Core) FindAvailableTruck(parcel Parcel) (*Truck, bool) {
	for _, truck := range c.Trucks {
		totalWeight := truck.totalWeight()
		if (totalWeight + parcel.Weight) >= truck.MaxWeight {
			continue
		}

		if truck.Available {
			return &truck, true
		}
	}

	return nil, false
}

func (c *Core) ForkliftWithoutParcel(forklift *Forklift) {
	if forklift.IsNextToTarget(forklift.TargetParcel.Position) {
		// Take package
		forklift.TakeParcel()
		fmt.Println("Take Parcel")
	} else {
		// Move forklift to package
		forklift.MoveTowardsParcel(c)
		fmt.Println("Move")
	}
}

func (c Core) FindTargetTruck(forklift *Forklift) bool {

	truck, available := c.FindAvailableTruck(*forklift.Content)
	if available {
		forklift.TargetTruck = truck
	} else {
		fmt.Println("Waiting...")
		return false
	}

	return true
}

func (c *Core) ForkliftWithParcel(forklift *Forklift) {
	if forklift.IsNextToTarget(forklift.TargetTruck.Position) {
		// Load package into truck
		forklift.LoadTruck()
		fmt.Println("Load truck")
	} else {
		// Move forklift to truck
		forklift.MoveTowardsTruck(c)
		fmt.Println("Move")
	}
}

func (c *Core) UnavailableTrucksCounter() {
	for _, truck := range c.Trucks {
		if !truck.Available {
			truck.RemainingTime -= 1
		}
	}
}

func (c *Core) Run() {
	globalQuit := make(chan struct{})
	defer close(globalQuit)

	for _, truck := range c.Trucks {
		go truck.InitTruck(globalQuit)
	}

	c.OrderParcels()

	for i := 0; i < c.Ticks; i++ {
		c.UnavailableTrucksCounter()

		for _, forklift := range c.Forklifts {
			if forklift.Content == nil && forklift.TargetParcel == nil {
				if len(c.Parcels) == 0 {
					// TODO End condition: Show end string
					return
				}
				// Find package
				forklift.TargetParcel = &c.Parcels[0]
				c.Parcels = c.Parcels[1:]
			}

			if forklift.Content == nil {
				c.ForkliftWithoutParcel(&forklift)
				continue
			} else {
				if forklift.TargetTruck == nil {
					if res := c.FindTargetTruck(&forklift); !res {
						continue
					}
				}

				c.ForkliftWithParcel(&forklift)
				continue
			}
		}
	}
}
