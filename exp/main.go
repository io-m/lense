package main

// import (
// 	"fmt"
// 	"log"

// 	"github.com/io-m/lenses/pkg/models"
// 	"github.com/io-m/lenses/pkg/models/storages"
// 	"golang.org/x/crypto/bcrypt"
// )

// var pepper = "my-pepper"
// var sqlStatement = `
// INSERT INTO users (name, email, passwordhash)
// VALUES ($1, $2, $3)`

// func main() {

// 	incommingUser := models.User{}
// 	postgres := storages.NewDB()
// 	db, err := postgres.Connect()
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	defer db.Close()

// 	incommingUser.Name = "Bla blaadsf"
// 	incommingUser.Email = "pisoaadcadsaj@outasddslsadfasook.hoho"
// 	incommingUser.Password = "Joasadscdsippasswordas"

// 	passwordBytes := []byte("blablabla" + pepper)
// 	// Hashing password
// 	bytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
// 	if err != nil {
// 		return
// 	}
// 	incommingUser.PasswordHash = string(bytes)

// 	fmt.Println(incommingUser)

// 	res := db.QueryRow(sqlStatement, &incommingUser.Name, &incommingUser.Email, &incommingUser.PasswordHash)
// 	_ = res.Scan(&incommingUser)
// 	fmt.Println(incommingUser)

// }
