package registration

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Controller : registration controller to manage user business process
type Controller struct {
	UserService Service
}

//New : Instantiate new Controller instance
func New(service Service) *Controller {
	return &Controller{UserService: service}
}

//RegisterUser function to register user email
func (c *Controller) RegisterUser(r *gin.Context) {
	r.JSON(http.StatusOK, gin.H{"result": "success"})
}
