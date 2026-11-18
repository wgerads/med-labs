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

type PrescriptionHandler struct {
	prescriptionService *service.PrescriptionService
}

func NewPrescriptionHandler(prescriptionService *service.PrescriptionService) *PrescriptionHandler {
	return &PrescriptionHandler{prescriptionService: prescriptionService}
}

// CreatePrescription godoc
// @summary     Назначить медикамент
// @description Добавляет назначение в медицинскую карту
// @tags Prescription
// @success 200
// @router /prescriptions [post]
// @param request body dto.PrescriptionCreateRequest true "request"
func (h *PrescriptionHandler) CreatePrescription(w http.ResponseWriter, r *http.Request) {
	var request dto.PrescriptionCreateRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(400)
		w.Write([]byte("Неверный JSON для назначения медикамента"))
		return
	}

	prescription := model.Prescription{
		MedicalCardID: request.MedicalCardID,
		MedicationID:  request.MedicationID,
		Frequency:     request.Frequency,
		Duration:      request.Duration,
	}

	err = h.prescriptionService.CreatePrescription(prescription)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(500)
		w.Write([]byte("Не удалось назначить медикамент"))
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)
	w.Write([]byte("Медикамент назначен"))
}

// DeletePrescription godoc
// @summary     Отменить назначение
// @description Удаляет назначение по идентификатору
// @tags Prescription
// @success 200
// @router /prescriptions/{id} [delete]
// @param id path int true "id"
func (h *PrescriptionHandler) DeletePrescription(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	prescriptionID, _ := strconv.Atoi(vars["id"])

	err := h.prescriptionService.DeletePrescription(prescriptionID)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(500)
		w.Write([]byte("Не удалось отменить назначение: " + err.Error()))
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)
	w.Write([]byte("Назначение отменено"))
}
