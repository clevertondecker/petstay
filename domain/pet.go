package domain

import "fmt"

type SpeciesType int

const (
	Dog  SpeciesType = iota // 0
	Cat                     // 1
	Bird                    // 2
	Fish                    // 3
)

func (s SpeciesType) String() string {
	switch s {
	case Dog:
		return "Dog"
	case Cat:
		return "Cat"
	case Bird:
		return "Bird"
	case Fish:
		return "Fish"
	default:
		return "Unknown"
	}
}

type Pet struct {
	ID           int
	Name         string
	Species      SpeciesType
	Age          int
	SpecialNeeds string
}

func (p Pet) Display() string {
	return fmt.Sprintf("Pet ID: %d, Name: %s, Species: %s, Age: %d, Special Needs: %s",
		p.ID, p.Name, p.Species, p.Age, p.SpecialNeeds)
}
