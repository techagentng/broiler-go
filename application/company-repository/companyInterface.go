package company_repository

import (
	domain "github.com/techagentng/boiler-go/domain/company-repo"
)

// Repository interface
type Repository interface {
	CreateUser(user *domain.User) (*domain.User, error)
}
