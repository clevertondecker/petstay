package infra

import (
	"database/sql"
	"fmt"
	"petstay/domain"
)

// Defining the repository for database manipulation.
type OwnerRepository struct {
	DB *sql.DB
}

// Create a new Owner repository.
func NewOwnerRepository(db *sql.DB) domain.OwnerRepository {
	return &OwnerRepository{DB: db}
}

// Implementing the Save method for the repository.
func (r *OwnerRepository) Save(owner *domain.Owner) error {
	query := "INSERT INTO owners (name, email, phone, address) VALUES (?, ?, ?, ?)"
	result, err := r.DB.Exec(query, owner.Name, owner.Email, owner.Phone, owner.Address)
	if err != nil {
		return fmt.Errorf("erro ao salvar o owner no banco: %v", err)
	}

	// Getting the generated ID.
	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("erro ao recuperar o ID do owner: %v", err)
	}

	// Assigning the ID to the owner object.
	owner.ID = int(id)
	return nil
}
