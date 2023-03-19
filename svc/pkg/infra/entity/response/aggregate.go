package entity

import (
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/response"
)

const ResponseRootName = "Responses"

type (
	Response struct {
		ID                id.ResponseID               `json:"-"`
		OrgID             string                      `json:"org_id"`
		AuthorID          string                      `json:"author_id"`
		FormID            string                      `json:"form_id"`
		QuestionResponses map[string]QuestionResponse `json:"questions"`
	}
	QuestionResponse struct {
		QuestionID   id.QuestionID          `json:"-"`
		ResponseData map[string]interface{} `json:"data"`
	}
)

func NewResponse(
	rid id.ResponseID,
	oid string,
	aid string,
	fid string,
	data map[string]QuestionResponse,
) Response {
	return Response{
		rid, oid, aid, fid, data,
	}
}

func (r Response) ToModel() (*response.Response, error) {
	orgID, err := identity.ImportID(r.OrgID)
	if err != nil {
		return nil, err
	}
	authorID, err := identity.ImportID(r.AuthorID)
	if err != nil {
		return nil, err
	}
	formID, err := identity.ImportID(r.FormID)
	if err != nil {
		return nil, err
	}

	return &response.Response{
		ID:       r.ID,
		OrgID:    orgID,
		AuthorID: authorID,
		FormID:   formID,
		Data:     r.Data,
	}, nil
}
