package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/iamaloneforever/GoTraining/db"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"apikey"`
}

func databaseUserToUser(dbUserr db.User) User {
	return User{
		ID:        dbUserr.ID,
		Name:      dbUserr.Name,
		UpdatedAt: dbUserr.UpdatedAt,
		ApiKey:    dbUserr.ApiKey,
		CreatedAt: dbUserr.CreatedAt,
	}
}
