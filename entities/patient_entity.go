package entities

import "database/sql"

type Patient struct {
	ID                  int            `json:"id"`
	FirstNameTH         string         `json:"first_name_th"`
	MiddleNameTH        sql.NullString `json:"middle_name_th"`
	LastNameTH          string         `json:"last_name_th"`
	FirstNameEN         string         `json:"first_name_en"`
	MiddleNameEN        sql.NullString `json:"middle_name_en"`
	LastNameEN          string         `json:"last_name_en"`
	DateOfBirth         string         `json:"date_of_birth"`
	PatientHospitalName string         `json:"patient_hn"`
	NationalId          sql.NullString `json:"national_id"`
	PassportId          sql.NullString `json:"passport_id"`
	PhoneNumber         string         `json:"phone_number"`
	Email               string         `json:"email"`
	Gender              string         `json:"gender"`
}
