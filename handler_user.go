package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/scaranaraa/rss_aggregator/internal/database"
)

func (apiConfig *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		resopondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
	}

	user, err := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})

	if err != nil {
		resopondWithError(w, 400, fmt.Sprintf("Couldnt create user: %v",err))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}