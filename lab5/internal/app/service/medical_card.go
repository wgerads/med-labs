package service

import (
	"gitlab.com/med-group5591241/med/internal/app/model"
	"gitlab.com/med-group5591241/med/internal/app/repository"
)

type MedicalCardService struct {
	medicalCardRepository *repository.MedicalCardRepository
}

func NewMedicalCardService(medicalCardRepository *repository.MedicalCardRepository) *MedicalCardService {
	return &MedicalCardService{medicalCardRepository: medicalCardRepository}
}

func (s *MedicalCardService) CreateMedicalCard(card model.MedicalCard) error {
	return s.medicalCardRepository.CreateMedicalCard(card)
}

func (s *MedicalCardService) GetMedicalCard(id int) *model.MedicalCard {
	return s.medicalCardRepository.GetMedicalCard(id)
}

func (s *MedicalCardService) UpdateMedicalCard(id int, diagnosisID int, recommendations string) error {
	return s.medicalCardRepository.UpdateMedicalCard(id, diagnosisID, recommendations)
}
