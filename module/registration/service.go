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

//NewService : instantiate new service
func NewService(Db *gorm.DB) (*Service, error) {

	if Db == nil {
		return nil, errors.New("failed to instantiate Service , Db instance is null")
	}
	return &Service{db: Db}, nil
}

//CreateUser : Registration Service to create new user
func (s *Service) CreateUser(email string) (bool, error, uint) {

	var user model.User

	user.ID = 0
	user.Email = email

	tx := s.db.Begin()

	if tx.Error != nil {
		glog.Errorf("Failed to begin transaction when creating User : %s", tx.Error)
		return false, tx.Error, 0
	}

	if !tx.Where("Email = ?", email).RecordNotFound() {
		if err := tx.Create(&user).Error; err != nil {
			glog.V(2).Infof("Error when saving user : %s", err)
			tx.Rollback()
			return false, errors.New("Error when saving user :" + err.Error()), 0

		}
		tx.Commit()
		s.db.First(&user)
		return true, nil, user.ID

	} else {
		return false, errors.New("Email already registered"), 0
	}

}
