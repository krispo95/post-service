package db

import (
	"context"
	"fmt"
	"krispogram-grpc/pb"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DbInteractor struct {
	Url    string
	client *mongo.Client
}

func NewDbInteractor(url string) *DbInteractor {
	actor := DbInteractor{
		Url: url,
	}
	return &actor
}

func (d *DbInteractor) Connect() {
	// Set client options
	clientOptions := options.Client().ApplyURI(d.Url)
	var err error
	// Connect to MongoDB
	d.client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = d.client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}
func (d *DbInteractor) Disconnect() {
	// Close the connection once no longer needed
	err := d.client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection to MongoDB closed.")
	}
}

func (d *DbInteractor) Create(post *pb.Post) error {
	// Get a handle for your collection
	collection := d.client.Database("krispogram").Collection("posts")

	// Insert a single document
	insertResult, err := collection.InsertOne(context.TODO(), post)
	if err != nil {
		return err
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return err
}

func (d *DbInteractor) GetById(req *pb.GetPostByIdReq) (*pb.Post, error) {
	// Get a handle for your collection
	collection := d.client.Database("krispogram").Collection("posts")

	filter := bson.D{{"Id", req.Id}}
	var result pb.Post
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Found a single document: %+v\n", result)
	return &result, nil
}

func (d *DbInteractor) GetByAuthorId(req *pb.GetPostsByAuthorIdReq) (*pb.GetPostsByAuthorIdResp, error) {
	// Get a handle for your collection
	collection := d.client.Database("krispogram").Collection("posts")
	findOptions := options.Find()
	findOptions.SetLimit(20)
	filter := bson.D{{"AuthorId", req.AuthorId}}
	var result pb.GetPostsByAuthorIdResp
	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	// Iterate through the cursor
	for cur.Next(context.TODO()) {
		var elem pb.Post
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		result.Posts = append(result.Posts, &elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	if err := cur.Close(context.TODO()); err != nil {
		return nil, err
	}
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", result)

	return &result, nil
}
