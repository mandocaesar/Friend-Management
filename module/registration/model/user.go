package registration

import (
	"github.com/jinzhu/gorm"
)

//User : User data model
type User struct {
	gorm.Model
	Email string `gorm:"type:varchar(100);unique_index"`
}
