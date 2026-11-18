package model

import "time"

type MedicalCard struct {
	ID              int
	PatientID       int        `db:"patient_id"`
	DoctorID        int        `db:"doctor_id"`
	WardID          int        `db:"ward_id"`
	DiagnosisID     *int       `db:"diagnosis_id"`
	OpenedAt        time.Time  `db:"opened_at"`
	ClosedAt        *time.Time `db:"closed_at"`
	Recommendations *string
	Prescriptions   []Prescription
}
