package ports

import (
	"context"

	"github.com/upgradeskill/beta-team/internal/requests"
)

type ICompanyService interface {
}

type ICompanyRepository interface {
	CreateCompany(ctx context.Context, input *requests.CreateCompanyRequest) (uint64, error)
}
