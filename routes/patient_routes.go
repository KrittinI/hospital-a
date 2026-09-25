package routes

import (
	"hospital-a/handler"

	"github.com/gin-gonic/gin"
)

func PatientRoute(router *gin.RouterGroup, patientHandler *handler.PatientHandler) {
	patient := router.Group("/patient")
	patient.POST("/", patientHandler.CreatePatient)

	search := patient.Group("/search")
	search.Use()
	search.GET("/", patientHandler.GetAllPatients)
	search.GET("/:id", patientHandler.GetPatientById)
}
