package registry

import (
	"ynufes-mypage-backend/pkg/firebase"
	"ynufes-mypage-backend/svc/pkg/domain/command"
	"ynufes-mypage-backend/svc/pkg/domain/query"
	"ynufes-mypage-backend/svc/pkg/infra/reader"
	"ynufes-mypage-backend/svc/pkg/infra/writer"
)

var repo Repository

type Repository struct {
	fb *firebase.Firebase
}

func NewRepository() (Repository, error) {
	fb := firebase.New()
	repo = Repository{
		fb: &fb,
	}
	return repo, nil
}

func (repo Repository) NewUserQuery() query.User {
	return reader.NewUser(repo.fb)
}

func (repo Repository) NewUserCommand() command.User {
	return writer.NewUser(repo.fb)
}

func (repo Repository) NewEventQuery() query.Event {
	return reader.NewEvent(repo.fb)
}

func (repo Repository) NewEventCommand() command.Event {
	return writer.NewEvent(repo.fb)
}

func (repo Repository) NewOrgQuery() query.Org {
	return reader.NewOrg(repo.fb)
}

func (repo Repository) NewOrgCommand() command.Org {
	return writer.NewOrg(repo.fb)
}

func (repo Repository) NewFormCommand() command.Form {
	return writer.NewForm(repo.fb)
}

func (repo Repository) NewFormQuery() query.Form {
	return reader.NewForm(repo.fb)
}

func (repo Repository) NewQuestionQuery() query.Question {
	return reader.NewQuestion(repo.fb)
}

func (repo Repository) NewQuestionCommand() command.Question {
	return writer.NewQuestion(repo.fb)
}

func (repo Repository) NewRelationQuery() query.Relation {
	return reader.NewRelation(repo.fb)
}

func (repo Repository) NewRelationCommand() command.Relation {
	return writer.NewRelation(repo.fb)
}

func (repo Repository) NewSectionCommand() command.Section {
	return writer.NewSection(repo.fb)
}

func (repo Repository) NewSectionQuery() query.Section {
	return reader.NewSection(repo.fb)
}

func (repo Repository) NewLineCommand() command.Line {
	return writer.NewLine(repo.fb)
}

func (repo Repository) NewLineQuery() query.Line {
	return reader.NewLine(repo.fb)
}

func (repo Repository) NewResponseCommand() command.Response {
	return writer.NewResponse(repo.fb)
}

func (repo Repository) NewResponseQuery() query.Response {
	return reader.NewResponse(repo.fb)
}
