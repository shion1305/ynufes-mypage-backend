package uc

import (
	"context"
	"ynufes-mypage-backend/pkg/jwt"
	"ynufes-mypage-backend/pkg/setting"
	userDomain "ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/domain/query"
	"ynufes-mypage-backend/svc/pkg/domain/service/util"
	"ynufes-mypage-backend/svc/pkg/registry"
)

type LoginUseCase struct {
	userQuery query.User
	jwtSecret string
	idManager util.IDManager
}

type LoginInput struct {
	JWT userDomain.JWT
}

type LoginOutput struct {
	User userDomain.User
}

func NewLoginUseCase(registry registry.Registry) LoginUseCase {
	config := setting.Get()
	return LoginUseCase{
		userQuery: registry.Repository().NewUserQuery(),
		jwtSecret: config.Application.Admin.JwtSecret,
		idManager: registry.Service().NewIDManager(),
	}
}

func (uc LoginUseCase) Do(ctx context.Context, input LoginInput) (*LoginOutput, error) {
	claims, err := jwt.Verify(input.JWT, uc.jwtSecret)
	if err != nil {
		return nil, err
	}
	id, err := uc.idManager.ImportID(claims.Id)
	if err != nil {
		return nil, err
	}
	userData, err := uc.userQuery.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &LoginOutput{
		User: *userData,
	}, nil
}
