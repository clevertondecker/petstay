package domain

import (
	"fmt"
	"time"
)

// Definindo a struct Booking (Reserva)
type Booking struct {
	ID        int
	Pet       Pet
	Owner     Owner
	PetSitter PetSitter
	StartDate time.Time
	EndDate   time.Time
	Status    string // "Pending", "Confirmed", "Completed", "Cancelled"
}

// Método para exibir informações da reserva
func (b Booking) Display() string {
	return fmt.Sprintf("Booking ID: %d, Pet: %s, Owner: %s, PetSitter: %s, Dates: %s to %s, Status: %s",
		b.ID, b.Pet.Name, b.Owner.Name, b.PetSitter.Name, b.StartDate.Format("2006-01-02"), b.EndDate.Format("2006-01-02"), b.Status)
}
