package service

import (
	"gitlab.com/med-group5591241/med/internal/app/model"
	"gitlab.com/med-group5591241/med/internal/app/repository"
)

type MedicationService struct {
	medicationRepository *repository.MedicationRepository
}

func NewMedicationService(medicationRepository *repository.MedicationRepository) *MedicationService {
	return &MedicationService{medicationRepository: medicationRepository}
}

func (s *MedicationService) GetMedications() []model.Medication {
	return s.medicationRepository.GetMedications()
}
