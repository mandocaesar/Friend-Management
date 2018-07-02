package registration

import (
	"net/http"

	"github.com/gin-gonic/gin/binding"

	messages "github.com/Friend-Management/module/registration/messages"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//Controller : registration controller to manage user business process
type Controller struct {
	UserService *Service
	db *gorm.DB
}

//NewController : Instantiate new Controller instance
func NewController(service *Service) *Controller {

	return &Controller{UserService: service}
}

//RegisterUser function to register user email
func (c *Controller) RegisterUser(r *gin.Context) {
	var req messages.RequestMessage
	 err := r.ShouldBindWith(&req, binding.JSON);
	 if err == nil {
		success, err, id := c.UserService.CreateUser(req.Email)
		if err == nil {
			r.JSON(http.StatusOK, gin.H{"result": success, "id":id})
			return
		}

		r.JSON(http.StatusBadRequest, gin.H{"result": success, "error": err.Error(), "id":0})
		return
	}else {
		 r.JSON(http.StatusBadRequest, gin.H{"result": "failed", "error": err.Error(), "id":0})
		 return
	 }

}
