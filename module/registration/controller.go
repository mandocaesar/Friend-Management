package registration

import (
	"net/http"

	"github.com/gin-gonic/gin/binding"

	messages "github.com/Friend-Management/module/registration/messages"
	"github.com/gin-gonic/gin"
)

//Controller : registration controller to manage user business process
type Controller struct {
	UserService *Service
}

//NewController : Instantiate new Controller instance
func NewController(service *Service) *Controller {
	return &Controller{UserService: service}
}

//RegisterUser function to register user email
func (c *Controller) RegisterUser(r *gin.Context) {
	var req messages.RequestMessage
	if err := r.ShouldBindWith(&req, binding.JSON); err != nil {
		success, err := c.UserService.CreateUser(req.Email)
		if success {
			r.JSON(http.StatusOK, gin.H{"result": "success"})
		}

		r.JSON(http.StatusOK, gin.H{"result": "success", "error": err.Error()})
	}
	r.JSON(http.StatusOK, gin.H{"result": "failed"})

}
