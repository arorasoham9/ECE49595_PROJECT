package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type App struct {
	ID             primitive.ObjectID
	Name           *string
	PermisionLevel *string // "Admin", "User" or 0,1
}
