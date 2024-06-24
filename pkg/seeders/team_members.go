package seeders

import (
	"github.com/adamnasrudin03/go-skeleton-gin/app/models"
	"gorm.io/gorm"
)

func InitTeamMembers(db *gorm.DB) {
	tx := db.Begin()
	var teamMembers = []models.TeamMember{}
	tx.Select("id").Limit(1).Find(&teamMembers)
	if len(teamMembers) == 0 {
		teamMembers = []models.TeamMember{
			{},
		}
		tx.Create(&teamMembers)
	}

	tx.Commit()
}
