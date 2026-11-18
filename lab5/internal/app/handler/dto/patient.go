package dto

import "time"

type PatientCreateRequest struct {
	FullName  string  `json:"full_name"`
	BirthDate string  `json:"birth_date"`
	Phone     *string `json:"phone"`
	Policy    string  `json:"policy"`
}

type PatientResponse struct {
	ID        int       `json:"id"`
	FullName  string    `json:"full_name"`
	BirthDate time.Time `json:"birth_date"`
	Phone     *string   `json:"phone"`
	Policy    string    `json:"policy"`
}
