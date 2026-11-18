package repository

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"gitlab.com/med-group5591241/med/internal/app/model"
)

type MedicalCardRepository struct {
	db *sqlx.DB
}

func NewMedicalCardRepository(db *sqlx.DB) *MedicalCardRepository {
	return &MedicalCardRepository{db: db}
}

const getFreeBedsQuery = `
SELECT ward.beds - COUNT(medical_card.id)
FROM ward LEFT JOIN medical_card ON medical_card.ward_id = ward.id AND medical_card.closed_at IS NULL
WHERE ward.id=$1
GROUP BY ward.id
`

func (r *MedicalCardRepository) CreateMedicalCard(card model.MedicalCard) error {
	freeBeds := 0
	err := r.db.Get(&freeBeds, getFreeBedsQuery, card.WardID)
	if err != nil {
		return err
	}
	if freeBeds <= 0 {
		return errors.New("в палате нет свободных коек")
	}

	_, err = r.db.NamedExec("INSERT INTO medical_card (patient_id, doctor_id, ward_id) VALUES (:patient_id, :doctor_id, :ward_id)", card)
	if err != nil {
		return err
	}

	return nil
}

func (r *MedicalCardRepository) GetMedicalCard(id int) *model.MedicalCard {
	card := model.MedicalCard{}
	err := r.db.Get(&card, "SELECT id, patient_id, doctor_id, ward_id, diagnosis_id, opened_at, closed_at, recommendations FROM medical_card WHERE id=$1", id)
	if err != nil {
		fmt.Println("get medical card err", err)
		return nil
	}

	err = r.db.Select(&card.Prescriptions, "SELECT id, medical_card_id, medication_id, frequency, duration FROM prescription WHERE medical_card_id=$1 ORDER BY id", id)
	if err != nil {
		fmt.Println("get prescriptions err", err)
	}

	return &card
}

func (r *MedicalCardRepository) UpdateMedicalCard(id int, diagnosisID int, recommendations string) error {
	result, err := r.db.Exec("UPDATE medical_card SET diagnosis_id=$1, recommendations=$2, closed_at=CURRENT_DATE WHERE id=$3", diagnosisID, recommendations, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("медицинская карта не найдена")
	}

	return nil
}
