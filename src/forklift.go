package entrego

import (
	"fmt"
)

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
	// TODO: Only call when necessary, not every lap
	nextPositions := FindShortestPath(c.SpaceMap, f.Position, target)

	c.SpaceMap[f.Position.Y][f.Position.X] = 0 // Number representing empty
	c.SpaceMap[nextPositions[1].Y][nextPositions[1].X] = targetType

	f.Position = nextPositions[1]

	fmt.Println(f.Name, "GO", CoordinatesToString(f.Position))
}

func (f *Forklift) MoveTowardsParcel(c *Core) {
	f.move(c, f.TargetParcel.Position, 2)
}

func (f *Forklift) MoveTowardsTruck(c *Core) {
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
		fmt.Println(f.Name, "LEAVE", f.Content.Name, f.Content.Color)
		f.Content = nil
		f.TargetTruck = nil
	} else {
		fmt.Println(f.Name, "WAIT")
	}
}

func (f *Forklift) IsNextToTarget(target Position) bool {
	if (f.Position.Y == target.Y && f.Position.X == target.X+1) || (f.Position.Y == target.Y && f.Position.X == target.X-1) {
		return true
	}

	if (f.Position.X == target.X && f.Position.Y == target.Y+1) || (f.Position.X == target.X && f.Position.Y == target.Y-1) {
		return true
	}

	return false
}
