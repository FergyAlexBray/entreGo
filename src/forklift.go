package entrego

import "fmt"

type Position struct {
	X int
	Y int
}

type Forklift struct {
	Name         string
	Position     Position
	Content      *Parcel
	TargetTruck  *Truck
	TargetParcel *Parcel
}

func (f *Forklift) move(c *Core, target Position, targetType int) {
	nextPosition := FindShortestPath(c.SpaceMap, f.Position, target)[0]

	c.SpaceMap[f.Position.X][f.Position.Y] = 0 // Number representing empty
	c.SpaceMap[nextPosition.X][nextPosition.Y] = targetType

	f.Position = nextPosition
}

func (f *Forklift) MoveTowardsParcel(c *Core) {
	// TODO: Set the right targetType number
	f.move(c, f.TargetParcel.Position, 1)
}

func (f *Forklift) MoveTowardsTruck(c *Core) {
	// TODO: Set the right targetType number
	f.move(c, f.TargetTruck.Position, 2)
}

func (f *Forklift) TakeParcel() {
	f.Content = f.TargetParcel
	f.TargetParcel = nil
}

func (f *Forklift) LoadTruck() {
	f.TargetTruck.LoadTruck <- LoadPackage{
		TruckName: f.TargetTruck.Name,
		Parcel:    f.Content,
	}

	res := <-f.TargetTruck.LoadTruck

	if res.Loaded {
		f.Content = nil
		f.TargetTruck = nil
	} else {
		fmt.Println("Waiting...")
	}
}

func (f *Forklift) IsNextToTarget(target Position) bool {
	if f.Position.X == target.X+1 || f.Position.X == target.X-1 {
		return true
	}

	if f.Position.Y == target.Y+1 || f.Position.Y == target.Y-1 {
		return true
	}

	return false
}
