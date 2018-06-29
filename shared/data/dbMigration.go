package data

import (
	registration "github.com/Friend-Management/module/registration/model"
	"github.com/Friend-Management/shared/config"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

//DbMigration : instance to hold dbmigration instance
type DbMigration struct {
	connection *gorm.DB
}

//NewDbMigration : intantiate new DBMigration instance
func NewDbMigration(cfg *config.Configuration) (*DbMigration, error) {
	dbFactory, err := NewDbFactory(cfg)

	if err != nil {
		glog.Errorf("%s", err)
		return nil, err
	}

	conn, err := dbFactory.DBConnection()

	if err != nil {
		glog.Errorf("%s", err)
		return nil, err
	}

	return &DbMigration{connection: conn}, nil
}

//Migrate : function to invoke gotm's automigrate
func (d *DbMigration) Migrate() (bool, error) {
	glog.Info("Start Database Migration")

	d.connection.AutoMigrate(
		&registration.User{},
	)

	glog.Info("Database migration finished")
	return true, nil
}
