package entrego

type Forklift struct {
	Name     string
	Position [2]int
	IsUsed   bool
	Content  Parcel
}
