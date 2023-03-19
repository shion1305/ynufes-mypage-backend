package writer

import (
	"context"
	"firebase.google.com/go/v4/db"
	"fmt"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/pkg/typecast"
	"ynufes-mypage-backend/svc/pkg/domain/model/response"
	"ynufes-mypage-backend/svc/pkg/exception"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/response"
)

type Response struct {
	ref *db.Ref
}

func NewResponse(f *db.Ref) Response {
	return Response{
		ref: f,
	}
}

func (w Response) Create(ctx context.Context, resp *response.Response) error {
	if resp.ID != nil && resp.ID.HasValue() {
		return exception.ErrIDAlreadyAssigned
	}
	if resp.OrgID == nil || !resp.OrgID.HasValue() ||
		resp.AuthorID == nil || !resp.AuthorID.HasValue() ||
		resp.FormID == nil || !resp.FormID.HasValue() {
		return exception.ErrIDNotAssigned
	}
	newID := identity.IssueID()
	data := make(map[string]entity.QuestionResponse, len(resp.Data))
	for key, val := range resp.Data {
		mapInterface, err := typecast.ConvertToStringMapInterface(val)
		if err != nil {
			return fmt.Errorf("failed to convert to entity QuestionResponse: %w", err)
		}
		data[key.ExportID()] = entity.NewQuestionResponse(key, mapInterface)
	}
	e := entity.NewResponse(
		newID, resp.OrgID.ExportID(), resp.AuthorID.ExportID(), resp.FormID.ExportID(), data,
	)
	if err := w.ref.Child(newID.ExportID()).Set(ctx, e); err != nil {
		return err
	}
	resp.ID = newID
	return nil
}

func (w Response) Set(ctx context.Context, resp response.Response) error {
	if resp.ID == nil || !resp.ID.HasValue() {
		return exception.ErrIDNotAssigned
	}
	data := make(map[string]entity.QuestionResponse, len(resp.Data))
	for key, val := range resp.Data {
		mapInterface, err := typecast.ConvertToStringMapInterface(val)
		if err != nil {
			return fmt.Errorf("failed to convert to entity QuestionResponse: %w", err)
		}
		data[key.ExportID()] = entity.NewQuestionResponse(key, mapInterface)
	}
	e := entity.NewResponse(
		resp.ID, resp.OrgID.ExportID(), resp.AuthorID.ExportID(), resp.FormID.ExportID(), data,
	)
	if err := w.ref.Child(resp.ID.ExportID()).Set(ctx, e); err != nil {
		return err
	}
	return nil
}
