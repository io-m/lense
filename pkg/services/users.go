package services

import (
	// "fmt"
	"log"
	"net/http"

	"github.com/io-m/lenses/pkg/models"
	"github.com/io-m/lenses/pkg/models/storages"
	"github.com/io-m/lenses/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

const (
	pepper    = "my-app-level-pepper"
	secretKey = "my-ultimately-secret-jwt-key"
)

// UserService is an interface that defines methods passed to the storage queries
type UserService interface {
	Save(user models.User) (*models.User, *utils.Response)
	GetOne(id string) (*models.User, *utils.Response)
	GetByEmail(email string) *utils.Response
	Update(id uint) (string, *utils.Response)
	Delete(id uint) (string, *utils.Response)
	GetAll() ([]*models.User, *utils.Response)
	Validator
}

// Validator is an interface that defines methods for validating and authenticating users
type Validator interface {
	Validate(user *models.User) (*models.User, *utils.Response)
	Authenticate(user *models.User, password string) (string, *utils.Response)
}

type userService struct {
	// hmac utils.HMAC
}

// NewUserService is a constructor function for making new instances of UserService interface
func NewUserService() UserService {
	return &userService{}
}

var s = storages.NewStorage()

// ===========================
// Implementing interface on userSerive struct

func (*userService) GetOne(id string) (*models.User, *utils.Response) {
	var user = &models.User{}
	user, err := s.GetOne(id)
	if err != nil {
		log.Printf("ERROR: %s CODE: %d", err.Msg, err.StatusCode)
		return nil, err
	}
	return user, nil
}

func (*userService) GetByEmail(email string) *utils.Response {
	err := s.GetByEmail(email)
	if err != nil {
		log.Printf("ERROR: %s CODE: %d", err.Msg, err.StatusCode)
		return err
	}
	return nil
}

func (us *userService) Save(user models.User) (*models.User, *utils.Response) {
	passwordBytes := []byte(user.Password + pepper)
	// fmt.Println(passwordBytes)
	// Hashing password
	bytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return nil, utils.Back(http.StatusInternalServerError, err.Error())
	}
	user.Hash = string(bytes)
	user.Password = ""
	newID, errResp := s.Save(user)
	if errResp != nil {
		log.Printf("ERROR: %s CODE: %d", errResp.Msg, errResp.StatusCode)
		return nil, utils.Back(errResp.StatusCode, errResp.Msg)
	}
	user.ID = newID
	return &user, nil
}

func (*userService) GetAll() ([]*models.User, *utils.Response) {
	// allUsers := []*models.User{}
	allUsers, err := s.GetAll()
	if err != nil {
		log.Printf("ERROR: %s CODE: %d", err.Msg, err.StatusCode)
		return nil, err
	}
	return allUsers, nil
}
func (*userService) Update(id uint) (string, *utils.Response) {
	success, err := s.Update(id)
	if err != nil {
		log.Printf("ERROR: %s CODE: %d", err.Msg, err.StatusCode)
		return "", err
	}
	return success, nil
}
func (*userService) Delete(id uint) (string, *utils.Response) {
	success, err := s.Delete(id)
	if err != nil {
		log.Printf("ERROR: %s CODE: %d", err.Msg, err.StatusCode)
		return "", err
	}
	return success, nil
}

func (us *userService) Validate(user *models.User) (*models.User, *utils.Response) {
	if user == nil {
		return nil, utils.Back(http.StatusBadRequest, "Fill out your info")
	}
	if len(user.Password) < 6 {
		return nil, utils.Back(http.StatusBadRequest, "Your password is too short. At least 6 chars are required")
	}
	return user, nil
}
func (us *userService) Authenticate(user *models.User, password string) (string, *utils.Response) {
	if err := us.GetByEmail(user.Email); err != nil {
		return "", utils.Back(http.StatusNotFound, "User is not found by an email")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(password+pepper)); err != nil {
		return "", utils.Back(http.StatusUnauthorized, err.Error())
	}
	// Here comes generating jwt code bcs user's password is compared with one in DB
	jwtToken, err := utils.JWToken(user.Email)
	if err != nil {
		return "", err
	}
	log.Printf("User %v is authenticated", jwtToken)
	return jwtToken, nil
}
