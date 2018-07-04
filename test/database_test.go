package test

import (
	"testing"

	"github.com/Friend-Management/shared/data"
	"github.com/stretchr/testify/assert"
)

//TestDatabaseConnection : test database configuration load
func TestDatabaseConnection(t *testing.T) {
	cfg, err := LoadConfiguration()

	assert.Nil(t, err)
	assert.NotNil(t, cfg)
}

//TestDatabaseInstance : test database instance creation
func TestDatabaseInstance(t *testing.T) {
	cfg, err := LoadConfiguration()
	//test configuration
	assert.Nil(t, err)

	dbInstance, err := data.NewDbFactory(cfg)
	//test database factory
	assert.Nil(t, err)
	assert.NotNil(t, dbInstance)

	conn, err := dbInstance.DBConnection()
	//tst connection
	assert.Nil(t, err)
	assert.NotNil(t, conn)
}
