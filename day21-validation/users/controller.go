package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{service: service}
}

func formatErrors(err error) map[string]string {
	errs := make(map[string]string)
	for _, e := range err.(validator.ValidationErrors) {
		errs[e.Field()] = e.Tag()
	}
	return errs
}

func (ctrl *Controller) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, ctrl.service.GetAll())
}

func (ctrl *Controller) CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Input validation error", "errors": formatErrors(err)})
		return
	}

	user, err := ctrl.service.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (ctrl *Controller) DeleteUser(c *gin.Context) {
	if err := ctrl.service.Delete(c.Param("id")); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

func (ctrl *Controller) PurgeSoftDeleted(c *gin.Context) {
	count := ctrl.service.PurgeSoftDeleted()
	c.JSON(http.StatusOK, gin.H{"purged": count})
}

func (ctrl *Controller) UpdateUser(c *gin.Context) {
	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": formatErrors(err)})
		return
	}

	user, err := ctrl.service.Update(c.Param("id"), input)
	if err == ErrNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err == ErrDeleted {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
