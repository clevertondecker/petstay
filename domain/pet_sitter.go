package domain

import "fmt"

type PetSitter struct {
	ID              int
	Name            string
	Rating          float64 // Avaliação do pet sitter (de 1 a 5)
	ExperienceLevel string  // Nível de experiência: "Beginner", "Intermediate", "Expert"
}

func (p PetSitter) Display() string {
	return fmt.Sprintf("PetSitter ID: %d, Name: %s, Rating: %.1f, Experience Level: %s",
		p.ID, p.Name, p.Rating, p.ExperienceLevel)
}
