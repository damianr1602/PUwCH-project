package dao

import (
	"context"
	"log"

	logger "github.com/damianr1602/chmuryrest/logging"

	. "github.com/damianr1602/chmuryrest/gen/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MoviesDAO struct {
	Server   string
	Database string
}

var db *mongo.Database
var ctx context.Context = context.TODO()

// mongo collection
const (
	COLLECTION = "movies"
)

// Connect Establish a connection to database
func (m *MoviesDAO) Connect() {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.Server))
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	db = client.Database(m.Database)

}

// FindAll find list of movies
func (m *MoviesDAO) FindAll(since int32, limit int32) ([]*Movie, error) {
	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(since))
	res, _ := db.ListCollectionNames(ctx, bson.D{{}})
	logger.Log.Println("FindAll findOptions: ", res)

	var movies []*Movie
	cur, err := db.Collection(COLLECTION).Find(ctx, bson.D{{}}, findOptions)
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var movie Movie
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, &movie)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	logger.Log.Println("TEST-dao", movies)

	// Close the cursor once finished
	cur.Close(ctx)
	return movies, err
}

// FindbyID find a movie by its id
func (m *MoviesDAO) FindbyID(id string) (Movie, error) {
	var movie Movie
	objID, _ := primitive.ObjectIDFromHex(id)
	err := db.Collection(COLLECTION).FindOne(ctx, bson.D{{"_id", objID}}).Decode(&movie)
	return movie, err
}

// Create a movie into database
func (m *MoviesDAO) Create(movie Movie) (string, error) {
	inserted, err := db.Collection(COLLECTION).InsertOne(ctx, &movie)
	return inserted.InsertedID.(string), err
}

// Delete an existing movie
func (m *MoviesDAO) Delete(movie Movie) error {
	objID, _ := primitive.ObjectIDFromHex(movie.ID)

	_, err := db.Collection(COLLECTION).DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

// Update an existing movie
func (m *MoviesDAO) Update(movie Movie) error {
	objID, _ := primitive.ObjectIDFromHex(movie.ID)

	_, err := db.Collection(COLLECTION).UpdateOne(ctx, bson.M{"_id": objID}, &movie)
	return err
}