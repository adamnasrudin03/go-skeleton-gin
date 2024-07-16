package dto

type TeamMemberDetailReq struct {
	ID             uint64 `json:"id"`
	Name           string `json:"name"`
	UsernameGithub string `json:"username_github"`
	Email          string `json:"email"`
	CustomColumn   string `json:"custom_column"`
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
