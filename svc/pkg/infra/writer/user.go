package writer

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/user"
)

type User struct {
	collection *firestore.CollectionRef
}

func NewUser(c *firestore.Client) User {
	return User{
		collection: c.Collection("users"),
	}
}

func (u User) Create(ctx context.Context, model user.User) error {
	log.Printf("CREATE USER: %v", model)
	e := entity.User{
		//ID is not required as it will not be used by firestore
		//ID: int64(model.ID),
		Status: int(model.Status),
		UserDetail: entity.UserDetail{
			NameFirst:     model.Detail.Name.FirstName,
			NameFirstKana: model.Detail.Name.FirstNameKana,
			NameLast:      model.Detail.Name.LastName,
			NameLastKana:  model.Detail.Name.LastNameKana,
			Gender:        int(model.Detail.Gender),
			StudentID:     string(model.Detail.StudentID),
			Email:         string(model.Detail.Email),
			Type:          int(model.Detail.Type),
		},
		Line: entity.Line{
			LineServiceID:         string(model.Line.LineServiceID),
			LineProfileURL:        string(model.Line.LineProfilePictureURL),
			LineDisplayName:       model.Line.LineDisplayName,
			EncryptedAccessToken:  string(model.Line.EncryptedAccessToken),
			EncryptedRefreshToken: string(model.Line.EncryptedRefreshToken),
		},
	}
	//NOTE: Create fails if the document already exists
	_, err := u.collection.Doc(model.ID.ExportID()).
		Create(ctx, e)
	if err != nil {
		return err
	}
	return nil
}

func (u User) UpdateAll(ctx context.Context, model user.User) error {
	log.Printf("UPDATE USER: %v", model)
	e := entity.User{
		Status: int(model.Status),
		UserDetail: entity.UserDetail{
			NameFirst:     model.Detail.Name.FirstName,
			NameFirstKana: model.Detail.Name.FirstNameKana,
			NameLast:      model.Detail.Name.LastName,
			NameLastKana:  model.Detail.Name.LastNameKana,
			Gender:        int(model.Detail.Gender),
			StudentID:     string(model.Detail.StudentID),
			Email:         string(model.Detail.Email),
			Type:          int(model.Detail.Type),
		},
		Line: entity.Line{},
	}
	_, err := u.collection.Doc(model.ID.ExportID()).
		Set(ctx, e)
	return err
}

func (u User) UpdateLine(ctx context.Context, oldUser *user.User, update user.Line) error {
	log.Printf("UPDATE USER LINE: %v -> %v\n", oldUser, update)
	targets := map[string]struct {
		oldValue string
		newValue string
	}{
		"line-id": {
			oldValue: string(oldUser.Line.LineServiceID),
			newValue: string(update.LineServiceID),
		},
		"line-profile_url": {
			oldValue: string(oldUser.Line.LineProfilePictureURL),
			newValue: string(update.LineProfilePictureURL),
		},
		"line-display_name": {
			oldValue: oldUser.Line.LineDisplayName,
			newValue: update.LineDisplayName,
		},
		"line-access_token": {
			oldValue: string(oldUser.Line.EncryptedAccessToken),
			newValue: string(update.EncryptedAccessToken),
		},
		"line-refresh_token": {
			oldValue: string(oldUser.Line.EncryptedRefreshToken),
			newValue: string(update.EncryptedRefreshToken),
		},
	}
	var updates []firestore.Update
	for key, value := range targets {
		if value.oldValue != value.newValue {
			updates = append(updates, firestore.Update{Path: key, Value: value.newValue})
		}
	}
	if len(updates) == 0 {
		return nil
	}
	_, err := u.collection.
		Doc(oldUser.ID.ExportID()).
		Update(ctx, updates)
	if err == nil {
		oldUser.Line = update
	}
	return err
}

func (u User) UpdateUserDetail(ctx context.Context, oldUser *user.User, update user.Detail) error {
	log.Printf("UPDATE USER INFO: %v -> %v\n", oldUser, update)

	var updateTargets []firestore.Update
	var newStatus int
	switch update.MeetsBasicRequirement() {
	case true:
		newStatus = int(user.StatusRegistered)
		break
	case false:
		newStatus = int(user.StatusNew)
		break
	}
	targets := map[string]struct {
		oldValue interface{}
		newValue interface{}
	}{
		"detail-name_first": {
			oldValue: oldUser.Detail.Name.FirstName,
			newValue: update.Name.FirstName,
		},
		"detail-name_first_kana": {
			oldValue: oldUser.Detail.Name.FirstNameKana,
			newValue: update.Name.FirstNameKana,
		},
		"detail-name_last": {
			oldValue: oldUser.Detail.Name.LastName,
			newValue: update.Name.LastName,
		},
		"detail-name_last_kana": {
			oldValue: oldUser.Detail.Name.LastNameKana,
			newValue: update.Name.LastNameKana,
		},
		"detail-email": {
			oldValue: string(oldUser.Detail.Email),
			newValue: update.Email,
		},
		"detail-gender": {
			oldValue: oldUser.Detail.Gender,
			newValue: update.Gender,
		},
		"detail-student_id": {
			oldValue: string(oldUser.Detail.StudentID),
			newValue: update.StudentID,
		},
		"status": {
			oldValue: oldUser.Status,
			newValue: newStatus,
		},
	}
	for key, value := range targets {
		if value.oldValue != value.newValue {
			updateTargets = append(updateTargets, firestore.Update{Path: key, Value: value.newValue})
		}
	}
	if len(updateTargets) == 0 {
		return nil
	}
	_, err := u.collection.Doc(oldUser.ID.ExportID()).
		Update(ctx, updateTargets)
	if err == nil {
		// do not update field: Type
		update.Type = oldUser.Detail.Type
		oldUser.Detail = update
	}
	return err
}

func (u User) Delete(ctx context.Context, model user.User) error {
	log.Printf("DELETE USER: %v", model)
	_, err := u.collection.Doc(model.ID.ExportID()).
		Delete(ctx)
	return err
}
