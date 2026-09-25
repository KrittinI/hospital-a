package dtos

import (
	"hospital-a/entities"
	"time"
)

type PatientFilter struct {
	NationalId  *string    `form:"nationalId"`
	PassportId  *string    `form:"passportId"`
	FirstName   *string    `form:"firstName"`
	MiddleName  *string    `form:"middleName"`
	LastName    *string    `form:"lastName"`
	Email       *string    `form:"email"`
	PhoneNumber *string    `form:"phoneNumber"`
	DateOfBirth *time.Time `form:"dateOfBirth"`
}

type PatientResponse struct {
	FirstNameTH         string `json:"first_name_th"`
	MiddleNameTH        string `json:"middle_name_th"`
	LastNameTH          string `json:"last_name_th"`
	FirstNameEN         string `json:"first_name_en"`
	MiddleNameEN        string `json:"middle_name_en"`
	LastNameEN          string `json:"last_name_en"`
	DateOfBirth         string `json:"date_of_birth"`
	PatientHospitalName string `json:"patient_hn"`
	NationalId          string `json:"national_id"`
	PassportId          string `json:"passport_id"`
	PhoneNumber         string `json:"phone_number"`
	Email               string `json:"email"`
	Gender              string `json:"gender"`
}

type GetAllPatientsResponse struct {
	Patients []*PatientResponse `json:"patients"`
}

type CreatePatientRequest struct {
	FirstNameTH         string `json:"first_name_th" binding:"required,max=50"`
	MiddleNameTH        string `json:"middle_name_th"`
	LastNameTH          string `json:"last_name_th" binding:"required,max=50"`
	FirstNameEN         string `json:"first_name_en" binding:"required,max=50"`
	MiddleNameEN        string `json:"middle_name_en"`
	LastNameEN          string `json:"last_name_en" binding:"required,max=50"`
	Email               string `json:"email" binding:"required,email,max=254"`
	PhoneNumber         string `json:"phone_number" binding:"required"`
	PatientHospitalName string `json:"patient_hn" binding:"required"`
	DateOfBirth         string `json:"date_of_birth" binding:"required"`
	NationalId          string `json:"national_id" binding:"required_without=PassportId"`
	PassportId          string `json:"passport_id" binding:"required_without=NationalId"`
	Gender              string `json:"gender"`
}

type UpdatePatientRequest struct {
	FirstNameTH         string `json:"first_name_th" binding:"required,max=50"`
	MiddleNameTH        string `json:"middle_name_th"`
	LastNameTH          string `json:"last_name_th" binding:"required,max=50"`
	FirstNameEN         string `json:"first_name_en" binding:"required,max=50"`
	MiddleNameEN        string `json:"middle_name_en"`
	LastNameEN          string `json:"last_name_en" binding:"required,max=50"`
	Email               string `json:"email" binding:"required,email,max=254"`
	PhoneNumber         string `json:"phone_number" binding:"required"`
	PatientHospitalName string `json:"patient_hn" binding:"required"`
	DateOfBirth         string `json:"date_of_birth" binding:"required"`
	NationalId          string `json:"national_id" binding:"required_without=PassportId"`
	PassportId          string `json:"passport_id" binding:"required_without=NationalId"`
	Gender              string `json:"gender"`
}

type CreatePatientResponse struct {
	FirstNameTH         string `json:"first_name_th" binding:"required,max=50"`
	MiddleNameTH        string `json:"middle_name_th"`
	LastNameTH          string `json:"last_name_th" binding:"required,max=50"`
	FirstNameEN         string `json:"first_name_en" binding:"required,max=50"`
	MiddleNameEN        string `json:"middle_name_en"`
	LastNameEN          string `json:"last_name_en" binding:"required,max=50"`
	Email               string `json:"email" binding:"required,email,max=254"`
	PhoneNumber         string `json:"phone_number" binding:"required"`
	PatientHospitalName string `json:"patient_hn" binding:"required"`
	DateOfBirth         string `json:"date_of_birth" binding:"required"`
	NationalId          string `json:"national_id" binding:"required_without=PassportId"`
	PassportId          string `json:"passport_id" binding:"required_without=NationalId"`
	Gender              string `json:"gender"`
	Message             string `json:"message" binding:"required"`
}

