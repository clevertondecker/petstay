package main

import (
	"fmt"
	"log"
	"petstay/domain"
	"petstay/infra"
	"petstay/service"
)

func main() {
	// Initialize the database connection.
	infra.InitDB()
	db := infra.GetDB()

	// Create the repository and service.
	ownerRepo := infra.NewOwnerRepository(db)
	ownerService := service.OwnerService{
		Repo: ownerRepo,
	}

	// Create the controller with the service.
	ownerController := controller.NewOwnerController(&ownerService)

	// Configure the route to create an Owner.
	http.HandleFunc("/owners", ownerController.CreateOwner)

	// Start the HTTP server.
	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
