package handler

import (
	"errors"
	"hospital-a/models/dtos"
	"hospital-a/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PatientHandler struct {
	paitentService *service.PatientService
}

func NewPatientHandler(paitentService *service.PatientService) *PatientHandler {
	return &PatientHandler{paitentService: paitentService}
}

func (h *PatientHandler) GetAllPatients(ctx *gin.Context) {
	var filter dtos.PatientFilter
	hospital_name := ""

	if err := ctx.ShouldBindQuery(&filter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	patients, patientErr := h.paitentService.GetAllPatients(hospital_name, &filter)

	if patientErr != nil {
		ctx.AbortWithStatusJSON(patientErr.Code, patientErr)

		return
	}

	ctx.JSON(http.StatusOK, patients)
}

func (h *PatientHandler) GetPatientById(ctx *gin.Context) {
	patientId := ctx.Param("id")
	hospital_name := ""

	patient, patientErr := h.paitentService.GetPatientById(hospital_name, patientId)
	if patientErr != nil {
		ctx.AbortWithStatusJSON(patientErr.Code, patientErr)

		return
	}

	ctx.JSON(http.StatusOK, patient)
}

func (h *PatientHandler) DeletePatient(ctx *gin.Context) {
	patientId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User ID not valid"})

		return
	}

	deleteError := h.paitentService.DeletePatient(patientId)

	if deleteError != nil {
		ctx.AbortWithStatusJSON(deleteError.Code, deleteError)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Patient deleted"})
}

func (h *PatientHandler) CreatePatient(ctx *gin.Context) {
	var createPatientRequest dtos.CreatePatientRequest

	if err := ctx.ShouldBindJSON(&createPatientRequest); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make(map[string]string)
			for _, fe := range ve {
				out[fe.Field()] = msgForTag(fe)
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})

			return
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	CreatePatientResponse, createErr := h.paitentService.CreatePatient(&createPatientRequest)

	if createErr != nil {
		ctx.AbortWithStatusJSON(createErr.Code, createErr)

		return
	}

	ctx.JSON(http.StatusCreated, CreatePatientResponse)
}