func (r *GetAllPatientsResponse) MapPatientsResponse(patients []*entities.Patient) {
	for _, patients := range patients {
		patient := &PatientResponse{
			FirstNameTH:         patients.FirstNameTH,
			MiddleNameTH:        patients.MiddleNameTH,
			LastNameTH:          patients.LastNameTH,
			FirstNameEN:         patients.FirstNameEN,
			MiddleNameEN:        patients.MiddleNameEN,
			LastNameEN:          patients.LastNameEN,
			Email:               patients.Email,
			PhoneNumber:         patients.PhoneNumber,
			DateOfBirth:         patients.DateOfBirth,
			PatientHospitalName: patients.PatientHospitalName,
			NationalId:          patients.NationalId,
			PassportId:          patients.PassportId,
			Gender:              patients.Gender,
		}
		r.Patients = append(r.Patients, patient)
	}
}

func (r *PatientResponse) MapPatientResponse(patient *entities.Patient) {
	r.DateOfBirth = patient.DateOfBirth
	r.Email = patient.Email
	r.FirstNameEN = patient.FirstNameEN
	r.MiddleNameEN = patient.MiddleNameEN
	r.LastNameEN = patient.LastNameEN
	r.FirstNameTH = patient.FirstNameTH
	r.MiddleNameTH = patient.MiddleNameTH
	r.LastNameTH = patient.LastNameTH
	r.Gender = patient.Gender
	r.NationalId = patient.NationalId
	r.PassportId = patient.PassportId
	r.PatientHospitalName = patient.PatientHospitalName
	r.PhoneNumber = patient.PhoneNumber
}

func (pr *CreatePatientRequest) ToPatient() *entities.Patient {
	return &entities.Patient{
		FirstNameTH:         pr.FirstNameTH,
		MiddleNameTH:        pr.MiddleNameTH,
		LastNameTH:          pr.LastNameTH,
		FirstNameEN:         pr.FirstNameEN,
		MiddleNameEN:        pr.MiddleNameEN,
		LastNameEN:          pr.LastNameEN,
		Email:               pr.Email,
		PhoneNumber:         pr.PhoneNumber,
		DateOfBirth:         pr.DateOfBirth,
		PatientHospitalName: pr.PatientHospitalName,
		NationalId:          pr.NationalId,
		PassportId:          pr.PassportId,
		Gender:              pr.Gender,
	}
}

func (pr *UpdatePatientRequest) ToPatient() *entities.Patient {
	return &entities.Patient{
		FirstNameTH:         pr.FirstNameTH,
		MiddleNameTH:        pr.MiddleNameTH,
		LastNameTH:          pr.LastNameTH,
		FirstNameEN:         pr.FirstNameEN,
		MiddleNameEN:        pr.MiddleNameEN,
		LastNameEN:          pr.LastNameEN,
		Email:               pr.Email,
		PhoneNumber:         pr.PhoneNumber,
		DateOfBirth:         pr.DateOfBirth,
		PatientHospitalName: pr.PatientHospitalName,
		NationalId:          pr.NationalId,
		PassportId:          pr.PassportId,
		Gender:              pr.Gender,
	}
}

func (pr *CreatePatientResponse) FromPatient(patient *entities.Patient) *CreatePatientResponse {
	return &CreatePatientResponse{
		FirstNameTH:         pr.FirstNameTH,
		MiddleNameTH:        pr.MiddleNameTH,
		LastNameTH:          pr.LastNameTH,
		FirstNameEN:         pr.FirstNameEN,
		MiddleNameEN:        pr.MiddleNameEN,
		LastNameEN:          pr.LastNameEN,
		Email:               pr.Email,
		PhoneNumber:         pr.PhoneNumber,
		DateOfBirth:         pr.DateOfBirth,
		PatientHospitalName: pr.PatientHospitalName,
		NationalId:          pr.NationalId,
		PassportId:          pr.PassportId,
		Gender:              pr.Gender,
		Message:             "Patient created successfully.",
	}
}
