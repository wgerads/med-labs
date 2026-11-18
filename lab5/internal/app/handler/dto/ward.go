package dto

type WardResponse struct {
	ID       int `json:"id"`
	Number   int `json:"number"`
	Beds     int `json:"beds"`
	FreeBeds int `json:"free_beds"`
}
