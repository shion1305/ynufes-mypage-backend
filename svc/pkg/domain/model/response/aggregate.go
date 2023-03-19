package response

import (
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
)

type (
	Response struct {
		ID       id.ResponseID
		OrgID    id.OrgID
		AuthorID id.UserID
		FormID   id.FormID
		Data     map[id.QuestionID]QuestionResponse
	}
	QuestionResponse struct {
		QuestionID   id.QuestionID
		ResponseData map[string]interface{}
	}
)

func NewResponse(
	rid id.ResponseID,
	oid id.OrgID,
	aid id.UserID,
	fid id.FormID,
	data map[id.QuestionID]QuestionResponse,
) Response {
	return Response{
		rid, oid, aid, fid, data,
	}
}

func NewQuestionResponse(
	qid id.QuestionID,
	data map[string]interface{},
) QuestionResponse {
	return QuestionResponse{qid, data}
}
