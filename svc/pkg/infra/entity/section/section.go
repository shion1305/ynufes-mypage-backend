package entity

import (
	"sort"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/section"
	"ynufes-mypage-backend/svc/pkg/domain/model/util"
)

const SectionRootName = "Sections"

type Section struct {
	ID id.SectionID `json:"-"`

	FormID id.FormID `json:"form_id"`

	// Questions map[QID]order
	// Order of questions are managed by fractional indexing.
	// Reference: https://en-jp.wantedly.com/companies/wantedly/post_articles/386188
	Questions map[string]float64 `json:"questions"`

	// ConditionQuestion a question which determines next section based on its answer
	// Only some of the questions can be condition questions. (e.g. radio, checkbox)
	// If !ConditionQuestion.HasValue(), then proceed to next section
	ConditionQuestion string `json:"c_question"`

	// ConditionCustoms map[OptionID]NextSectionID
	ConditionCustoms map[string]string `json:"c_customs"`
}

func NewSection(
	sectionID id.SectionID,
	formID id.FormID,
	questions map[string]float64,
	conditionQID string,
	conditionCustoms map[string]string,
) Section {
	return Section{
		ID:                sectionID,
		FormID:            formID,
		Questions:         questions,
		ConditionQuestion: conditionQID,
		ConditionCustoms:  conditionCustoms,
	}
}

func (s Section) ToModel() (*section.Section, error) {
	qs := make(map[id.QuestionID]float64, len(s.Questions))
	for k, v := range s.Questions {
		i, err := identity.ImportID(k)
		if err != nil {
			return nil, err
		}
		qs[i] = v
	}

	conditionCustoms := make(map[util.ID]id.SectionID, len(s.ConditionCustoms))
	for k, v := range s.ConditionCustoms {
		i, err := identity.ImportID(k)
		if err != nil {
			return nil, err
		}
		nextS, err := identity.ImportID(v)
		if err != nil {
			return nil, err
		}
		conditionCustoms[i] = nextS
	}

	cid, err := identity.ImportID(s.ConditionQuestion)
	if err != nil {
		return nil, err
	}

	sec := section.NewSection(
		s.ID,
		s.FormID,
		qs,
		cid,
		conditionCustoms,
	)
	return &sec, nil
}

type question struct {
	order int
	id    id.QuestionID
}

type questions []question

func sortQuestions(qs map[string]int) ([]id.QuestionID, error) {
	q, err := newQuestions(qs)
	if err != nil {
		return nil, err
	}
	sort.Sort(q)
	ids := make([]id.QuestionID, len(q))
	for i := range q {
		ids[i] = q[i].id
	}
	return ids, nil
}

func newQuestions(qs map[string]int) (questions, error) {
	q := make(questions, 0, len(qs))
	for k, v := range qs {
		tid, err := identity.ImportID(k)
		if err != nil {
			return nil, err
		}
		q = append(q, question{
			order: v,
			id:    tid,
		})
	}
	return q, nil
}

func (q questions) Len() int {
	return len(q)
}

func (q questions) Less(i, j int) bool {
	return q[i].order < q[j].order
}

func (q questions) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}
