package drivers

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo(collection_name string) (*mongo.Collection, *mongo.Client, error) {

	// auth := options.Credential{
	// 	Username: os.Getenv("MONGO_USERNAME"),
	// 	Password: os.Getenv("MONGO_PASSWORD"),
	// }
	// clientOptions := options.Client().ApplyURI(os.Getenv("DB")).SetAuth(auth)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://muboy:muboy@cluster0.evqfa.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, nil, err
	}

	collection := client.Database("goapi").Collection(collection_name)

	return collection, client, nil
}
