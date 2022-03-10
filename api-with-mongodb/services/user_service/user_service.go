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

	err := models.CheckFieldsToInsert(user)
	if err != nil {
		return err
	}

	user.CreatedAt = time.Now()

	err = user_repository.InsertOne(user)
	if err != nil {
		return err
	}

	return nil
}

func UpdateOne(user models.User, userId string) error {

	update, err := models.CheckFieldsToUpdate(user)
	if err != nil {
		return err
	}

	update["$set"].(bson.M)["updated_at"] = time.Now()

	_, err = user_repository.FindOneByID(userId)
	if err != nil {
		return err
	}

	err = user_repository.UpdateOne(userId, update)
	if err != nil {
		return err
	}

	return nil
}

func DeleteOne(userId string) error {

	_, err := user_repository.FindOneByID(userId)
	if err != nil {
		return err
	}

	err = user_repository.DeleteOne(userId)
	if err != nil {
		return err
	}

	return nil
}
