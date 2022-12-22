package writer

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
	"strconv"
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
			LineServiceID: string(model.Line.LineServiceID),
			//EncryptedAccessToken:  string(model.Line.EncryptedAccessToken),
			//EncryptedRefreshToken: string(model.Line.EncryptedRefreshToken),
		},
		//UserDashboard: entity.UserDashboard{
		//	Grants: model.Dashboard.Grants,
		//},
	}
	//NOTE: Create fails if the document already exists
	_, err := u.collection.Doc(strconv.FormatInt(int64(model.ID), 10)).
		Create(ctx, e)
	if err != nil {
		return err
	}
	return nil
}

func (u User) UpdateAll(ctx context.Context, model user.User) error {
	log.Printf("UPDATE USER: %v", model)
	_, err := u.collection.Doc(strconv.FormatInt(int64(model.ID), 10)).
		Set(ctx, model)
	return err
}

func (u User) UpdateLine(ctx context.Context, model user.User) error {
	log.Printf("UPDATE USER LINE: %v", model)
	_, err := u.collection.Doc(strconv.FormatInt(int64(model.ID), 10)).
		Update(ctx, []firestore.Update{
			{Path: "line-id", Value: string(model.Line.LineServiceID)},
			{Path: "line-profile_url", Value: string(model.Line.LineProfilePictureURL)},
			{Path: "line-access_token", Value: string(model.Line.EncryptedAccessToken)},
			{Path: "line-refresh_token", Value: string(model.Line.EncryptedRefreshToken)},
			{Path: "line-display_name", Value: model.Line.LineDisplayName},
		})
	return err
}

func (u User) Delete(ctx context.Context, model user.User) error {
	log.Printf("DELETE USER: %v", model)
	_, err := u.collection.Doc(strconv.FormatInt(int64(model.ID), 10)).
		Delete(ctx)
	return err
}
