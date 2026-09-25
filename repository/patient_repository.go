package repository

import (
	"database/sql"
	"fmt"
	"hospital-a/database"
	"hospital-a/entities"
	"hospital-a/models/dtos"
)

type Patient struct {
	database.BaseSQLRepository[entities.Patient]
}

func NewPatientRepository(db *sql.DB) *Patient {
	return &Patient{
		BaseSQLRepository: database.BaseSQLRepository[entities.Patient]{DB: db},
	}
}

func mapPatient(rows *sql.Row, u *entities.Patient) error {
	return rows.Scan(&u.ID, &u.FirstNameTH, &u.MiddleNameTH, &u.LastNameTH, &u.FirstNameEN, &u.MiddleNameEN, &u.LastNameEN, &u.DateOfBirth, &u.Email, &u.Gender, &u.NationalId, &u.PassportId, &u.PatientHospitalName, &u.PhoneNumber)
}

func mapPatients(rows *sql.Rows, u *entities.Patient) error {
	return rows.Scan(&u.ID, &u.FirstNameTH, &u.MiddleNameTH, &u.LastNameTH, &u.FirstNameEN, &u.MiddleNameEN, &u.LastNameEN, &u.DateOfBirth, &u.Email, &u.Gender, &u.NationalId, &u.PassportId, &u.PatientHospitalName, &u.PhoneNumber)
}

func (r *Patient) FindByNationalIdOrPassportIdAndPatientHn(hospital_name string, uniqueid string) (*entities.Patient, error) {
	return r.SelectSingle(
		mapPatient,
		"SELECT * FROM patient u WHERE (national_id = $1 OR passport_id = $1) AND patient_hn = $2",
		uniqueid, hospital_name,
	)
}

func (r *Patient) FindById(id int) (*entities.Patient, error) {
	return r.SelectSingle(
		mapPatient,
		"SELECT * FROM patient u WHERE u.id = $1",
		id,
	)
}

func (r *Patient) GetAllPatients(hospital_name string, filter *dtos.PatientFilter) ([]*entities.Patient, error) {
	query := `
	SELECT * FROM patient WHERE patient_hn = $1
	`

	var args []any
	args = append(args, hospital_name)
	index := 2

	if filter.NationalId != nil {
		query += fmt.Sprintf(" AND national_id = $%d", index)
		args = append(args, *filter.NationalId)
		index++
	}

	if filter.PassportId != nil {
		query += fmt.Sprintf(" AND passport_id = $%d", index)
		args = append(args, *filter.PassportId)
		index++
	}

	if filter.FirstName != nil {
		query += fmt.Sprintf(" AND (first_name_th ILIKE $%d OR first_name_en ILIKE $%d)", index, index)
		args = append(args, "%"+*filter.FirstName+"%")
		index++
	}

	if filter.MiddleName != nil {
		query += fmt.Sprintf(" AND (middle_name_th ILIKE $%d OR middle_name_en ILIKE $%d)", index, index)
		args = append(args, "%"+*filter.MiddleName+"%")
		index++
	}

	if filter.LastName != nil {
		query += fmt.Sprintf(" AND (last_name_th ILIKE $%d OR last_name_en ILIKE $%d)", index, index)
		args = append(args, "%"+*filter.LastName+"%")
		index++
	}

	if filter.Email != nil {
		query += fmt.Sprintf(" AND email ILIKE $%d", index)
		args = append(args, "%"+*filter.Email+"%")
		index++
	}

	if filter.PhoneNumber != nil {
		query += fmt.Sprintf(" AND phone_number ILIKE $%d", index)
		args = append(args, "%"+*filter.PhoneNumber+"%")
		index++
	}

	if filter.DateOfBirth != nil {
		query += fmt.Sprintf(" AND date_of_birth = $%d", index)
		args = append(args, *filter.DateOfBirth)
		index++
	}

	return r.SelectMultiple(
		mapPatients,
		query,
		args...,
	)
}

func (r *Patient) CreatePatient(patient *entities.Patient) error {
	_, err := r.Insert(
		`INSERT INTO patient (
        first_name_th,
        middle_name_th,
        last_name_th,
        first_name_en,
        middle_name_en,
        last_name_en,
        date_of_birth,
        email,
        gender,
        national_id,
        passport_id,
        patient_hospital_name,
        phone_number
    )
    VALUES (
        $1, $2, $3, $4, $5, $6, $7,
        $8, $9, $10, $11, $12, $13
    )`,
		patient.FirstNameTH,
		patient.MiddleNameTH,
		patient.LastNameTH,
		patient.FirstNameEN,
		patient.MiddleNameEN,
		patient.LastNameEN,
		patient.DateOfBirth,
		patient.Email,
		patient.Gender,
		patient.NationalId,
		patient.PassportId,
		patient.PatientHospitalName,
		patient.PhoneNumber,
	)

	return err
}

func (r *Patient) Update(patient *entities.Patient) error {
	_, err := r.ExecuteQuery(
		`UPDATE patient
        SET
            first_name_th = $1,
            middle_name_th = $2,
            last_name_th = $3,
            first_name_en = $4,
            middle_name_en = $5,
            last_name_en = $6,
            date_of_birth = $7,
            email = $8,
            gender = $9,
            national_id = $10,
            passport_id = $11,
            patient_hospital_name = $12,
            phone_number = $13
        WHERE id = $14`,
		patient.FirstNameTH,
		patient.MiddleNameTH,
		patient.LastNameTH,
		patient.FirstNameEN,
		patient.MiddleNameEN,
		patient.LastNameEN,
		patient.DateOfBirth,
		patient.Email,
		patient.Gender,
		patient.NationalId,
		patient.PassportId,
		patient.PatientHospitalName,
		patient.PhoneNumber,
		patient.ID,
	)

	return err
}

func (r *Patient) DeleteUser(id int) error {
	_, err := r.ExecuteQuery("DELETE FROM patient WHERE id = $1", id)

	return err
}
