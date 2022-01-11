package repository

import (
	"context"
	"go-api/drivers"
	"go-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListAllMember() ([]models.Member, error) {
	var members []models.Member

	db, client, err := drivers.ConnectMongo("member")
	if err != nil {
		return members, err
	}

	cur, err := db.Find(context.TODO(), bson.M{})
	if err != nil {
		return members, err
	}

	for cur.Next(context.Background()) {
		var member models.Member
		cur.Decode(&member)
		members = append(members, member)
	}

	err = client.Disconnect(context.Background())

	if err != nil {
		return members, err
	}
	return members, nil
}

func FindOneMember(id primitive.ObjectID) (models.Member, error) {
	var member models.Member

	db, client, err := drivers.ConnectMongo("member")
	if err != nil {
		return member, err
	}

	err = db.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&member)
	if err != nil {
		return member, err
	}

	err = client.Disconnect(context.Background())

	if err != nil {
		return member, err
	}
	return member, nil
}

func InsertMember(member models.Member) error {
	db, client, err := drivers.ConnectMongo("member")
	if err != nil {
		return err
	}
	_, err = db.InsertOne(context.Background(), member)
	if err != nil {
		return err
	}
	err = client.Disconnect(context.Background())

	if err != nil {
		return err
	}
	return nil
}

func UpdateMember(id primitive.ObjectID, member models.Member) error {
	db, client, err := drivers.ConnectMongo("member")
	if err != nil {
		return err
	}

	_, err = db.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": member})
	if err != nil {
		return err
	}

	err = client.Disconnect(context.Background())

	if err != nil {
		return err
	}
	return nil
}

func DeleteMember(id primitive.ObjectID) error {
	db, client, err := drivers.ConnectMongo("member")
	if err != nil {
		return err
	}

	_, err = db.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}

	err = client.Disconnect(context.Background())

	if err != nil {
		return err
	}
	return nil
}
