package routes

import (
	"hospital-a/handler"

	"github.com/gin-gonic/gin"
)

func StaffRoute(router *gin.RouterGroup, staffHandler *handler.StaffHandler) {
	staff := router.Group("/staff")
	staff.POST("/create", staffHandler.CreateStaff)
	staff.POST("/login", staffHandler.StaffLogin)
}
