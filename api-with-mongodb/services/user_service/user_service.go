package user_service

import (
	"api-with-mongodb/models"
	"api-with-mongodb/repositories/user_repository"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func FindMany() (models.Users, error) {
	users, err := user_repository.FindMany()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func FindOneByID(id string) (models.User, error) {
	user, err := user_repository.FindOneByID(id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func InsertOne(user models.User) error {

	err := models.CheckFields(user)
	if err != nil {
		return err
	}

	user.CreatedAt = time.Now()
	user.Active = true

	err = user_repository.InsertOne(user)
	if err != nil {
		return err
	}

	return nil
}

func UpdateOne(user models.User, userId string) error {

	err := models.CheckFields(user)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"name":       user.Name,
			"email":      user.Email,
			"roles":      user.Roles,
			"updated_at": time.Now(),
		},
	}

	err = user_repository.UpdateOne(userId, update)
	if err != nil {
		return err
	}

	return nil
}

func InactivateOne(userId string) error {

	update := bson.M{
		"$set": bson.M{
			"active":     false,
			"updated_at": time.Now(),
		},
	}

	err := user_repository.UpdateOne(userId, update)
	if err != nil {
		return err
	}

	return nil
}

func DeleteOne(userId string) error {
	err := user_repository.DeleteOne(userId)
	if err != nil {
		return err
	}

	return nil
}
