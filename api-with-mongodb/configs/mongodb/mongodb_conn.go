package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var (
	usr  = ""
	pwd  = ""
	host = "localhost"
	port = "27017"
	db   = "test"
	uri  = "mongodb://" + host + ":" + port + "/" + db
	//uri  = "mongodb://" + usr + ":" + pwd + "@" + host + ":" + port + "/" + db
)

func GetCollection(collection string) *mongo.Collection {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		panic(err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error())
	}

	return client.Database(db).Collection(collection)
}
