package model

type Doctor struct {
	ID        int
	FullName  string `db:"full_name"`
	Specialty string
}
