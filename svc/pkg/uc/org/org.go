package org

import (
	"context"
	"fmt"
	"ynufes-mypage-backend/svc/pkg/domain/model/form"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/org"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/domain/query"
	"ynufes-mypage-backend/svc/pkg/domain/service/access"
	"ynufes-mypage-backend/svc/pkg/exception"
	"ynufes-mypage-backend/svc/pkg/registry"
)

type (
	OrgUseCase struct {
		orgQ    query.Org
		formQ   query.Form
		accessC *access.AccessController
	}
	OrgInput struct {
		Ctx   context.Context
		User  user.User
		OrgID id.OrgID
	}
	OrgOutput struct {
		Forms []form.Form
		Org   org.Org
	}
)

func NewOrg(rgst registry.Registry) OrgUseCase {
	return OrgUseCase{
		orgQ:    rgst.Repository().NewOrgQuery(),
		formQ:   rgst.Repository().NewFormQuery(),
		accessC: rgst.Service().AccessController(),
	}
}

func (uc OrgUseCase) Do(ipt OrgInput) (*OrgOutput, error) {
	targetOrg, err := uc.orgQ.GetByID(ipt.Ctx, ipt.OrgID)
	if err != nil {
		return nil, fmt.Errorf("failed to get targetOrg in OrgUC: %w", err)
	}

	if !uc.accessC.CanAccessOrg(ipt.Ctx, ipt.User.ID, targetOrg.ID) {
		return nil, exception.ErrUnauthorized
	}

	forms, err := uc.formQ.ListByEventID(ipt.Ctx, targetOrg.Event.ID)
	if err != nil {
		return nil, err
	}
	return &OrgOutput{
		Forms: forms,
		Org:   *targetOrg,
	}, nil
}
