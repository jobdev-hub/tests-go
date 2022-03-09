package models

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	Roles     []string           `json:"roles"`
	Active    bool               `json:"active"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt *time.Time         `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type Users []*User

func CheckFields(user User) error {

	var err []string

	if user.Name == "" {
		err = append(err, "name")
	}

	if user.Email == "" {
		err = append(err, "email")
	}

	if user.Roles == nil {
		err = append(err, "roles")
	}

	if err != nil {
		return errors.New("[" + strings.Join(err, ", ") + "] is required")
	}

	return nil
}
