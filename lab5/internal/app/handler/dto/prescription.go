package dto

type PrescriptionCreateRequest struct {
	MedicalCardID int    `json:"medical_card_id"`
	MedicationID  int    `json:"medication_id"`
	Frequency     string `json:"frequency"`
	Duration      string `json:"duration"`
}

type PrescriptionResponse struct {
	ID           int    `json:"id"`
	MedicationID int    `json:"medication_id"`
	Frequency    string `json:"frequency"`
	Duration     string `json:"duration"`
}
