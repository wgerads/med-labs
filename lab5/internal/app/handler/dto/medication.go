package dto

type MedicationResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Form string `json:"form"`
}
