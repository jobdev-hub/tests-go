package user_service

import (
	"api-with-mongodb/models"
	"api-with-mongodb/repositories/user_repository"
)

func Create(user models.User) error {

	err := models.VerifyFields(user)
	if err != nil {
		return err
	}

	err = user_repository.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func Read() (models.Users, error) {
	users, err := user_repository.Read()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func ReadByID(id string) (models.User, error) {
	user, err := user_repository.ReadByID(id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func Update(user models.User, userId string) error {

	err := user_repository.Update(user, userId)
	if err != nil {
		return err
	}

	return nil
}

func Delete(userId string) error {
	err := user_repository.Delete(userId)
	if err != nil {
		return err
	}

	return nil
}
