package infra

import (
	"database/sql"
	"fmt"
	"petstay/domain"
)

// Definindo o repositório para manipulação do banco
type OwnerRepository struct {
	DB *sql.DB
}

// Criar um novo repositório de Owner
func NewOwnerRepository(db *sql.DB) domain.OwnerRepository {
	return &OwnerRepository{DB: db}
}

// Implementação do método Save para o repositório
func (r *OwnerRepository) Save(owner *domain.Owner) error {
	// Preparando a consulta de inserção
	query := "INSERT INTO owners (name, email, phone, address) VALUES (?, ?, ?, ?)"
	result, err := r.DB.Exec(query, owner.Name, owner.Email, owner.Phone, owner.Address)
	if err != nil {
		return fmt.Errorf("erro ao salvar o owner no banco: %v", err)
	}

	// Obtendo o ID gerado
	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("erro ao recuperar o ID do owner: %v", err)
	}

	// Atribuindo o ID ao objeto owner
	owner.ID = int(id)
	return nil
}
