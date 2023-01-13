package entrego

type Truck struct {
	Name          string
	Position      Position
	MaxWeight     int
	Delay         int
	Load          []*Parcel
	Available     bool
	RemainingTime int
	LoadTruck     chan LoadPackage
}

type LoadPackage struct {
	TruckName string
	Parcel    *Parcel
	Loaded    bool
}

func (t Truck) totalWeight() int {
	weight := 0

	for _, parcel := range t.Load {
		weight += parcel.Weight
	}

	return weight
}

func (t *Truck) load(loadPackage LoadPackage) bool {
	if !t.Available {
		return false
	}

	totalWeight := t.totalWeight()
	sumWeight := totalWeight + loadPackage.Parcel.Weight

	if totalWeight >= t.MaxWeight || sumWeight > t.MaxWeight {
		t.Available = false
		t.RemainingTime = t.Delay

		return false
	}

	t.Load = append(t.Load, loadPackage.Parcel)

	if sumWeight == t.MaxWeight {
		t.Available = false
		t.RemainingTime = t.Delay
	}

	return true
}

func (t *Truck) InitTruck(globalQuit chan struct{}) {
	for {
		select {
		case loadPackage := <-t.LoadTruck:
			res := t.load(loadPackage)
			t.LoadTruck <- LoadPackage{Loaded: res}
		case <-globalQuit:
			return
		}
	}
}
