package question

import (
	"errors"
	"fmt"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
)

type (
	FileQuestion struct {
		Basic
		FileType   FileType
		Constraint FileConstraint
	}
	FileType int
)

const (
	Image                       FileType = 1
	PDF                         FileType = 2
	Any                         FileType = 3
	FileQuestionFileTypeField            = "fileType"
	FileConstraintsCustomsField          = "fileConstraint"
)

func (t FileType) String() string {
	switch t {
	case Image:
		return "image"
	case PDF:
		return "pdf"
	case Any:
		return "any"
	default:
		return "unknown"
	}
}

func NewFileQuestion(
	id id.QuestionID, text string, fileType FileType, constraint FileConstraint, formID id.FormID,
) *FileQuestion {
	return &FileQuestion{
		Basic:      NewBasic(id, text, TypeFile, formID),
		FileType:   fileType,
		Constraint: constraint,
	}
}

func NewFileType(v int) (FileType, error) {
	switch FileType(v) {
	case Image, PDF, Any:
		return FileType(v), nil
	}
	return 0, errors.New("invalid file type")
}

func ImportFileQuestion(q StandardQuestion) (*FileQuestion, error) {
	// check if customs has "fileType" as int, return error if not
	fileTypeDataI, has := q.Customs[FileQuestionFileTypeField]
	if !has {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" is required for FileQuestion", FileQuestionFileTypeField))
	}
	fileTypeData, ok := fileTypeDataI.(int64)
	if !ok {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" must be int for FileQuestion", FileQuestionFileTypeField))
	}
	fileType, err := NewFileType(int(fileTypeData))
	if err != nil {
		return nil, err
	}

	if fileType == Any {
		return NewFileQuestion(q.ID, q.Text, fileType, nil, q.FormID), nil
	}

	constraintsCustomsData, has := q.Customs[FileConstraintsCustomsField]
	// if FileConstraintsCustomsField is not present, return FileQuestion without constraint
	if !has {
		return NewFileQuestion(q.ID, q.Text, fileType, nil, q.FormID), nil
	}

	constraintsCustoms, ok := constraintsCustomsData.(map[string]interface{})
	// if FileConstraintsCustomsField Found, but it is not map[string]interface{}, return error
	if !ok {
		return nil, errors.New(
			fmt.Sprintf("\"%s\" must be map[string]interface{} for FileQuestion", FileConstraintsCustomsField))
	}

	constraint := NewStandardFileConstraint(fileType, constraintsCustoms)
	question := NewFileQuestion(q.ID, q.Text, fileType, ImportFileConstraint(constraint), q.FormID)
	if err != nil {
		return nil, err
	}
	return question, nil
}

func (q FileQuestion) Export() StandardQuestion {
	customs := make(map[string]interface{})

	customs[FileQuestionFileTypeField] = q.FileType

	if q.Constraint != nil {
		customs[FileConstraintsCustomsField] = q.Constraint.Export().Customs
	}
	return NewStandardQuestion(TypeFile, q.ID, q.FormID, q.Text, customs)
}
