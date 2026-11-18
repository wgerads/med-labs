package handler

import (
	"encoding/json"
	"net/http"

	"gitlab.com/med-group5591241/med/internal/app/handler/dto"
	"gitlab.com/med-group5591241/med/internal/app/service"
)

type WardHandler struct {
	wardService *service.WardService
}

func NewWardHandler(wardService *service.WardService) *WardHandler {
	return &WardHandler{wardService: wardService}
}

// GetWards godoc
// @summary     Список палат
// @description Возвращает палаты с количеством свободных коек
// @tags Ward
// @success 200 {object} []dto.WardResponse
// @router /wards [get]
func (h *WardHandler) GetWards(w http.ResponseWriter, r *http.Request) {
	wards := h.wardService.GetWards()
	if len(wards) == 0 {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Палат нет"))
		return
	}

	wardsResponse := []dto.WardResponse{}
	for _, ward := range wards {
		wardsResponse = append(wardsResponse, dto.WardResponse{
			ID:       ward.ID,
			Number:   ward.Number,
			Beds:     ward.Beds,
			FreeBeds: ward.FreeBeds,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(wardsResponse)
	w.Write(b)
}
