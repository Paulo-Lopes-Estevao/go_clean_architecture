package entity

type Park struct {
	Name string
}

type ParkVague struct {
	Vague     string
	Limit     string
	Available string
}

type ParkVagueType struct {
	Type  string
	Price float32
	Time  int32
}
