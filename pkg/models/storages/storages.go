package storages

import (
	"github.com/io-m/lenses/pkg/models"
	"github.com/io-m/lenses/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

// Storage is an interface that defines methods for
// interacting with database
type Storage interface {
	Save(user models.User) (primitive.ObjectID, *utils.Response)
	GetOne(id string) (*models.User, *utils.Response)
	GetByEmail(email string) *utils.Response
	Update(id uint) (string, *utils.Response)
	Delete(id uint) (string, *utils.Response)
	GetAll() ([]*models.User, *utils.Response)
}
