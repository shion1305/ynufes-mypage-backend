package id

import "ynufes-mypage-backend/svc/pkg/domain/model/util"

type (
	UserID     util.ID
	OrgID      util.ID
	EventID    util.ID
	FormID     util.ID
	SectionID  util.ID
	QuestionID util.ID
	ResponseID util.ID
	OrgIDs     []OrgID
)

func (i OrgIDs) HasOrgID(oid OrgID) bool {
	for _, t := range i {
		if t == oid {
			return true
		}
	}
	return false
}
