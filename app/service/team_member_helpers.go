package service

import (
	"context"

	"github.com/adamnasrudin03/go-skeleton-gin/app/dto"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

func (s TeamMemberSrv) checkDuplicate(ctx context.Context, req dto.TeamMemberDetailReq) error {
	var (
		opName = "TeamMemberService-checkDuplicate"
		err    error
	)
	detail, err := s.Repo.GetDetail(ctx, dto.TeamMemberDetailReq{
		CustomColumn: "id",
		Email:        req.Email,
	})
	if err != nil {
		s.Logger.Errorf("%s, failed check duplicate email: %v", opName, err)
		return helpers.ErrDB()
	}

	if detail != nil && detail.ID > 0 {
		return helpers.ErrIsDuplicate("email", "email")
	}

	detail, err = s.Repo.GetDetail(ctx, dto.TeamMemberDetailReq{
		CustomColumn:   "id",
		UsernameGithub: req.UsernameGithub,
	})
	if err != nil {
		s.Logger.Errorf("%s, failed check duplicate username_github: %v", opName, err)
		return helpers.ErrDB()
	}

	if detail != nil && detail.ID > 0 {
		return helpers.ErrIsDuplicate("username_github", "username_github")
	}

	return nil

}
