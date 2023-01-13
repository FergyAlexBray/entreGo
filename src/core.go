package entrego

import (
	"fmt"
)

const (
	EMPTY int = 0
	PARCEL
	FORKLIFT
	TRUCK
	TARGET_PARCEL
)

type Core struct {
	Rules     GameRules
	SpaceMap  [][]int
	Parcels   []Parcel
	Trucks    []Truck
	Forklifts []Forklift
	Ticks     int
}

func (c *Core) FindEmptyForklift() (*Forklift, bool) {
	for _, forklift := range c.Forklifts {
		if forklift.Content == nil && forklift.TargetParcel == nil {
			return &forklift, true
		}
	}

	return nil, false
}

func (c *Core) FindExistingSpaceMapIndex(p Position) bool {
	cellExists := p.X >= 0 && p.Y >= 0 && p.X < c.Rules.Width && p.Y < c.Rules.Length
	return cellExists
}

func (c *Core) FindAvailableTruck(parcel Parcel) (*Truck, bool) {
	for _, truck := range c.Trucks {
		totalWeight := truck.totalWeight()
		if (totalWeight + parcel.Weight) >= truck.MaxWeight {
			continue
		}

		if truck.Available {
			return &truck, true
		}
	}

	for i, truck := range c.Trucks {
		if parcel.Weight <= truck.MaxWeight {
			c.Trucks[i].Available = false
			c.Trucks[i].RemainingTime = truck.Delay

			return &c.Trucks[i], true
		}
	}

	return nil, false
}

func (c *Core) ForkliftWithoutParcel(forklift *Forklift) {
	if forklift.IsNextToTarget(forklift.TargetParcel.Position) {
		// Take package
		forklift.TakeParcel()
		c.SpaceMap[forklift.Content.Position.Y][forklift.Content.Position.X] = EMPTY
		fmt.Println(forklift.Name, "TAKE", forklift.Content.Name, forklift.Content.Color)
	} else {
		// Move forklift to package
		forklift.MoveTowardsParcel(c)
	}
}

func (c *Core) FindTargetTruck(forklift *Forklift) bool {

	truck, available := c.FindAvailableTruck(*forklift.Content)
	if available {
		forklift.TargetTruck = truck
	} else {
		fmt.Println(forklift.Name, "WAIT")
		return false
	}

	return true
}

func (c *Core) ForkliftWithParcel(forklift *Forklift) {
	if forklift.IsNextToTarget(forklift.TargetTruck.Position) {
		// Load package into truck
		forklift.LoadTruck()
	} else {
		// Move forklift to truck
		forklift.MoveTowardsTruck(c)
	}
}

func (c *Core) UnavailableTrucksCounter() {
	for i := range c.Trucks {
		if !c.Trucks[i].Available {
			c.Trucks[i].RemainingTime -= 1

			if c.Trucks[i].RemainingTime == 0 {
				c.Trucks[i].Available = true
				c.Trucks[i].Load = make([]*Parcel, 0)
			}
		}
	}
}

func (c *Core) Init() chan struct{} {
	globalQuit := make(chan struct{})

	for i := range c.Trucks {
		go c.Trucks[i].InitTruck(globalQuit)
	}

	c.OrderParcels()

	return globalQuit
}

func (c *Core) isWareHouseEmpty() bool {
	forkliftsEmpty := true

	for _, forklift := range c.Forklifts {
		if forklift.Content != nil || forklift.TargetParcel != nil {
			forkliftsEmpty = false
		}
	}

	return forkliftsEmpty && len(c.Parcels) == 0
}

func (c *Core) Run() {
	globalQuit := c.Init()
	defer close(globalQuit)

	for i := 0; i < c.Ticks; i++ {
		if c.isWareHouseEmpty() {
			fmt.Println("😎")
			return
		}
		fmt.Println("tour", i+1)
		c.UnavailableTrucksCounter()

		for j, forklift := range c.Forklifts {
			if forklift.Content == nil && forklift.TargetParcel == nil {
				// Find package
				if len(c.Parcels) == 0 {
					if forklift.IsNextToTarget(forklift.StartPosition) {
						fmt.Println(forklift.Name, "WAIT")
					} else {
						c.Forklifts[j].MoveToStart(c)
					}
					continue
				}
				parcel := &c.Parcels[0]
				c.SpaceMap[parcel.Position.Y][parcel.Position.X] = TARGET_PARCEL
				c.Forklifts[j].TargetParcel = parcel
				c.Parcels = c.Parcels[1:]
			}

			if forklift.Content == nil {
				c.ForkliftWithoutParcel(&c.Forklifts[j])
				continue
			} else {
				if forklift.TargetTruck == nil {
					if res := c.FindTargetTruck(&c.Forklifts[j]); !res {
						continue
					}
				}
				c.ForkliftWithParcel(&c.Forklifts[j])
				continue
			}
		}

		DisplayTruckStates(*c)
		fmt.Println()
	}

	fmt.Println("🙂")
}
