package contacts

import (
	"github.com/jinzhu/gorm"
	"errors"
	"github.com/golang/glog"
)

//Service : Contact service to interact directly with business process
type Service struct {
	db 	*gorm.DB
}

//NewService : instantiate new contact service
func NewService(Db *gorm.DB) (*Service, error) {

	if Db == nil {
		glog.Error("failed to intantiate Service , Db instance is null")
		return nil, errors.New("failed to intantiate Service , Db instance is null")
	}

	return &Service{db: Db}, nil
}