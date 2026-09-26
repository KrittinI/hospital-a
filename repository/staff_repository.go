package repository

import (
	"database/sql"
	"hospital-a/database"
	"hospital-a/entities"
)

type Staff struct {
	database.BaseSQLRepository[entities.Staff]
}

func NewStaffRepository(db *sql.DB) *Staff {
	return &Staff{
		BaseSQLRepository: database.BaseSQLRepository[entities.Staff]{DB: db},
	}
}

func mapStaff(rows *sql.Row, s *entities.Staff) error {
	return rows.Scan(&s.ID, &s.UserName, &s.HospitalName, &s.Password)
}

func (r *Staff) FindStaffByUsername(username string) (*entities.Staff, error) {
	return r.SelectSingle(
		mapStaff,
		"SELECT s.id, s.username, s.hospital_name, s.password FROM staff s WHERE s.username = $1",
		username,
	)
}

func (r *Staff) FindStaffById(id int) (*entities.Staff, error) {
	return r.SelectSingle(
		mapStaff,
		"SELECT s.id, s.username, s.hospital_name FROM staff s WHERE s.id = $1",
		id,
	)
}

func (r *Staff) CreateStaff(staff *entities.Staff) error {
	_, err := r.Insert(
		"INSERT INTO staff (username, password, hospital_name) VALUES ($1, $2, $3)",
		staff.UserName, staff.Password, staff.HospitalName,
	)

	return err
}

func (r *Staff) UpdateStaff(staff *entities.Staff) error {
	_, err := r.ExecuteQuery(
		"UPDATE staff SET username = $1, password = $2, hospital_name = $3 WHERE id = $4",
		staff.UserName, staff.Password, staff.HospitalName, staff.ID,
	)

	return err
}

func (r *Staff) DeleteStaff(id int) error {
	_, err := r.ExecuteQuery("DELETE FROM staff WHERE id = $1", id)

	return err
}
