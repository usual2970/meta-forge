package domain

import "context"

type MingdaoGetPassIdReq struct {
	UserName string `json:"userName"`
	Password  string `json:"password"`
}
type IMingdaoUsecase interface {
	GetPassId(ctx context.Context, param *MingdaoGetPassIdReq) (string, error)
}
