package handler

import (
	"encoding/json"
	"net/http"

	"gitlab.com/med-group5591241/med/internal/app/handler/dto"
	"gitlab.com/med-group5591241/med/internal/app/service"
)

type MedicationHandler struct {
	medicationService *service.MedicationService
}

func NewMedicationHandler(medicationService *service.MedicationService) *MedicationHandler {
	return &MedicationHandler{medicationService: medicationService}
}

// GetMedications godoc
// @summary     Список медикаментов
// @description Возвращает справочник медикаментов
// @tags Medication
// @success 200 {object} []dto.MedicationResponse
// @router /medications [get]
func (h *MedicationHandler) GetMedications(w http.ResponseWriter, r *http.Request) {
	medications := h.medicationService.GetMedications()
	if len(medications) == 0 {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Медикаментов нет"))
		return
	}

	medicationsResponse := []dto.MedicationResponse{}
	for _, medication := range medications {
		medicationsResponse = append(medicationsResponse, dto.MedicationResponse{
			ID:   medication.ID,
			Name: medication.Name,
			Form: medication.Form,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(medicationsResponse)
	w.Write(b)
}
