package model

type Ward struct {
	ID       int
	Number   int
	Beds     int
	FreeBeds int `db:"free_beds"`
}
