package cthulhu

import (
	"encoding/json"
	"github.com/ahonesa/new-character-repository/internal/model/cthulhu"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// GetCthulhuCharacterByID retrieves a character by its ID.
func GetCthulhuCharacterByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	repository := NewCthulhuCharacterRepository()
	character, err := repository.GetByID(id)
	if err != nil {
		http.Error(w, "Character not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(character)
}

// GetAllCthulhuCharacters retrieves all Cthulhu characters.
func GetAllCthulhuCharacters(w http.ResponseWriter, r *http.Request) {
	repository := NewCthulhuCharacterRepository()
	characters, err := repository.GetAll()
	if err != nil {
		http.Error(w, "Error fetching characters", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(characters)
}

// CreateCharacterHandler handles POST requests to create new Cthulhu characters.
func CreateCharacterHandler(w http.ResponseWriter, req *http.Request) {
	var cthulhuChar cthulhu.CthulhuCharacter
	if err := json.NewDecoder(req.Body).Decode(&cthulhuChar); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	// Set the Content-Type of response
	w.Header().Set("Content-Type", "application/json")

	// Assuming the OwnerId should be extracted from the authenticated user's context.
	// Here, it's set directly for simplicity.
	cthulhuChar.OwnerId = "ExtractedOwnerId"

	// Generate a new CharacterId if not provided. This is optional and depends on your application logic.
	if cthulhuChar.Character.CharacterId == "" {
		cthulhuChar.Character.CharacterId = uuid.New().String()
	}

	repo := NewCthulhuCharacterRepository()
	err := repo.Create(req.Context(), cthulhuChar.Character.CharacterId, &cthulhuChar)
	if err != nil {
		http.Error(w, "Failed to create character", http.StatusInternalServerError)
		return
	}

	// Respond with the ID of the newly created character
	response := map[string]string{"createdId": cthulhuChar.Character.CharacterId}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		// Log the error, as it's too late to send a new HTTP status code
		log.Printf("Error encoding response: %v", err)
	}
}

// UpdateCharacterHandler handles PUT requests to update existing Cthulhu characters.
func UpdateCharacterHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	characterId := vars["id"] // Extract CharacterId from the URL path.

	var cthulhuChar cthulhu.CthulhuCharacter
	if err := json.NewDecoder(req.Body).Decode(&cthulhuChar); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	// Ensure that the characterId from the path is used to update the correct character.
	// This is important if the characterId in the path and the one in the body might differ.
	cthulhuChar.Character.CharacterId = characterId

	repo := NewCthulhuCharacterRepository()
	if err := repo.Update(req.Context(), characterId, &cthulhuChar); err != nil {
		http.Error(w, "Failed to update character", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Character updated successfully"))

	if _, err := w.Write([]byte("Character updated successfully")); err != nil {
		// Log the error, as it's too late to send a new HTTP status code
		log.Printf("Error sending response: %v", err)
	}
}

// DeleteCharacterHandler handles DELETE requests to remove a Cthulhu character.
func DeleteCharacterHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	characterId := vars["id"] // Ensure this matches the URL parameter defined in your router setup.

	repo := NewCthulhuCharacterRepository()
	err := repo.Delete(req.Context(), characterId)
	if err != nil {
		http.Error(w, "Failed to delete character", http.StatusInternalServerError)
		log.Printf("Failed to delete character with ID %s: %v", characterId, err)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content is appropriate for successful DELETE operations with no response body.
}
