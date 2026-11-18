package model

type Diagnosis struct {
	ID      int
	Name    string
	MKBCode string `db:"mkb_code"`
}
