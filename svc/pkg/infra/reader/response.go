package reader

import (
	"context"
	"firebase.google.com/go/v4/db"
	"fmt"
	"ynufes-mypage-backend/pkg/firebase"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/response"
	"ynufes-mypage-backend/svc/pkg/exception"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/response"
)

type Response struct {
	ref *db.Ref
}

func NewResponse(f firebase.Firebase) Response {
	return Response{
		ref: f.Client(entity.ResponseRootName),
	}
}

func (r Response) GetByID(ctx context.Context, oid id.ResponseID) (*response.Response, error) {
	var responseEntity entity.Response
	hits, err := r.ref.OrderByKey().
		EqualTo(oid.ExportID()).GetOrdered(ctx)
	if err != nil {
		return nil, err
	}
	if len(hits) == 0 {
		return nil, exception.ErrNotFound
	}
	if err := hits[0].Unmarshal(&responseEntity); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response entity: %w", err)
	}
	responseEntity.ID = oid
	m, err := responseEntity.ToModel()
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r Response) ListByFormID(ctx context.Context, fid id.FormID) ([]response.Response, error) {
	hits, err := r.ref.OrderByChild("form_id").
		EqualTo(fid.ExportID()).GetOrdered(ctx)
	if err != nil {
		return nil, err
	}
	if len(hits) == 0 {
		return nil, exception.ErrNotFound
	}
	resps := make([]response.Response, 0, len(hits))
	for _, hit := range hits {
		var e entity.Response
		if err := hit.Unmarshal(&e); err != nil {
			return nil, fmt.Errorf("failed to unmarshal response entity: %w", err)
		}
		e.ID = fid
		m, err := e.ToModel()
		if err != nil {
			return nil, err
		}
		resps = append(resps, *m)
	}
	return resps, nil
}

func (r Response) ListByOrgID(ctx context.Context, oid id.OrgID) ([]response.Response, error) {
	hits, err := r.ref.OrderByChild("org_id").
		EqualTo(oid.ExportID()).GetOrdered(ctx)
	if err != nil {
		return nil, err
	}
	resps := make([]response.Response, 0, len(hits))
	for _, hit := range hits {
		var e entity.Response
		if err := hit.Unmarshal(&e); err != nil {
			return nil, fmt.Errorf("failed to unmarshal response entity: %w", err)
		}
		e.ID = oid
		m, err := e.ToModel()
		if err != nil {
			return nil, err
		}
		resps = append(resps, *m)
	}
	return resps, nil
}

func (r Response) ListByAuthorID(ctx context.Context, uid id.UserID) ([]response.Response, error) {
	hits, err := r.ref.OrderByChild("author_id").
		EqualTo(uid.ExportID()).GetOrdered(ctx)
	if err != nil {
		return nil, err
	}
	resps := make([]response.Response, 0, len(hits))
	for _, hit := range hits {
		var e entity.Response
		if err := hit.Unmarshal(&e); err != nil {
			return nil, fmt.Errorf("failed to unmarshal response entity: %w", err)
		}
		e.ID = uid
		m, err := e.ToModel()
		if err != nil {
			return nil, err
		}
		resps = append(resps, *m)
	}
	return resps, nil
}
