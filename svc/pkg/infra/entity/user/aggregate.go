package entity

import (
	"time"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
)

const UserCollectionName = "Users"

type User struct {
	// ignore id from firestore
	ID     user.ID `firestore:"-"`
	Status int     `firestore:"status"`
	UserDetail
	Line
	Admin
	Agent
}

func (u User) ToModel() (*user.User, error) {
	email, err := user.NewEmail(u.Email)
	if err != nil {
		return nil, err
	}
	gender, err := user.NewGender(u.Gender)
	if err != nil {
		return nil, err
	}
	roles := make([]user.Role, len(u.Roles))
	for i, role := range u.Roles {
		lv, err := user.NewRoleLevel(role.Level)
		if err != nil {
			return nil, err
		}
		roles[i] = user.Role{
			ID:          identity.NewID(role.ID),
			Level:       lv,
			GrantedTime: time.UnixMilli(role.GrantedTime),
		}
	}
	return &user.User{
		ID:     u.ID,
		Status: user.Status(u.Status),
		Detail: user.Detail{
			Name: user.Name{
				FirstName:     u.NameFirst,
				LastName:      u.NameLast,
				FirstNameKana: u.NameFirstKana,
				LastNameKana:  u.NameLastKana,
			},
			Email:  email,
			Gender: gender,
			// TODO: add validation for StudentID, Type
			StudentID: user.StudentID(u.StudentID),
			Type:      user.Type(u.Type),
		},
		Line: user.Line{
			LineServiceID:         user.LineServiceID(u.LineServiceID),
			LineProfilePictureURL: user.LineProfilePictureURL(u.LineProfileURL),
			LineDisplayName:       u.LineDisplayName,
			EncryptedAccessToken:  user.EncryptedAccessToken(u.EncryptedAccessToken),
			EncryptedRefreshToken: user.EncryptedRefreshToken(u.EncryptedRefreshToken),
		},
		Admin: user.Admin{
			IsSuperAdmin: u.IsSuperAdmin,
			GrantedTime:  time.UnixMilli(u.GrantedTime),
		},
		Agent: user.Agent{
			Roles: roles,
		},
	}, nil
}
