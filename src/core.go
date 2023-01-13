package entrego

import (
	"fmt"
)

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
	Ticks       int
}

func (c Core) FindEmptyForklift() (*Forklift, bool) {
	for _, forklift := range c.Forklifts {
		if forklift.Content == nil && forklift.TargetParcel == nil {
			return &forklift, true
		}
	}

	return nil, false
}

func (c Core) FindExistingSpaceMapIndex(p Position) bool {
	cellExists := p.X >= 0 && p.Y >= 0 && p.X < c.Rules.Width && p.Y < c.Rules.Length
	return cellExists
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

func (c *Core) ForkliftWithoutParcel(forklift *Forklift) *Forklift {
	if forklift.IsNextToTarget(forklift.TargetParcel.Position) {
		// Take package
		forklift.TakeParcel()
		fmt.Println("Take Parcel")
	} else {
		// Move forklift to package
		forklift.MoveTowardsParcel(c)
		fmt.Println("Move")
	}
	return forklift
}

func (c Core) FindTargetTruck(forklift *Forklift) (bool, *Forklift) {

	truck, available := c.FindAvailableTruck(*forklift.Content)
	if available {
		forklift.TargetTruck = truck
	} else {
		fmt.Println("Waiting...")
		return false, forklift
	}

	return true, forklift
}

func (c *Core) ForkliftWithParcel(forklift *Forklift) *Forklift {
	if forklift.IsNextToTarget(forklift.TargetTruck.Position) {
		// Load package into truck
		forklift.LoadTruck()
		fmt.Println("Load truck")
	} else {
		// Move forklift to truck
		forklift.MoveTowardsTruck(c)
		fmt.Println("Move")
	}
	return forklift
}

func (c *Core) UnavailableTrucksCounter() {
	for _, truck := range c.Trucks {
		if !truck.Available {
			truck.RemainingTime -= 1
		}
	}
}

func (c *Core) Init() chan struct{} {
	globalQuit := make(chan struct{})

	for _, truck := range c.Trucks {
		go truck.InitTruck(globalQuit)
	}

	c.OrderParcels()

	return globalQuit
}

func (c *Core) Run() {
	globalQuit := c.Init()
	defer close(globalQuit)

	for i := 0; i < c.Ticks; i++ {
		c.UnavailableTrucksCounter()

		for j, forklift := range c.Forklifts[:] {
			if forklift.Content == nil && forklift.TargetParcel == nil {
				if len(c.Parcels) == 0 {
					// TODO End condition: Show end string
					return
				}
				// Find package
				if len(c.Parcels) > 1 {
					forklift.TargetParcel = &c.Parcels[j+1]
				} else {
					forklift.TargetParcel = &c.Parcels[j]
				}

				c.Parcels = c.Parcels[1:]
			}

			if forklift.Content == nil {
				c.Forklifts[j] = *c.ForkliftWithoutParcel(&forklift)
				continue
			} else {
				if forklift.TargetTruck == nil {
					res, updatedForklift := c.FindTargetTruck(&forklift)
					c.Forklifts[j] = *updatedForklift
					if !res {
						continue
					}
				}
				c.Forklifts[j] = *c.ForkliftWithParcel(&forklift)
				continue
			}
		}
	}
}
