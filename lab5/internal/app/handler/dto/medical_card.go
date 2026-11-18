package dto

import "time"

type MedicalCardCreateRequest struct {
	PatientID int `json:"patient_id"`
	DoctorID  int `json:"doctor_id"`
	WardID    int `json:"ward_id"`
}

type MedicalCardUpdateRequest struct {
	DiagnosisID     int    `json:"diagnosis_id"`
	Recommendations string `json:"recommendations"`
}

type MedicalCardResponse struct {
	ID              int                    `json:"id"`
	PatientID       int                    `json:"patient_id"`
	DoctorID        int                    `json:"doctor_id"`
	WardID          int                    `json:"ward_id"`
	DiagnosisID     *int                   `json:"diagnosis_id"`
	OpenedAt        time.Time              `json:"opened_at"`
	ClosedAt        *time.Time             `json:"closed_at"`
	Recommendations *string                `json:"recommendations"`
	Prescriptions   []PrescriptionResponse `json:"prescriptions"`
}
