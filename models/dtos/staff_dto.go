package dtos

import (
	"hospital-a/entities"
)

type StaffResponse struct {
	ID           int    `json:"id"`
	UserName     string `json:"username"`
	HospitalName string `json:"hospital_name"`
}

type CreateStaffRequest struct {
	UserName        string `json:"username" binding:"required,min=3,max=50"`
	Password        string `json:"password" binding:"required,min=3,max=50"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
	HospitalName    string `json:"hospital_name" binding:"required"`
}

type LoginStaffRequest struct {
	UserName     string `json:"username" binding:"required,min=3,max=50"`
	Password     string `json:"password" binding:"required,min=3,max=50"`
	HospitalName string `json:"hospital_name" binding:"required,min=3,max=50"`
}

type CreateStaffResponse struct {
	UserName string `json:"username" binding:"required,min=3,max=50"`
	Message  string `json:"message" binding:"required"`
}

type LoginStaffResponse struct {
	Token   string `json:"token" binding:"required"`
	Message string `json:"message" binding:"required"`
}

func (r *StaffResponse) MapStaffResponse(staff *entities.Staff) {
	r.UserName = staff.UserName
	r.HospitalName = staff.HospitalName
	r.ID = staff.ID
}

func (sr *CreateStaffRequest) ToStaff(hashPassword string) *entities.Staff {
	return &entities.Staff{
		UserName:     sr.UserName,
		HospitalName: sr.HospitalName,
		Password:     hashPassword,
	}
}

func (sr *CreateStaffResponse) FromStaff(staff *entities.Staff) *CreateStaffResponse {
	return &CreateStaffResponse{
		UserName: staff.UserName,
		Message:  "User created successfully.",
	}
}
