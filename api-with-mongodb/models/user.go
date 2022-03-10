package models

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	Roles     []string           `json:"roles"`
	Active    *bool              `json:"active"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt *time.Time         `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type Users []*User

func CheckFieldsToInsert(user User) error {

	var err []string

	if user.Name == "" {
		err = append(err, "name")
	}

	if user.Email == "" {
		err = append(err, "email")
	}

	if user.Roles == nil || len(user.Roles) == 0 {
		err = append(err, "roles")
	}

	if user.Active == nil {
		err = append(err, "active")
	}

	if err != nil {
		return errors.New("[" + strings.Join(err, ", ") + "] is required")
	}

	return nil

}

func CheckFieldsToUpdate(user User) (bson.M, error) {

	if user.Roles != nil && len(user.Roles) == 0 {
		return nil, errors.New("rules field needs at least 1 value to be updated")
	}

	update := bson.M{"$set": bson.M{}}
	count := 0

	if user.Name != "" {
		update["$set"].(bson.M)["name"] = user.Name
		count++
	}

	if user.Email != "" {
		update["$set"].(bson.M)["email"] = user.Email
		count++
	}

	if user.Roles != nil && len(user.Roles) > 0 {
		update["$set"].(bson.M)["roles"] = user.Roles
		count++
	}

	if user.Active != nil {
		update["$set"].(bson.M)["active"] = user.Active
		count++
	}

	if count == 0 {
		return nil, errors.New("no field identified to update, check body the request")
	}

	return update, nil

}
