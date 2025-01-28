package controller

import (
	"fmt"
	"encoding/json"
	"net/http"
	"petstay/domain"
	"petstay/service"
)

// OwnerController manages operations related to Owners.
// It exposes methods to handle the creation of owners via HTTP API.
type OwnerController struct {
	Service *service.OwnerService
}

// NewOwnerController creates a new instance of OwnerController.
// This function initializes the controller with the provided owner service.
//
// Parameters:
// - service: the owner service responsible for the business logic.
//
// Returns:
// - A new instance of OwnerController initialized with the provided service.
func NewOwnerController(service *service.OwnerService) *OwnerController {
	return &OwnerController{Service: service}
}

// CreateOwner creates a new owner from an HTTP POST request.
// This function receives the owner data in JSON format, validates it, saves it to the database,
// and returns a success or error status.
//
// Parameters:
// - w: The ResponseWriter used to write the HTTP response.
// - r: The HTTP request containing the owner data in JSON format.
//
// Returns:
// - On success, it returns HTTP status 201 (Created) with the owner ID.
// - On failure, it returns an appropriate error status (400 or 500).
func (c *OwnerController) CreateOwner(w http.ResponseWriter, r *http.Request) {
	var owner domain.Owner

	// Decoding the JSON request body into the owner object
	err := json.NewDecoder(r.Body).Decode(&owner)
	if err != nil {
		http.Error(w, "Error reading the request body", http.StatusBadRequest)
		return
	}

	// Saving the owner using the service
	err = c.Service.Save(&owner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Returning the ID of the newly created owner
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Owner created successfully! ID: %d", owner.ID)))
}
