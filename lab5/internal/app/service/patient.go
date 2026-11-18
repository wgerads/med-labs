package service

import (
	"gitlab.com/med-group5591241/med/internal/app/model"
	"gitlab.com/med-group5591241/med/internal/app/repository"
)

type PatientService struct {
	patientRepository *repository.PatientRepository
}

func NewPatientService(patientRepository *repository.PatientRepository) *PatientService {
	return &PatientService{patientRepository: patientRepository}
}

func (s *PatientService) CreatePatient(patient model.Patient) error {
	return s.patientRepository.CreatePatient(patient)
}

func (s *PatientService) GetPatient(id int) *model.Patient {
	return s.patientRepository.GetPatient(id)
}

func (s *PatientService) GetPatients() []model.Patient {
	return s.patientRepository.GetPatients()
}
