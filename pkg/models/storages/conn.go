package storages

import (
	// "database/sql"
	"log"
	// "fmt"
	// "context"
	"github.com/io-m/lenses/conn"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "KierKeg44rD"
// 	dbname   = "lenslocked-dev"
// )

// var (
// 	cred = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
// 	//DB is sql.DB instance
// 	DB *sql.DB
// 	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

// )

// // MongoClient is struct for new mongo clients
// type MongoClient struct {
// 	Client *mongo.Client
// }

// // NewMongoClient is constructor for making new instances of mongodb clients
// func NewMongoClient() *MongoClient {
// 	return &MongoClient{
// 		Client: &mongo.Client{},
// 	}
// }
// // NewMongoDatabase is constructor for making new instances of mongodb databases
// func (cl *MongoClient) NewMongoDatabase(db string) *mongo.Database {
// 	return cl.Client.Database(db)
// }
// // NewCollection is function for making new mongodb collection
// func NewCollection (db *mongo.Database, name string) error {
// 	return db.Collection(name)
// // }

// // Connect is function for establishing connection with MongoDB
// func (cl *MongoClient) Connect (ctx context.Context, cred string) error {
// 	var err error
// 	cl.Client, err = mongo.Connect(ctx, options.Client().ApplyURI(cred))
// 	if err != nil {
// 		log.Println("DB connection error: ", err)
// 		return err
// 	}
// 	if err = cl.Client.Ping(ctx, readpref.Primary()); err != nil {
// 		log.Println("DB connection error [ping]: ", err)
// 		return err
// 	}
// 	return nil
// }

var(
	// Client is instance of mogodb client struct
	Client = &mongo.Client{}
	// MongoDB is mongo Db instance
	MongoDB = &mongo.Database{}
	// Collection is mongo db collection instance
	Collection = &mongo.Collection{}
	mongoconn = conn.Mongoconn
)


func init() {
	var (
		err error
		db = "lenses"
		collection = "users"
	) 
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoconn))
	if err != nil {
		log.Fatalln("DB connection error: ", err) 
	}
	if err = Client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalln("DB connection error [ping]: ", err) 
	}
	MongoDB = Client.Database(db)
	Collection = MongoDB.Collection(collection)
	if Collection == nil {
		if err := MongoDB.CreateCollection(ctx, collection); err != nil {
			log.Fatal(err)
		}
		log.Printf("Collection %s is created", collection)
	}
	log.Printf("COLLECTION : %s", collection)
}



	


