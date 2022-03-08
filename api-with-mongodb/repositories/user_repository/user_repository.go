package user_repository

import (
	"api-with-mongodb/configs/mongodb"
	"api-with-mongodb/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var (
	collection = mongodb.GetCollection("users")
	ctx        = context.Background()
)

func Create(user models.User) error {

	var err error
	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func Read() (models.Users, error) {

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

func ReadByID(id string) (models.User, error) {

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

func Update(user models.User, userId string) error {

	var err error
	oid, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.M{"_id": oid}
	update := bson.M{
		"$set": bson.M{
			"name":       user.Name,
			"email":      user.Email,
			"updated_at": time.Now(),
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func Delete(userId string) error {

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
