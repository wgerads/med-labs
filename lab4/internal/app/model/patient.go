package model

import "time"

type Patient struct {
	ID        int
	FullName  string    `db:"full_name"`
	BirthDate time.Time `db:"birth_date"`
	Phone     *string
	Policy    string
	Cards     []MedicalCard
}
