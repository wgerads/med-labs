package service

import (
	"gitlab.com/med-group5591241/med/internal/app/model"
	"gitlab.com/med-group5591241/med/internal/app/repository"
)

type WardService struct {
	wardRepository *repository.WardRepository
}

func NewWardService(wardRepository *repository.WardRepository) *WardService {
	return &WardService{wardRepository: wardRepository}
}

func (s *WardService) GetWards() []model.Ward {
	return s.wardRepository.GetWards()
}
