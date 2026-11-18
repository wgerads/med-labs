package repository

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"gitlab.com/med-group5591241/med/internal/app/model"
)

type PrescriptionRepository struct {
	db *sqlx.DB
}

func NewPrescriptionRepository(db *sqlx.DB) *PrescriptionRepository {
	return &PrescriptionRepository{db: db}
}

func (r *PrescriptionRepository) CreatePrescription(prescription model.Prescription) error {
	_, err := r.db.NamedExec("INSERT INTO prescription (medical_card_id, medication_id, frequency, duration) VALUES (:medical_card_id, :medication_id, :frequency, :duration)", prescription)
	if err != nil {
		return err
	}

	return nil
}

func (r *PrescriptionRepository) DeletePrescription(id int) error {
	result, err := r.db.Exec("DELETE FROM prescription WHERE id=$1", id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("назначение не найдено")
	}

	return nil
}
