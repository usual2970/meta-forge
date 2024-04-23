package data

import (
	"context"

	"github.com/usual2970/meta-forge/internal/domain"
)

type usecase struct {
	repo domain.IDataRepository
}

func NewUsecase(repo domain.IDataRepository) domain.IDataUsecase {
	return &usecase{
		repo: repo,
	}
}

func (u *usecase) List(ctx context.Context, req *domain.DataListReq) (*domain.DataListResp, error) {
	rs, err := u.repo.List(ctx, req)
	if err != nil {
		return nil, err
	}
	return rs, nil
}
