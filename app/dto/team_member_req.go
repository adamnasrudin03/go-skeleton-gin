package dto

import (
	"github.com/adamnasrudin03/go-skeleton-gin/app/models"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

type TeamMemberDetailReq struct {
	ID             uint64 `json:"id"`
	Name           string `json:"name"`
	UsernameGithub string `json:"username_github"`
	Email          string `json:"email"`
	CustomColumn   string `json:"custom_column"`
	NotID          uint64 `json:"not_id"`
}

type TeamMemberCreateReq struct {
	Name           string `json:"name" validate:"required"`
	UsernameGithub string `json:"username_github" validate:"required"`
	Email          string `json:"email" validate:"required,email"`
}

type TeamMemberUpdateReq struct {
	ID             uint64 `json:"id" validate:"min=1"`
	Name           string `json:"name" validate:"required"`
	UsernameGithub string `json:"username_github" validate:"required"`
	Email          string `json:"email" validate:"required,email"`
}

type TeamMemberListReq struct {
	Search string `json:"search" form:"search"`
	models.BasedFilter
}

func (m *TeamMemberListReq) Validate() error {
	if m.Page <= 0 {
		m.Page = 1
	}

	if m.Limit <= 0 {
		m.Limit = 10
	}

	m.Search = helpers.ToLower(m.Search)

	m.OrderBy = helpers.ToUpper(m.OrderBy)
	if !models.IsValidOrderBy[m.OrderBy] && m.OrderBy != "" {
		return helpers.ErrInvalidFormat("order_by", "order_by")
	}

	m.SortBy = helpers.ToLower(m.SortBy)
	if m.OrderBy != "" && m.SortBy == "" {
		return helpers.ErrIsRequired("sort_by", "sort_by")
	}

	return nil
}
