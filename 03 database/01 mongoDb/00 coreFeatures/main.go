package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type product struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Name      string        `json:"name" bson:"name"`
	Price     int           `json:"price" bson:"price"`
	Available bool          `json:"available" bson:"available"`
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Connecting to MongoDb")
	session, e := mgo.Dial("mongodb://localhost")
	if e != nil {
		log.Fatalf("Failed to dial to MongoDb: %v", e)
	}
	log.Println("Connecting established to MongoDb")

	database := "shop"
	collection := "products"
	data := product{
		ID:        bson.NewObjectId(),
		Name:      "Nintendo Switch",
		Price:     309.00,
		Available: true,
	}

	printAllDatabasesWithCollections(session)
	clearCollection(session, database, collection)
	insertIntoCollection(session, database, collection, data)

	log.Println("Closing the MongoDb Connection!")
	session.Close()
}

func printAllDatabasesWithCollections(session *mgo.Session) {
	databaseNames, e := session.DatabaseNames()
	if e != nil {
		log.Fatalf("Failed to query for Database Names: %v", e)
	}

	for _, database := range databaseNames {
		fmt.Println("--- --- ---")
		fmt.Printf("Database: \t%s\n", database)

		collectionNames, e := session.DB(database).CollectionNames()
		if e != nil {
			log.Fatalf("Failed to query for Collection Names: %v", e)
		}

		fmt.Printf("Collections: \t%v\n", collectionNames)
		fmt.Println("--- --- ---")
	}
}

func clearCollection(session *mgo.Session, database string, collection string) {
	e := session.DB(database).C(collection).DropCollection()
	if e != nil {
		log.Fatalf("Failed to clear the Collection: %v", e)
	}
}

func insertIntoCollection(session *mgo.Session, database string, collection string, data product) {
	e := session.DB(database).C(collection).Insert(data)
	if e != nil {
		log.Fatalf("Failed to insert in Collection: %v", e)
	}
}
