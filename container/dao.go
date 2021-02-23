package container

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type FIDataSourceDao struct {
	client       *mongo.Client
	indexer      *mongo.Database
	fIDataSource *mongo.Collection
}

func (f *FIDataSourceDao) SetUp() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection has been established.")
	f.client = client
	f.indexer = client.Database("indexer")
	f.fIDataSource = f.indexer.Collection("fIDataSource")
}

func (f *FIDataSourceDao) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := f.client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection has been closed.")
}

func (f *FIDataSourceDao) FetchAll() []FIDataSourceDto {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	cursor, err := f.fIDataSource.Find(ctx, bson.M{}, options.Find())
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = cursor.Close(ctx)
		if err != nil {
			log.Fatal("Unable to close the cursor.")
		}
	}()
	result := make([]FIDataSourceDto, 0)
	for cursor.Next(ctx) {
		result = append(result, ConvertToDto(&cursor.Current))
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}
