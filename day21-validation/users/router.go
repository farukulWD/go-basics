package users

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, ctrl *Controller) {
	r.GET("/users", ctrl.GetUsers)
	r.POST("/users", ctrl.CreateUser)
	r.DELETE("/users/purge", ctrl.PurgeSoftDeleted)
	r.DELETE("/users/:id", ctrl.DeleteUser)
	r.PATCH("/users/:id", ctrl.UpdateUser)
}
