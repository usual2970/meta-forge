package domain

import (
	"context"
	"mime/multipart"

	"github.com/tmc/langchaingo/schema"
)

type AiUploadDocByLinkReq struct {
	Link  string `json:"link"`
	Index string `json:"index"`
}
type AiUploadDocReq struct {
	Index string `json:"index" form:"index"`
	File  *multipart.FileHeader
}

type AiUploadDocResp struct {
	Index     string `json:"index" form:"index"`
	Namespace string `json:"namespace"`
}

type AiSearchDocReq struct {
	Question  string `json:"question" query:"question"`
	Index     string `json:"index" query:"index"`
	Namespace string `json:"namespace" query:"namespace"`
	Func      string `json:"func" query:"func"`
}

type AiSearchDocResp struct {
	Result map[string]any    `json:"result"`
	Docs   []schema.Document `json:"docs"`
}

type AiSummaryDocReq struct {
	Link string `json:"link"`
}

type AiSummaryDocResp struct {
	Result string `json:"result"`
}

type AiOrderNumReq struct {
	Question string `json:"question"`
}

type AiOrderNumResp struct {
	Result string `json:"result"`
}

type IAiUsecase interface {
	UploadDoc(ctx context.Context, req *AiUploadDocReq) (*AiUploadDocResp, error)

	UploadDocByLink(ctx context.Context, req *AiUploadDocByLinkReq) (*AiUploadDocResp, error)

	SearchDoc(ctx context.Context, req *AiSearchDocReq) (*AiSearchDocResp, error)

	SummaryDoc(ctx context.Context, req *AiSummaryDocReq) (*AiSummaryDocResp, error)

	OrderNum(ctx context.Context, req *AiOrderNumReq) (*AiOrderNumResp, error)
}
