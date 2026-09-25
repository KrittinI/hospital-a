package entities

type Staff struct {
	ID           int    `json:"id"`
	UserName     string `json:"username"`
	Password     string `json:"password"`
	HospitalName string `json:"hospital_name"`
}
