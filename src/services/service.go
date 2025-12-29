package services

import (
	"oidc_project/repository"
	repo "oidc_project/repository"
)

type MainService struct {
	repo repo.MainRepository
}

func NewMainService(repo repository.MainRepository) *MainService {
	return &MainService{repo: repo}
}

func (m MainService) WriteData(data string) error {
	err := m.repo.WriteData(data)
	return err
}

func (m MainService) ReadData() (string, error) {
	data, err := m.repo.ReadData()
	if err != nil {
		return "", err
	}
	return data, nil
}

func (m MainService) CleanData() error {
	err := m.repo.CleanData()
	return err
}
