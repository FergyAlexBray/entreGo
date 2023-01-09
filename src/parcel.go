package entrego

var Colors = map[string]int{
	"BLUE":   500,
	"GREEN":  200,
	"YELLOW": 100,
}

type Parcel struct {
	Name     string
	Weight   int
	Position Position
}

func (c Core) FindParcel() Parcel {
	// Get the smallest parcel

	return Parcel{}
}
