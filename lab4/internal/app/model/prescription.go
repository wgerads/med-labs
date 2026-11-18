package model

type Prescription struct {
	ID            int
	MedicalCardID int `db:"medical_card_id"`
	MedicationID  int `db:"medication_id"`
	Frequency     string
	Duration      string
}
