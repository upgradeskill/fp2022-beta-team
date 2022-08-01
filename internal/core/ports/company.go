package ports

import (
	"github.com/upgradeskill/beta-team/internal/core/domain"
)

type ICompanyRepository interface {
	CreateCompany(company *domain.Companies) error
}

type ICompanyUsecase interface {
	CreateCompany(company *domain.Companies) error
}
