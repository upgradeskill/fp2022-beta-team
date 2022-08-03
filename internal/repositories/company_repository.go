package repositories

import (
	"github.com/upgradeskill/beta-team/internal/core/domain"
	"gorm.io/gorm"
)

type companyRepository struct {
	DB *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) companyRepository {
	return companyRepository{DB: db}
}

func (cr *companyRepository) CreateCompany(company *domain.Companies) error {
	return cr.DB.Create(company).Error
}

func (cr *companyRepository) GetCompanyByID(id uint64) (*domain.Companies, error) {
	var company domain.Companies
	err := cr.DB.First(&company, id).Error
	if err != nil {
		return nil, err
	}
	return &company, nil
}

func (cr *companyRepository) ListCompanies() ([]*domain.Companies, error) {
	var companies []*domain.Companies
	err := cr.DB.Find(&companies).Error
	if err != nil {
		return nil, err
	}
	return companies, nil
}

func (cr *companyRepository) UpdateCompany(company *domain.Companies) error {
	return cr.DB.Save(company).Error
}

func (cr *companyRepository) DeleteCompany(id uint64) error {
	return cr.DB.Delete(&domain.Companies{ID: id}).Error
}
