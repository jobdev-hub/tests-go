package user_repository

import (
	"api-with-mongodb/configs/mongodb"
	"api-with-mongodb/models"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	collection = mongodb.GetCollection("users")
	ctx        = context.Background()
)

func FindMany() (models.Users, error) {

	var users models.Users
	filter := bson.D{}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return users, err
	}

	for cursor.Next(ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return users, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func FindOneByID(id string) (models.User, error) {

	var user models.User
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	filter := bson.D{{"_id", objectID}}
	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil

}

func InsertOne(user models.User) error {

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func UpdateOne(userId string, update bson.M) error {
	oid, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.M{"_id": oid}

	var err error
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func DeleteOne(userId string) error {

	var err error
	var oid primitive.ObjectID
	oid, err = primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": oid}
	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func CheckEmailUnique(email string, userID string) error {

	var user models.User

	filter := bson.M{"email": email}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return err
	}

	if count == 0 {
		return nil
	} else {
		err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
		if err != nil {
			return err
		}
		if user.ID != nil && user.ID.Hex() != userID {
			return errors.New("email already exists to userID: " + user.ID.Hex())
		}
	}

	return nil
}
