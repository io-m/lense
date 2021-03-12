package storages

import (
	"log"
	"net/http"
	// "database/sql"
	// "time"
	"context"

	"github.com/io-m/lenses/pkg/models"
	"github.com/io-m/lenses/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)

const (
// 	saveStmt = "INSERT INTO users(name, email, password_hash) VALUES($1, $2, $3) RETURNING id;"
// 	getAllStmt = "SELECT * FROM users;"
// 	getByEmailStmt = "SELECT * FROM users WHERE email=$1;"
)

var allUsers = []*models.User{}


type storage struct{}

// NewStorage is constructor function for making new instances of
// storage interface with all of the method needed for interacting
// with database
func NewStorage() Storage {
	return &storage{}
}

// =============================
// Implementing interface functions

var (
	ctx = context.Background()
	// mongoClient = &MongoClient{}
) 




func (*storage) GetOne(id string) (*models.User, *utils.Response) {
	user := &models.User{}
	// mongoid, _ := primitive.ObjectIDFromHex(id)
	// err := mongocollection.FindOne(ctx, models.User{ID: mongoid}).Decode(user)
	// if err != nil {
	// 	return nil, utils.Back(http.StatusInternalServerError, "Can not find specified user")
	// }
	return user, nil
}

func (*storage) GetByEmail(theEmail string) *utils.Response {
	
	return nil	
}
func (*storage) Save(user models.User) (primitive.ObjectID, *utils.Response) {
	result, err := Collection.InsertOne(ctx, bson.D{
		{Key: "name", Value: user.Name},
		{Key: "email", Value: user.Email},
		{Key: "hash", Value: user.Hash},
	})
	if err != nil {
		return primitive.ObjectID{}, utils.Back(http.StatusInternalServerError, "Can not create new user")
	}
	newID := result.InsertedID.(primitive.ObjectID)
	log.Println(newID)
	return newID, nil
	// stmt, err := DB.Prepare(saveStmt)
	// if err != nil {
	// 	log.Println(err)
	// }

	// defer stmt.Close()
	// result, err := stmt.Exec(user.Name, user.Email, user.PasswordHash)
	// if err != nil {
	// 	return utils.Back(http.StatusInternalServerError, "Can not create new user")
	// }
	// if err != nil {
	// 	return utils.Back(http.StatusInternalServerError, "Can not create new user")
	// }
	// userid, err := result.LastInsertId() // does not work with postgres
	// if err != nil {
	// 	return utils.Back(http.StatusInternalServerError, err.Error())
	// }
	// log.Println(userid)
	// return nil
}
func (*storage) GetAll() ([]*models.User, *utils.Response) {
	cursor, err := Collection.Find(ctx, bson.M{})
	user := &models.User{}
	if err != nil {
		return nil, utils.Back(http.StatusInternalServerError, "Could not fetch all records")
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		if err := cursor.Decode(user); err != nil {
			return nil, utils.Back(http.StatusInternalServerError, "Could not fetch all records")
		}
		allUsers = append(allUsers, user)
	}
	// rows, err := DB.Query(getAllStmt)
	// if err != nil {
	// 	return nil, utils.Back(http.StatusInternalServerError, "Could not fetch all records")
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	user := &models.User{}
	// 	err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash)
	// 	if err != nil {
	// 		return nil, utils.Back(http.StatusInternalServerError, "Could not fetch all records")
	// 	}
	// 	allUsers = append(allUsers, user)
	// }
	return allUsers, nil
}
func (*storage) Update(id uint) (string, *utils.Response) {
	return "", nil

}
func (*storage) Delete(id uint) (string, *utils.Response) {
	return "", nil

}

// =========================
// HELPER FUNCTION FOR DB INTERACTIONS
// func first(db *gorm.DB, dest interface{}) *utils.Response {
// 	err := db.First(dest).Error
// 	if err == gorm.ErrRecordNotFound {
// 		return utils.Back(http.StatusNotFound, "User is not found")
// 	}
// 	return utils.Back(http.StatusInternalServerError, "Something went wrong")
// }
