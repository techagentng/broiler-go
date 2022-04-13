package service

import (
	"github.com/techagentng/boiler-go/application/company-repository"
	domain "github.com/techagentng/boiler-go/domain/company-repo"
)

// CompanyService interface
type CompanyService interface {
	CreateUser(user *domain.User) (*domain.User, error)
}

// DefaultWalletService struct
type DefaultWalletService struct {
	repo company_repository.Repository
}

func NewCompanyService(repo company_repository.Repository) *DefaultWalletService {
	return &DefaultWalletService{
		repo: repo,
	}
}
