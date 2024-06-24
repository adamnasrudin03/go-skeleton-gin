package models

import (
	"time"
)

// TeamMember represents the model for an TeamMember
type TeamMember struct {
	ID              uint64     `json:"id" gorm:"primaryKey"`
	Name            string     `json:"name" gorm:"not null"`
	Role            string     `json:"role" gorm:"not null;default:'USER'"`
	Username        string     `json:"username" gorm:"not null;uniqueIndex"`
	Email           string     `json:"email" gorm:"not null;uniqueIndex"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"  gorm:"null; default:null"`
	Password        string     `json:"password,omitempty" gorm:"not null"`
	Salt            string     `json:"salt,omitempty" gorm:"not null"`
	DefaultModel
}
