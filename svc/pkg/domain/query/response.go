package query

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/response"
)

type Response interface {
	GetByID(ctx context.Context, oid id.ResponseID) (*response.Response, error)
	ListByFormID(ctx context.Context, fid id.FormID) ([]response.Response, error)
	ListByOrgID(ctx context.Context, oid id.OrgID) ([]response.Response, error)
	ListByAuthorID(ctx context.Context, uid id.UserID) ([]response.Response, error)
}
