package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"gitlab.com/med-group5591241/med/internal/app/model"
)

type PatientRepository struct {
	db *sqlx.DB
}

func NewPatientRepository(db *sqlx.DB) *PatientRepository {
	return &PatientRepository{db: db}
}

func (r *PatientRepository) GetPatients() []model.Patient {
	patients := []model.Patient{}
	err := r.db.Select(&patients, "SELECT id, full_name, birth_date, phone, policy FROM patient ORDER BY id")
	if err != nil {
		fmt.Println("get patients err", err)
		return patients
	}

	return patients
}

func (r *PatientRepository) GetPatient(id int) *model.Patient {
	patient := model.Patient{}
	err := r.db.Get(&patient, "SELECT id, full_name, birth_date, phone, policy FROM patient WHERE id=$1", id)
	if err != nil {
		fmt.Println("get patient err", err)
		return nil
	}

	err = r.db.Select(&patient.Cards, "SELECT id, patient_id, doctor_id, ward_id, diagnosis_id, opened_at, closed_at, recommendations FROM medical_card WHERE patient_id=$1 ORDER BY id", id)
	if err != nil {
		fmt.Println("get patient cards err", err)
	}

	return &patient
}

func (r *PatientRepository) CreatePatient(patient model.Patient) error {
	_, err := r.db.NamedExec("INSERT INTO patient (full_name, birth_date, phone, policy) VALUES (:full_name, :birth_date, :phone, :policy)", patient)
	if err != nil {
		return err
	}

	return nil
}
