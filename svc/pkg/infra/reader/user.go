package reader

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
	"strconv"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/user"
)

const (
	UserCollectionName = "users"
)

type (
	User struct {
		Collection *firestore.CollectionRef
	}
)

func NewUser(c *firestore.Client) User {
	return User{
		c.Collection("users"),
	}
}

func (u User) GetByID(ctx context.Context, id user.ID) (model *user.User, err error) {
	log.Printf("GET USER: %v", id)
	var userEntity entity.User
	snap, err := u.Collection.Doc(strconv.FormatInt(int64(id), 10)).Get(ctx)
	if err != nil {
		return nil, err
	}
	err = snap.DataTo(&userEntity)
	if err != nil {
		return nil, err
	}
	model, err = userEntity.ToModel()
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (u User) GetByLineServiceID(ctx context.Context, id user.LineServiceID) (model *user.User, err error) {
	log.Printf("GET USER BY LINE ID: %v", id)
	var userEntity entity.User
	snap, err := u.Collection.Where("line-id", "==", string(id)).Documents(ctx).Next()
	if err != nil {
		// user not found

		return nil, err
	}
	err = snap.DataTo(&userEntity)
	if err != nil {
		return nil, err
	}
	userEntity.ID = snap.Ref.ID
	model, err = userEntity.ToModel()
	if err != nil {
		return nil, err
	}
	return model, nil
}
