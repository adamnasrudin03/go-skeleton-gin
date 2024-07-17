package dto

import (
	help "github.com/adamnasrudin03/go-helpers"
	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/adamnasrudin03/go-skeleton-gin/app/models"
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

	m.Search = help.ToLower(m.Search)

	m.OrderBy = help.ToUpper(m.OrderBy)
	if !models.IsValidOrderBy[m.OrderBy] && m.OrderBy != "" {
		return response_mapper.ErrInvalidFormat("order_by", "order_by")
	}

	m.SortBy = help.ToLower(m.SortBy)
	if m.OrderBy != "" && m.SortBy == "" {
		return response_mapper.ErrIsRequired("sort_by", "sort_by")
	}

	return nil
}
