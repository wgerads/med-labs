package service

import (
	"gitlab.com/med-group5591241/med/internal/app/model"
	"gitlab.com/med-group5591241/med/internal/app/repository"
)

type PrescriptionService struct {
	prescriptionRepository *repository.PrescriptionRepository
}

func NewPrescriptionService(prescriptionRepository *repository.PrescriptionRepository) *PrescriptionService {
	return &PrescriptionService{prescriptionRepository: prescriptionRepository}
}

func (s *PrescriptionService) CreatePrescription(prescription model.Prescription) error {
	return s.prescriptionRepository.CreatePrescription(prescription)
}

func (s *PrescriptionService) DeletePrescription(id int) error {
	return s.prescriptionRepository.DeletePrescription(id)
}
