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
