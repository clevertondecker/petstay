package main

import (
	"fmt"
	"petstay/domain"
	"time"
)

func main() {
	// Criando uma instância de Owner (Dono)
	owner := domain.Owner{
		ID:      1,
		Name:    "Carlos",
		Email:   "carlos@email.com",
		Phone:   "1234567890",
		Address: "Rua XYZ, 123",
	}

	// Criando uma instância de PetSitter (Cuidador)
	petSitter := domain.PetSitter{
		ID:              1,
		Name:            "João",
		Rating:          4.5,
		ExperienceLevel: "Expert",
	}

	// Criando uma instância de Pet
	pet := domain.Pet{
		ID:           1,
		Name:         "Rex",
		Species:      domain.Dog,
		Age:          5,
		SpecialNeeds: "None",
	}

	// Criando uma instância de Booking (Reserva)
	startDate := time.Now()
	endDate := startDate.Add(48 * time.Hour) // Exemplo de 2 dias de reserva

	booking := domain.Booking{
		ID:        1,
		Pet:       pet,
		Owner:     owner,
		PetSitter: petSitter,
		StartDate: startDate,
		EndDate:   endDate,
		Status:    "Confirmed",
	}

	// Exibindo as informações
	fmt.Println(owner.Display())
	fmt.Println(pet.Display())
	fmt.Println(petSitter.Display())
	fmt.Println(booking.Display())
}
