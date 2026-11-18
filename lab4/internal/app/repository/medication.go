package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"gitlab.com/med-group5591241/med/internal/app/model"
)

type MedicationRepository struct {
	db *sqlx.DB
}

func NewMedicationRepository(db *sqlx.DB) *MedicationRepository {
	return &MedicationRepository{db: db}
}

func (r *MedicationRepository) GetMedications() []model.Medication {
	medications := []model.Medication{}
	err := r.db.Select(&medications, "SELECT id, name, form FROM medication ORDER BY id")
	if err != nil {
		fmt.Println("get medications err", err)
		return medications
	}

	return medications
}
