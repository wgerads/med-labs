package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"gitlab.com/med-group5591241/med/internal/app/handler/dto"
	"gitlab.com/med-group5591241/med/internal/app/model"
	"gitlab.com/med-group5591241/med/internal/app/service"
)

type PatientHandler struct {
	patientService *service.PatientService
}

func NewPatientHandler(patientService *service.PatientService) *PatientHandler {
	return &PatientHandler{patientService: patientService}
}

// CreatePatient godoc
// @summary     Создать пациента
// @description Регистрирует нового пациента. Дата рождения в формате ГГГГ-ММ-ДД, телефон можно не указывать
// @tags Patient
// @success 200
// @router /patients [post]
// @param request body dto.PatientCreateRequest true "request"
func (h *PatientHandler) CreatePatient(w http.ResponseWriter, r *http.Request) {
	var request dto.PatientCreateRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(400)
		w.Write([]byte("Неверный JSON для создания пациента"))
		return
	}

	birthDate, err := time.Parse("2006-01-02", request.BirthDate)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(400)
		w.Write([]byte("Дата рождения должна быть в формате ГГГГ-ММ-ДД"))
		return
	}

	patient := model.Patient{
		FullName:  request.FullName,
		BirthDate: birthDate,
		Phone:     request.Phone,
		Policy:    request.Policy,
	}

	err = h.patientService.CreatePatient(patient)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(500)
		w.Write([]byte("Не удалось создать пациента"))
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)
	w.Write([]byte("Пациент создан"))
}

// GetPatient godoc
// @summary     Получить пациента
// @description Возвращает пациента по идентификатору
// @tags Patient
// @success 200 {object} dto.PatientResponse
// @router /patients/{id} [get]
// @param id path int true "id"
func (h *PatientHandler) GetPatient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	patientID, _ := strconv.Atoi(vars["id"])

	patient := h.patientService.GetPatient(patientID)
	if patient == nil {
		w.Write([]byte("Пациент не найден"))
		return
	}

	patientResponse := dto.PatientResponse{
		ID:        patient.ID,
		FullName:  patient.FullName,
		BirthDate: patient.BirthDate,
		Phone:     patient.Phone,
		Policy:    patient.Policy,
	}

	w.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(&patientResponse)
	w.Write(b)
}

// GetPatients godoc
// @summary     Список пациентов
// @description Возвращает всех пациентов
// @tags Patient
// @success 200 {object} []dto.PatientResponse
// @router /patients [get]
func (h *PatientHandler) GetPatients(w http.ResponseWriter, r *http.Request) {
	patients := h.patientService.GetPatients()
	if len(patients) == 0 {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Пациентов нет"))
		return
	}

	patientsResponse := []dto.PatientResponse{}
	for _, patient := range patients {
		patientsResponse = append(patientsResponse, dto.PatientResponse{
			ID:        patient.ID,
			FullName:  patient.FullName,
			BirthDate: patient.BirthDate,
			Phone:     patient.Phone,
			Policy:    patient.Policy,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(patientsResponse)
	w.Write(b)
}
