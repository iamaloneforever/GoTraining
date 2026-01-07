package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/iamaloneforever/GoTraining/db"
)

func (apiConf *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, http.StatusBadRequest,
			fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	if params.Name == "" {
		respondWithError(w, http.StatusBadRequest, "name is required")
		return
	}

	// Insert user
	user, err := apiConf.DB.CreateUser(r.Context(), db.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError,
			fmt.Sprintf("Can't create user: %v", err))
		return
	}

	// فقط پیام موفقیت بده
	respondWithJSON(w, http.StatusCreated, databaseUserToUser(user))
}
