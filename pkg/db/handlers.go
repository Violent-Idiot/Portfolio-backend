package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// type dir struct {
// 	id   interface{}
// 	name string
// }
type fs struct {
	name string `bson:"name"`
	data string `bson:"data, omitempty"`
	dir  string `bson:"dir, omitempty"`
}

func GetProjects() Tech {
	cur, err := getInstance("project").Find(context.Background(), bson.D{})
	if err == mongo.ErrNoDocuments {
		log.Print("no documents present")
	}
	tech := &Tech{}
	if err := cur.All(context.TODO(), tech); err != nil {
		log.Print(err)
	}
	return *tech
}
func PostProject(name, data string) {
	doc := bson.M{"name": name, "data": data}
	result, err := getInstance("project").InsertOne(context.TODO(), doc)
	if err != nil {
		log.Print(err)
	}
	log.Printf("%+v", result)
}
func PostLevel0(name, data, dir string) *mongo.InsertOneResult {
	temp := fs{
		name: name,
		data: data,
		dir:  dir,
	}
	// doc := bson.M{{temp}}

	result, err := getInstance("level0").InsertOne(context.TODO(), temp)
	if err != nil {
		log.Print(err)
	}
	return result
}
