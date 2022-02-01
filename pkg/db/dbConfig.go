package db

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	MongoUri string
}

type Tech []struct {
	Path  string
	Title string
}

var (
	projectInstance *mongo.Collection
	// techInstance    *mongo.Collection
	level0Instance *mongo.Collection
	client         *mongo.Client
)

func (s *Config) Db_init() {
	client, _ = mongo.Connect(context.TODO(), options.Client().ApplyURI(s.MongoUri))
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
func getInstance(value string) *mongo.Collection {
	var once sync.Once
	switch value {
	case "project":
		if projectInstance == nil {
			once.Do(func() {
				projectInstance = client.Database("portfolio").Collection("projects")
			})
		}
		return projectInstance
	case "level0":
		if level0Instance == nil {
			once.Do(func() {
				level0Instance = client.Database("portfolio").Collection("level0")
			})
		}
		return level0Instance
	default:
		return nil
	}

}
