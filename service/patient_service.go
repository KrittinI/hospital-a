package service

import (
	"database/sql"
	"errors"
	"fmt"
	"hospital-a/models"
	"hospital-a/models/dtos"
	"hospital-a/repository"
	"net/http"
)

type PatientService struct {
	patientRepo *repository.Patient
}

func NewPatientService(patientRepo *repository.Patient) *PatientService {
	return &PatientService{patientRepo: patientRepo}
}

func (p *PatientService) GetAllPatients(hospital_name string, filter *dtos.PatientFilter) (*dtos.GetAllPatientsResponse, *models.ErrorResponse) {
	response := &dtos.GetAllPatientsResponse{}

	queriedUsers, err := p.patientRepo.GetAllPatients(hospital_name, filter)
	if err != nil {
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
	}

	response.MapPatientsResponse(queriedUsers)

	return response, nil
}

func (p *PatientService) GetPatientById(hospital_name string, uniqueid string) (*dtos.PatientResponse, *models.ErrorResponse) {
	response := &dtos.PatientResponse{}

	patient, err := p.patientRepo.FindByNationalIdOrPassportIdAndPatientHn(hospital_name, uniqueid)

	fmt.Println(patient, err)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "Patient Not Found",
			}
		}
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
	}

	response.MapPatientResponse(patient)

	return response, nil
}

func (p *PatientService) DeletePatient(patientId int) *models.ErrorResponse {
	patient, err := p.patientRepo.FindById(patientId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "User not found",
			}
		}
		return &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
	}
	err = p.patientRepo.DeleteUser(patient.ID)
	if err != nil {
		return &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
	}

	return nil
}

func (p *PatientService) CreatePatient(createPatientRequest *dtos.CreatePatientRequest) (*dtos.CreatePatientResponse, *models.ErrorResponse) {
	userResponse := &dtos.CreatePatientResponse{}

	errUnique := p.checkIfPatientExists(createPatientRequest)
	if errUnique != nil {
		return nil, errUnique
	}

	patient := createPatientRequest.ToPatient()

	err := p.patientRepo.CreatePatient(patient)
	if err != nil {
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create patient",
		}
	}

	return userResponse.FromPatient(patient), nil
}

// func (p *Patient) UpdateUser(patientId int, UpdatePatientRequest *dtos.UpdatePatientRequest) *models.ErrorResponse {

// 	existingUser, err := p.patientRepo.FindById(patientId)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return &models.ErrorResponse{
// 				Code:    http.StatusNotFound,
// 				Message: "User not found",
// 			}
// 		}
// 		return &models.ErrorResponse{
// 			Code:    http.StatusInternalServerError,
// 			Message: "Internal Server Error",
// 		}
// 	}

// 	if UpdatePatientRequest.Email != existingUser.Email {
// 		errEmail := p.checkIfEmailExists(UpdatePatientRequest.Email)
// 		if errEmail != nil {
// 			return errEmail
// 		}
// 	}

// 	existingUser = UpdatePatientRequest.ToUser()
// 	existingUser.ID = userID

// 	err = p.patientRepo.Update(existingUser)

// 	if err != nil {
// 		return &models.ErrorResponse{
// 			Code:    http.StatusInternalServerError,
// 			Message: "Failed to update user",
// 		}
// 	}

// 	return nil
// }

func (p *PatientService) checkIfPatientExists(createPatientRequest *dtos.CreatePatientRequest) *models.ErrorResponse {
	patientNationalId, err := p.patientRepo.FindByNationalIdOrPassportIdAndPatientHn(createPatientRequest.PatientHospitalName, createPatientRequest.NationalId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
	}

	if patientNationalId != nil {
		return &models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Paitent already exists",
		}
	}

	patientPassportId, err := p.patientRepo.FindByNationalIdOrPassportIdAndPatientHn(createPatientRequest.PatientHospitalName, createPatientRequest.PassportId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
	}

	if patientPassportId != nil {
		return &models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Paitent already exists",
		}
	}

	return nil
}
