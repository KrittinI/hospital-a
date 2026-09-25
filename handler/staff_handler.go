package handler

import (
	"errors"
	"fmt"
	"hospital-a/models/dtos"
	"hospital-a/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type StaffHandler struct {
	staffService *service.StaffService
}

func NewStaffHandler(staffService *service.StaffService) *StaffHandler {
	return &StaffHandler{staffService: staffService}
}

func (h *StaffHandler) GetStaff(ctx *gin.Context) {
	staffId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Staff ID not valid"})

		return
	}

	staff, staffErr := h.staffService.GetStaff(staffId)
	if staffErr != nil {
		ctx.AbortWithStatusJSON(staffErr.Code, staffErr)

		return
	}

	ctx.JSON(http.StatusOK, staff)
}

func (h *StaffHandler) DeleteUser(ctx *gin.Context) {
	staffId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User ID not valid"})

		return
	}

	deleteError := h.staffService.DeleteStaff(staffId)
	if deleteError != nil {
		ctx.AbortWithStatusJSON(deleteError.Code, deleteError)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Staff deleted"})
}

func (h *StaffHandler) CreateStaff(ctx *gin.Context) {
	var createStaffRequest dtos.CreateStaffRequest

	if err := ctx.ShouldBindJSON(&createStaffRequest); err != nil {
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

	CreateStaffResponse, signupError := h.staffService.CreateStaff(&createStaffRequest)

	if signupError != nil {
		ctx.AbortWithStatusJSON(signupError.Code, signupError)

		return
	}

	ctx.JSON(http.StatusCreated, CreateStaffResponse)
}

func (h *StaffHandler) StaffLogin(ctx *gin.Context) {
	var loginStaffRequest dtos.LoginStaffRequest

	if err := ctx.ShouldBindJSON(&loginStaffRequest); err != nil {
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

	LoginStaffResponse, signupError := h.staffService.StaffLogin(&loginStaffRequest)

	if signupError != nil {
		ctx.AbortWithStatusJSON(signupError.Code, signupError)

		return
	}

	ctx.JSON(http.StatusCreated, LoginStaffResponse)
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "min":
		return fmt.Sprintf("Minimum length is %s", fe.Param())
	case "custom_password":
		return "Password must be at least 8 characters long and include uppercase, lowercase, number, and special character"
	default:
		return "Invalid value"
	}
}
