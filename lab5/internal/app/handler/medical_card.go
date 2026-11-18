package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gitlab.com/med-group5591241/med/internal/app/handler/dto"
	"gitlab.com/med-group5591241/med/internal/app/model"
	"gitlab.com/med-group5591241/med/internal/app/service"
)

type MedicalCardHandler struct {
	medicalCardService *service.MedicalCardService
}

func NewMedicalCardHandler(medicalCardService *service.MedicalCardService) *MedicalCardHandler {
	return &MedicalCardHandler{medicalCardService: medicalCardService}
}

// CreateMedicalCard godoc
// @summary     Открыть медицинскую карту
// @description Госпитализация: открывает карту пациента в палате, если в ней есть свободная койка
// @tags MedicalCard
// @success 200
// @router /medical-cards [post]
// @param request body dto.MedicalCardCreateRequest true "request"
func (h *MedicalCardHandler) CreateMedicalCard(w http.ResponseWriter, r *http.Request) {
	var request dto.MedicalCardCreateRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(400)
		w.Write([]byte("Неверный JSON для открытия медицинской карты"))
		return
	}

	card := model.MedicalCard{
		PatientID: request.PatientID,
		DoctorID:  request.DoctorID,
		WardID:    request.WardID,
	}

	err = h.medicalCardService.CreateMedicalCard(card)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(500)
		w.Write([]byte("Не удалось открыть медицинскую карту: " + err.Error()))
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)
	w.Write([]byte("Медицинская карта открыта"))
}

// GetMedicalCard godoc
// @summary     Получить медицинскую карту
// @description Возвращает карту вместе с назначениями
// @tags MedicalCard
// @success 200 {object} dto.MedicalCardResponse
// @router /medical-cards/{id} [get]
// @param id path int true "id"
func (h *MedicalCardHandler) GetMedicalCard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cardID, _ := strconv.Atoi(vars["id"])

	card := h.medicalCardService.GetMedicalCard(cardID)
	if card == nil {
		w.Write([]byte("Медицинская карта не найдена"))
		return
	}

	prescriptionsResponse := []dto.PrescriptionResponse{}
	for _, prescription := range card.Prescriptions {
		prescriptionsResponse = append(prescriptionsResponse, dto.PrescriptionResponse{
			ID:           prescription.ID,
			MedicationID: prescription.MedicationID,
			Frequency:    prescription.Frequency,
			Duration:     prescription.Duration,
		})
	}

	cardResponse := dto.MedicalCardResponse{
		ID:              card.ID,
		PatientID:       card.PatientID,
		DoctorID:        card.DoctorID,
		WardID:          card.WardID,
		DiagnosisID:     card.DiagnosisID,
		OpenedAt:        card.OpenedAt,
		ClosedAt:        card.ClosedAt,
		Recommendations: card.Recommendations,
		Prescriptions:   prescriptionsResponse,
	}

	w.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(&cardResponse)
	w.Write(b)
}

// UpdateMedicalCard godoc
// @summary     Закрыть медицинскую карту
// @description Выписка: указывает диагноз и рекомендации, закрывает карту
// @tags MedicalCard
// @success 200
// @router /medical-cards/{id} [patch]
// @param id path int true "id"
// @param request body dto.MedicalCardUpdateRequest true "request"
func (h *MedicalCardHandler) UpdateMedicalCard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cardID, _ := strconv.Atoi(vars["id"])

	var request dto.MedicalCardUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(400)
		w.Write([]byte("Неверный JSON для изменения медицинской карты"))
		return
	}

	err = h.medicalCardService.UpdateMedicalCard(cardID, request.DiagnosisID, request.Recommendations)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(500)
		w.Write([]byte("Не удалось изменить медицинскую карту: " + err.Error()))
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)
	w.Write([]byte("Медицинская карта изменена: указан диагноз, карта закрыта"))
}
