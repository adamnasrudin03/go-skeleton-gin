package dto

import (
	"testing"

	"github.com/adamnasrudin03/go-skeleton-gin/app/models"
)

func TestTeamMemberListReq_Validate(t *testing.T) {
	tests := []struct {
		name    string
		m       *TeamMemberListReq
		wantErr bool
	}{
		{
			name: "invalid order by",
			m: &TeamMemberListReq{
				BasedFilter: models.BasedFilter{
					OrderBy: "invalid",
					SortBy:  "",
				},
			},
			wantErr: true,
		},
		{
			name: "sort by required if order by provided",
			m: &TeamMemberListReq{
				BasedFilter: models.BasedFilter{
					OrderBy: models.OrderByASC,
					SortBy:  "",
				},
			},
			wantErr: true,
		},
		{
			name:    "success",
			m:       &TeamMemberListReq{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("TeamMemberListReq.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
