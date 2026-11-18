package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"gitlab.com/med-group5591241/med/internal/app/model"
)

type WardRepository struct {
	db *sqlx.DB
}

func NewWardRepository(db *sqlx.DB) *WardRepository {
	return &WardRepository{db: db}
}

const getWardsQuery = `
SELECT ward.id, ward.number, ward.beds, ward.beds - COUNT(medical_card.id) AS free_beds
FROM ward LEFT JOIN medical_card ON medical_card.ward_id = ward.id AND medical_card.closed_at IS NULL
GROUP BY ward.id
ORDER BY ward.number
`

func (r *WardRepository) GetWards() []model.Ward {
	wards := []model.Ward{}
	err := r.db.Select(&wards, getWardsQuery)
	if err != nil {
		fmt.Println("get wards err", err)
		return wards
	}

	return wards
}
