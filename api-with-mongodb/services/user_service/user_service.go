package user_service

import (
	"api-with-mongodb/models"
	"api-with-mongodb/repositories/user_repository"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

func FindMany() (int, models.Users, error) {

	if err := user_repository.IsConnected(); err != nil {
		return http.StatusInternalServerError, nil, err
	}

	users, err := user_repository.FindMany()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, users, nil

}

func FindOneByID(id string) (int, models.User, error) {

	if err := user_repository.IsConnected(); err != nil {
		return http.StatusInternalServerError, models.User{}, err
	}

	user, err := user_repository.FindOneByID(id)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return http.StatusNotFound, models.User{}, err

		} else if err.Error() == "the provided hex string is not a valid ObjectID" {
			return http.StatusBadRequest, models.User{}, err

		} else {
			return http.StatusInternalServerError, models.User{}, err
		}
	}

	return http.StatusOK, user, nil

}

func InsertOne(user models.User) (int, error) {

	if err := user_repository.IsConnected(); err != nil {
		return http.StatusInternalServerError, err
	}

	if err := models.CheckFieldsToInsert(user); err != nil {
		return http.StatusBadRequest, err
	}

	if err := models.CheckFieldsValues(user); err != nil {
		return http.StatusBadRequest, err
	}

	if err := user_repository.CheckEmailUnique(user.Email, ""); err != nil {
		return http.StatusBadRequest, err
	}

	user.CreatedAt = time.Now()

	err := user_repository.InsertOne(user)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func UpdateOne(user models.User, userId string) (int, error) {

	if err := user_repository.IsConnected(); err != nil {
		return http.StatusInternalServerError, err
	}

	if err := models.CheckFieldsValues(user); err != nil {
		return http.StatusBadRequest, err
	}

	if err := user_repository.CheckEmailUnique(user.Email, userId); err != nil {
		return http.StatusBadRequest, err
	}

	update, err := models.CheckFieldsToUpdate(user)
	if err != nil {
		return http.StatusBadRequest, err
	}

	update["$set"].(bson.M)["updated_at"] = time.Now()

	_, err = user_repository.FindOneByID(userId)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	err = user_repository.UpdateOne(userId, update)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, err
}

func DeleteOne(userId string) (int, error) {

	if err := user_repository.IsConnected(); err != nil {
		return http.StatusInternalServerError, err
	}

	_, err := user_repository.FindOneByID(userId)
	if err != nil {
		return http.StatusBadRequest, err
	}

	err = user_repository.DeleteOne(userId)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusNoContent, nil
}
