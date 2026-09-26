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

func mapPatient(rows *sql.Row, p *entities.Patient) error {
	return rows.Scan(&p.ID, &p.FirstNameTH, &p.MiddleNameTH, &p.LastNameTH, &p.FirstNameEN, &p.MiddleNameEN, &p.LastNameEN, &p.DateOfBirth, &p.Email, &p.Gender, &p.NationalId, &p.PassportId, &p.PatientHospitalName, &p.PhoneNumber)
}

func mapPatients(rows *sql.Rows, p *entities.Patient) error {
	return rows.Scan(&p.ID, &p.FirstNameTH, &p.MiddleNameTH, &p.LastNameTH, &p.FirstNameEN, &p.MiddleNameEN, &p.LastNameEN, &p.DateOfBirth, &p.Email, &p.Gender, &p.NationalId, &p.PassportId, &p.PatientHospitalName, &p.PhoneNumber)
}

func (r *Patient) FindByNationalIdOrPassportIdAndPatientHn(hospital_name string, uniqueid string) (*entities.Patient, error) {
	return r.SelectSingle(
		mapPatient,
		`SELECT 
		p.id,
	    p.first_name_th,
    	p.middle_name_th,
    	p.last_name_th,
    	p.first_name_en,
    	p.middle_name_en,
    	p.last_name_en,
    	p.date_of_birth,
    	p.email,
    	p.gender,
    	p.national_id,
    	p.passport_id,
    	p.patient_hn,
    	p.phone_number 
		FROM patient p WHERE (p.national_id = $1 OR p.passport_id = $1) AND p.patient_hn = $2`,
		uniqueid, hospital_name,
	)
}

func (r *Patient) FindById(id int) (*entities.Patient, error) {
	return r.SelectSingle(
		mapPatient,
		"SELECT * FROM patient p WHERE p.id = $1",
		id,
	)
}

func (r *Patient) GetAllPatients(hospital_name string, filter *dtos.PatientFilter) ([]*entities.Patient, error) {
	query := `
	SELECT 
	p.id,
	    p.first_name_th,
    	p.middle_name_th,
    	p.last_name_th,
    	p.first_name_en,
    	p.middle_name_en,
    	p.last_name_en,
    	p.date_of_birth,
    	p.email,
    	p.gender,
    	p.national_id,
    	p.passport_id,
    	p.patient_hn,
    	p.phone_number 
		FROM patient p WHERE p.patient_hn = $1
	`

	var args []any
	args = append(args, hospital_name)
	index := 2

	if filter.NationalId != nil {
		query += fmt.Sprintf(" AND p.national_id = $%d", index)
		args = append(args, *filter.NationalId)
		index++
	}

	if filter.PassportId != nil {
		query += fmt.Sprintf(" AND p.passport_id = $%d", index)
		args = append(args, *filter.PassportId)
		index++
	}

	if filter.FirstName != nil {
		query += fmt.Sprintf(" AND (p.first_name_th ILIKE $%d OR p.first_name_en ILIKE $%d)", index, index)
		args = append(args, "%"+*filter.FirstName+"%")
		index++
	}

	if filter.MiddleName != nil {
		query += fmt.Sprintf(" AND (p.middle_name_th ILIKE $%d OR p.middle_name_en ILIKE $%d)", index, index)
		args = append(args, "%"+*filter.MiddleName+"%")
		index++
	}

	if filter.LastName != nil {
		query += fmt.Sprintf(" AND (p.last_name_th ILIKE $%d OR p.last_name_en ILIKE $%d)", index, index)
		args = append(args, "%"+*filter.LastName+"%")
		index++
	}

	if filter.Email != nil {
		query += fmt.Sprintf(" AND p.email ILIKE $%d", index)
		args = append(args, "%"+*filter.Email+"%")
		index++
	}

	if filter.PhoneNumber != nil {
		query += fmt.Sprintf(" AND p.phone_number ILIKE $%d", index)
		args = append(args, "%"+*filter.PhoneNumber+"%")
		index++
	}

	if filter.DateOfBirth != nil {
		query += fmt.Sprintf(" AND p.date_of_birth = $%d", index)
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
        patient_hn,
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
            patient_hn = $12,
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
