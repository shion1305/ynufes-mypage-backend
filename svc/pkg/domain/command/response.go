package command

import (
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/response"
)

type Response interface {
	Create(ctx context.Context, resp *response.Response) error
	Set(ctx context.Context, resp response.Response) error
}
