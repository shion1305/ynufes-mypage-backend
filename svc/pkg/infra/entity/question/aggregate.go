package entity

import (
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/question"
)

const QuestionRootName = "Questions"

type Question struct {
	ID      id.QuestionID          `json:"-"`
	FormID  string                 `json:"form_id"`
	Text    string                 `json:"text"`
	Type    int                    `json:"type"`
	Customs map[string]interface{} `json:"customs"`
}

func NewQuestion(
	id id.QuestionID,
	formID string,
	text string,
	qType int,
	customs map[string]interface{},
) Question {
	return Question{
		ID:      id,
		FormID:  formID,
		Text:    text,
		Type:    qType,
		Customs: customs,
	}
}

func (q Question) ToModel() (question.Question, error) {
	fid, err := identity.ImportID(q.FormID)
	if err != nil {
		return nil, err
	}
	sq := question.NewStandardQuestion(
		question.Type(q.Type),
		q.ID,
		fid,
		q.Text,
		q.Customs,
	)
	return sq.ToQuestion()
}
