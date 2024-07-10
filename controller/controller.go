package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Ankit-692/API/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName string = "Cluster0"
const colName string = "Watchlist"

var collection *mongo.Collection

//initalising Collection
func init() {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	//Client Options
	opts := options.Client().ApplyURI("mongodb+srv://bhartiankit692:Atlas@123//@cluster0.z5aoibm.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0").SetServerAPIOptions(serverAPI)

	//Making Connection
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB Connection Successfull")

	collection = client.Database(dbName).Collection(colName)
}


//Helper Functions

func insertOneMovie(movie model.Netflix){
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err!= nil {
		log.Fatal(err)
	}
	
	fmt.Println("Inserted Movie with Id",inserted.InsertedID)
}

func updateOneMovie(movieId string){
	id,_:=primitive.ObjectIDFromHex(movieId)
	filter:=bson.M{"_id":id}
	update:=bson.M{"$set":bson.M{"watched":true}}

	result, _:=collection.UpdateOne(context.Background(),filter,update)

	fmt.Println("updated watched", result.ModifiedCount)
}

func deleteOneMovie(movieId string){
	id,_:=primitive.ObjectIDFromHex(movieId)
	filter:=bson.M{"_id":id}

	result, _:=collection.DeleteOne(context.Background(),filter)

	fmt.Println("deleted",result)
}

func deleteAll(){
	result ,_:=collection.DeleteMany(context.Background(),bson.D{{}})
	fmt.Println("deleted All", result.DeletedCount)
}


func getAllMovies() []primitive.M{
	curr,_:=collection.Find(context.Background(),bson.D{{}})

	var movies []primitive.M

	for curr.Next(context.Background()){
		var movie bson.M
		err:= curr.Decode(&movie)
		if err!=nil{
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer curr.Close(context.Background())
	return movies
}



//Actual Controllers

func GetAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-Type","application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkWatched(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	parms := mux.Vars(r)
	updateOneMovie(parms["id"])
	json.NewEncoder(w).Encode(parms["id"])
}


func DeleteONeMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	parms := mux.Vars(r)
	deleteOneMovie(parms["id"])
	json.NewEncoder(w).Encode(parms["id"])
}

func DeleteAll(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	deleteAll()

	json.NewEncoder(w).Encode("Success")
}