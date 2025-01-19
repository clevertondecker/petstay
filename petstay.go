package main

import (
	"fmt"
	"log"
	"petstay/domain"
	"petstay/infra"
	"petstay/service"
)

func main() {
	// Inicializa a conexão com o banco de dados
	infra.InitDB()
	db := infra.GetDB()

	// Cria o repositório e o serviço
	ownerRepo := infra.NewOwnerRepository(db)
	ownerService := service.OwnerService{
		Repo: ownerRepo,
	}

	// Cria um novo Owner
	owner := &domain.Owner{
		Name:    "John Doe",
		Email:   "john.doe@example.com",
		Phone:   "123-456-7890",
		Address: "123 Main St",
	}

	// Tenta salvar o Owner
	if err := ownerService.Save(owner); err != nil {
		log.Fatalf("Erro ao salvar Owner: %v", err)
	}

	// Exibe a resposta
	fmt.Printf("Owner salvo com sucesso: %s\n", owner.Display())
}
