package reader

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/domain/service/util"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/user"
)

const (
	UserCollectionName = "users"
)

type (
	User struct {
		collection *firestore.CollectionRef
		idManager  util.IDManager
	}
)

func NewUser(c *firestore.Client) User {
	return User{
		collection: c.Collection("users"),
		idManager:  identity.NewIDManager(),
	}
}

func (u User) GetByID(ctx context.Context, id user.ID) (model *user.User, err error) {
	log.Printf("GET USER: %v", id)
	var userEntity entity.User
	snap, err := u.collection.Doc(id.ExportID()).Get(ctx)
	if err != nil {
		return nil, err
	}
	err = snap.DataTo(&userEntity)
	if err != nil {
		return nil, err
	}
	userEntity.ID = id
	model, err = userEntity.ToModel()
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (u User) GetByLineServiceID(ctx context.Context, id user.LineServiceID) (model *user.User, err error) {
	log.Printf("GET USER BY LINE ID: %v", id)
	var userEntity entity.User
	snap, err := u.collection.Where("line-id", "==", string(id)).Documents(ctx).Next()
	if err != nil {
		// user not found

		return nil, err
	}
	err = snap.DataTo(&userEntity)
	if err != nil {
		return nil, err
	}
	userEntity.ID, err = u.idManager.ImportID(snap.Ref.ID)
	if err != nil {
		return nil, err
	}
	model, err = userEntity.ToModel()
	if err != nil {
		return nil, err
	}
	return model, nil
}
