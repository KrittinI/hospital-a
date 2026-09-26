package service

import (
	"database/sql"
	"errors"
	"hospital-a/models"
	"hospital-a/models/dtos"
	"hospital-a/repository"
	"hospital-a/share"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type StaffService struct {
	staffRepo *repository.Staff
}

func NewStaffService(staffRepo *repository.Staff) *StaffService {
	return &StaffService{staffRepo: staffRepo}
}

func (us *StaffService) GetStaff(staffId int) (*dtos.StaffResponse, *models.ErrorResponse) {
	response := &dtos.StaffResponse{}

	staff, err := us.staffRepo.FindStaffById(staffId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "Staff Not Found",
			}
		}

		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
	}

	response.MapStaffResponse(staff)

	return response, nil
}

func (us *StaffService) DeleteStaff(staffId int) *models.ErrorResponse {
	staff, err := us.staffRepo.FindStaffById(staffId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "Staff not found",
			}
		}
		return &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
	}

	err = us.staffRepo.DeleteStaff(staff.ID)

	if err != nil {
		return &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
	}

	return nil
}

func (us *StaffService) CreateStaff(createStaffRequest *dtos.CreateStaffRequest) (*dtos.CreateStaffResponse, *models.ErrorResponse) {
	staffResponse := &dtos.CreateStaffResponse{}

	errEmail := us.checkIfUsernameExists(createStaffRequest.UserName)

	if errEmail != nil {
		return nil, errEmail
	}

	hash, errHash := HashPassword(createStaffRequest.Password)

	if errHash != nil {
		return nil, &models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Password hash failed.",
		}
	}

	staff := createStaffRequest.ToStaff(hash)

	err := us.staffRepo.CreateStaff(staff)

	if err != nil {
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create staff",
		}
	}

	return staffResponse.FromStaff(staff), nil
}

func (us *StaffService) StaffLogin(loginStaffRequest *dtos.LoginStaffRequest) (*dtos.LoginStaffResponse, *models.ErrorResponse) {
	staffResponse := &dtos.LoginStaffResponse{}

	staff, err := us.staffRepo.FindStaffByUsername(loginStaffRequest.UserName)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
	}

	if staff == nil {
		return nil, &models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Username or Password incorrect",
		}
	}

	if staff.HospitalName != loginStaffRequest.HospitalName {
		return nil, &models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Username or Password incorrect",
		}
	}

	passCorrect := CheckPassword(loginStaffRequest.Password, staff.Password)

	if !passCorrect {
		return nil, &models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Username or Password incorrect",
		}
	}

	// jwt
	token, err := share.GenerateToken(staff)

	if err != nil {
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Token generate fail.",
		}
	}

	staffResponse.Message = "Login success"
	staffResponse.Token = token

	return staffResponse, nil
}

// func (us *Staff) UpdateStaff(userID int, updateStaffRequest *dtos.UpdateStaffRequest) *models.ErrorResponse {
// 	existingStaff, err := us.staffRepo.FindById(userID)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return &models.ErrorResponse{
// 				Code:    http.StatusNotFound,
// 				Message: "Staff not found",
// 			}
// 		}
// 		return &models.ErrorResponse{
// 			Code:    http.StatusInternalServerError,
// 			Message: "Internal Server Error",
// 		}
// 	}

// 	if updateStaffRequest.Email != existingStaff.Email {
// 		errEmail := us.checkIfUsernameExists(updateStaffRequest.Email)
// 		if errEmail != nil {
// 			return errEmail
// 		}
// 	}

// 	existingStaff = updateStaffRequest.ToStaff()
// 	existingStaff.ID = userID

// 	err = us.staffRepo.Update(existingStaff)

// 	if err != nil {
// 		return &models.ErrorResponse{
// 			Code:    http.StatusInternalServerError,
// 			Message: "Failed to update user",
// 		}
// 	}

// 	return nil
// }

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func (us *StaffService) checkIfUsernameExists(username string) *models.ErrorResponse {
	staff, err := us.staffRepo.FindStaffByUsername(username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
	}

	if staff != nil {
		return &models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Username already in use",
		}
	}

	return nil
}
