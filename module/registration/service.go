package registration

import (
	"errors"

	"github.com/golang/glog"

	model "github.com/Friend-Management/module/registration/model"

	"github.com/jinzhu/gorm"
)

//Service : User service to interarct directly with business process
type Service struct {
	db *gorm.DB
}

//NewService : intantiate new service
func NewService(Db *gorm.DB) (*Service, error) {

	if Db == nil {
		return nil, errors.New("failed to intantiate Service , Db intance is null")
	}
	return &Service{db: Db}, nil
}

//CreateUser : Registration Service to create new user
func (s *Service) CreateUser(email string) (bool, error) {

	var user model.User

	tx := s.db.Begin()

	if tx.Error != nil {
		glog.Errorf("Failed to begin transaction when creating User : %s", tx.Error)
		return false, tx.Error
	}

	if tx.First(&user).RecordNotFound() {
		if err := tx.Create(&user); err != nil {
			glog.Errorf("Error when saving user : %s", err.Error)
			tx.Rollback()
		}
	} else {
		return false, errors.New("Email already registered")
	}

	return true, nil
}
