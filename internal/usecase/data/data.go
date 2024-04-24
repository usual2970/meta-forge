package data

import (
	"context"

	"github.com/usual2970/meta-forge/internal/domain"
	"github.com/usual2970/meta-forge/internal/repository/data"
)

type usecase struct {
	repo domain.IDataRepository
}

func NewUsecase(repo domain.IDataRepository) domain.IDataUsecase {
	return &usecase{
		repo: repo,
	}
}

func (u *usecase) getRp(ctx context.Context) domain.IDataRepository {
	if u.repo != nil {
		return u.repo
	}

	u.repo = data.NewRepository()
	return u.repo
}

func (u *usecase) List(ctx context.Context, req *domain.DataListReq) (*domain.DataListResp, error) {
	rs, err := u.getRp(ctx).List(ctx, req)
	if err != nil {
		return nil, err
	}
	return rs, nil
}
