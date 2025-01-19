package service

import (
	"fmt"
	"petstay/domain"
)

type OwnerService struct {
	Repo domain.OwnerRepository
}

// Create a new owner
func (s *OwnerService) Save(owner *domain.Owner) error {
	if owner.Name == "" || owner.Email == "" {
		return fmt.Errorf("campos obrigatórios estão vazios: Name e Email são obrigatórios")
	}

	// Validate the email format.
	if len(owner.Email) < 5 || !contains(owner.Email, "@") {
		return fmt.Errorf("email inválido: %s", owner.Email)
	}

	// Call the save method.
	if err := s.Repo.Save(owner); err != nil {
		return fmt.Errorf("erro ao salvar o owner: %v", err)
	}

	fmt.Printf("Owner criado com sucesso! ID: %d\n", owner.ID)
	return nil
}

// Função auxiliar para verificar a presença de uma substring no email
func contains(str, substr string) bool {
	return len(str) >= len(substr) && len(substr) > 0 && (str[:len(substr)] == substr || contains(str[1:], substr))
}
