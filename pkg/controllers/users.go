package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/io-m/lenses/pkg/models"
	"github.com/io-m/lenses/pkg/services"
	"github.com/io-m/lenses/pkg/utils"
)

// UserController is an interface that defines methods passed to services
type UserController interface {
	GetOne(w http.ResponseWriter, r *http.Request)
	Save(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Middlewares
}

// Middlewares is an interface that defines methods for authentication
type Middlewares interface {
	WithAuth(next http.Handler) http.Handler
	// Valid(next http.HandlerFunc) http.HandlerFunc
}

type userController struct {
	// cookie *http.Cookie
}

// NewUserController is a contructor function for making new instances of UserController
func NewUserController() UserController {
	return &userController{}
}

var us = services.NewUserService()

// ===========================
// Implementing interface methods

func (uc *userController) GetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := utils.Params(r, "id", false)
	if err != nil {
		return
	}
	var user = &models.User{}
	user, errOne := us.GetOne(id.(string))
	if errOne != nil {
		log.Printf("ERROR: %s CODE: %d", errOne.Msg, errOne.StatusCode)
		return
	}
	if err := utils.ToJSON(w, user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (uc *userController) Save(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	newUser := models.User{}
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		return
	}
	savedUser, err := us.Save(newUser)
	if err != nil {
		http.Error(w, err.Msg, err.StatusCode)
		return
	}
	// Generating new jwtoken
	newToken, err := utils.JWToken(savedUser.Email)
	if err != nil {
		http.Error(w, err.Msg, err.StatusCode)
		return
	}
	// Here comes setting cookie, or write to response header with authUser credentials (id, email ...)
	cookie := &http.Cookie{
		Name:     "jwt-cookie",
		Value:    newToken,
		Expires:  time.Now().Add(time.Minute * 5),
		HttpOnly: true,
		MaxAge: 60 * 6,
	}
	w.Header().Set("Set-Cookie", fmt.Sprintf("%s=%s", cookie.Name, cookie.Value))
	if err := utils.ToJSON(w, savedUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (uc *userController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allUsers, err := us.GetAll()
	if err != nil {
		log.Printf("ERROR: %s CODE: %d", err.Msg, err.StatusCode)
		return
	}
	if err := utils.ToJSON(w, allUsers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (uc *userController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := utils.Params(r, "id", false)
	if err != nil {
		return
	}
	success, errUpdate := us.Update(id.(uint))
	if errUpdate != nil {
		log.Printf("ERROR: %s CODE: %d", errUpdate.Msg, errUpdate.StatusCode)
		return
	}
	log.Println(success)
	if err := utils.ToJSON(w, success); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (uc *userController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := utils.Params(r, "id", false)
	if err != nil {
		return
	}
	success, errDelete := us.Delete(id.(uint))
	if errDelete != nil {
		log.Printf("ERROR: %s CODE: %d", errDelete.Msg, errDelete.StatusCode)
		return
	}
	log.Println(success)
	if err := utils.ToJSON(w, success); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (uc *userController) Login(w http.ResponseWriter, r *http.Request) {
	newUser := models.User{}
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		return
	}
	
	jwtToken, err := us.Authenticate(&newUser, newUser.Password)
	if err != nil {
		http.Error(w, err.Msg, err.StatusCode)
		return
	}
	// Here comes setting cookie, or write to response header with authUser credentials (id, email ...)
	cookie := &http.Cookie{
		Name:     "jwt-cookie",
		Value:    jwtToken,
		Expires:  time.Now().Add(time.Minute * 5),
		HttpOnly: true,
	}
	w.Header().Set("Set-Cookie", fmt.Sprintf("%s=%s", cookie.Name, cookie.Value))
	w.Header().Set("Authorization", fmt.Sprintf("Authorized for %v time", cookie.Expires))
	type authMsg struct {
		success string
		err string
	}
	if err := utils.ToJSON(w, authMsg{"you are logged in","",}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (uc *userController) Logout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "jwt-cookie",
		Value:    "",
		HttpOnly: true,
		MaxAge: -1,
	}
	w.Header().Set("Set-Cookie", fmt.Sprintf("%s=%s", cookie.Name, cookie.Value))
}
// ============================================
// MIDDLEWARES
// WithAuth is middleware that checks if User is authorized to access certain content
func (uc *userController) WithAuth(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract jwt from cookie/headers and verify it with some method
		log.Println("Inside middleware : before")
		jwtCookie, err := r.Cookie("jwt-cookie")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "Please login, you are not authorized", http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		claimsToParseIn := &utils.CustomClaims{}
		tokenUnverified := jwtCookie.Value
		// Here we compare token from cookie/header with secret-key we used while creating it while login/signup
		token, err := jwt.ParseWithClaims(tokenUnverified, claimsToParseIn,
			func(t *jwt.Token) (interface{}, error) {
				return utils.JWTKey, nil
			})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(w, "Please login, you are not authorized", http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !token.Valid {
			http.Error(w, "Please login, you are not authorized", http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusOK)
		next.ServeHTTP(w, r)
		log.Println("After calling next")
	})
}

// *********** VALIDATE NEEDS TO BE IMPLEMENTED -> from us. Validate()

// func (uc *userController) Valid(next http.HandlerFunc) http.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		newUser := models.User{}
// 		if err := utils.FromJSON(r, newUser); err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 		}
// 		_, err := us.Validate(&newUser)
// 		if err != nil {
// 			http.Error(w, err.Msg, err.StatusCode)
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }

// ======================================================
// HELPER FUNCTIONS FOR CONTROLLERS
