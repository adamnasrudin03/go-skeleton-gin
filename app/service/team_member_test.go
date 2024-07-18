package service

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/adamnasrudin03/go-skeleton-gin/app/dto"
	"github.com/adamnasrudin03/go-skeleton-gin/app/models"
	"github.com/adamnasrudin03/go-skeleton-gin/app/repository/mocks"
	"github.com/adamnasrudin03/go-skeleton-gin/configs"
	"github.com/adamnasrudin03/go-skeleton-gin/pkg/driver"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TeamMemberServiceTestSuite struct {
	suite.Suite
	repo    *mocks.TeamMemberRepository
	ctx     context.Context
	service TeamMemberService
}

func (srv *TeamMemberServiceTestSuite) SetupTest() {
	var (
		cfg    = configs.GetInstance()
		logger = driver.Logger(cfg)
	)

	srv.repo = &mocks.TeamMemberRepository{}
	srv.ctx = context.Background()
	srv.service = NewTeamMemberService(srv.repo, cfg, logger)
}

func TestTeamMemberService(t *testing.T) {
	suite.Run(t, new(TeamMemberServiceTestSuite))
}

func (srv *TeamMemberServiceTestSuite) TestTeamMemberSrv_GetByID() {
	tests := []struct {
		name     string
		id       uint64
		mockFunc func(input uint64)
		want     *models.TeamMember
		wantErr  bool
	}{
		{
			name: "Success with cache",
			id:   1,
			mockFunc: func(input uint64) {
				key := models.KeyCacheTeamMemberDetail(input)
				res := models.TeamMember{
					ID: input,
				}
				srv.repo.On("GetCache", mock.Anything, key, &models.TeamMember{
					ID: 0,
				}).Return(true).Run(func(args mock.Arguments) {
					target := args.Get(2).(*models.TeamMember)
					*target = res
				}).Once()
			},
			want:    &models.TeamMember{ID: 1},
			wantErr: false,
		},
		{
			name: "failed ge db",
			id:   1,
			mockFunc: func(input uint64) {
				key := models.KeyCacheTeamMemberDetail(input)
				srv.repo.On("GetCache", mock.Anything, key, &models.TeamMember{
					ID: 0,
				}).Return(false).Once()

				srv.repo.On("GetDetail", mock.Anything, dto.TeamMemberDetailReq{
					ID: input,
				}).Return(nil, errors.New("invalid")).Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "not found",
			id:   1,
			mockFunc: func(input uint64) {
				key := models.KeyCacheTeamMemberDetail(input)
				srv.repo.On("GetCache", mock.Anything, key, &models.TeamMember{
					ID: 0,
				}).Return(false).Once()

				srv.repo.On("GetDetail", mock.Anything, dto.TeamMemberDetailReq{
					ID: input,
				}).Return(nil, nil).Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			id:   1,
			mockFunc: func(input uint64) {
				key := models.KeyCacheTeamMemberDetail(input)
				srv.repo.On("GetCache", mock.Anything, key, &models.TeamMember{
					ID: 0,
				}).Return(false).Once()

				srv.repo.On("GetDetail", mock.Anything, dto.TeamMemberDetailReq{
					ID: input,
				}).Return(&models.TeamMember{ID: input}, nil).Once()

				srv.repo.On("CreateCache", mock.Anything, key, &models.TeamMember{ID: input}, time.Minute).Return().Once()

			},
			want:    &models.TeamMember{ID: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		srv.T().Run(tt.name, func(t *testing.T) {
			if tt.mockFunc != nil {
				tt.mockFunc(tt.id)
			}

			got, err := srv.service.GetByID(srv.ctx, tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("TeamMemberSrv.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TeamMemberSrv.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (srv *TeamMemberServiceTestSuite) TestTeamMemberSrv_Create() {
	params := dto.TeamMemberCreateReq{
		Name:           "adam",
		UsernameGithub: "adamnasrudin03",
		Email:          "adam@example.com",
	}
	tests := []struct {
		name     string
		req      dto.TeamMemberCreateReq
		mockFunc func(input dto.TeamMemberCreateReq)
		want     *models.TeamMember
		wantErr  bool
	}{
		{
			name: "duplicate",
			req:  params,
			mockFunc: func(input dto.TeamMemberCreateReq) {
				srv.repo.On("GetDetail", mock.Anything, dto.TeamMemberDetailReq{
					CustomColumn: "id",
					Email:        input.Email,
				}).Return(nil, nil).Once()

				srv.repo.On("GetDetail", mock.Anything, dto.TeamMemberDetailReq{
					CustomColumn:   "id",
					UsernameGithub: input.UsernameGithub,
				}).Return(&models.TeamMember{ID: 1}, nil).Once()

			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed create record",
			req:  params,
			mockFunc: func(input dto.TeamMemberCreateReq) {
				srv.repo.On("GetDetail", mock.Anything, dto.TeamMemberDetailReq{
					CustomColumn: "id",
					Email:        input.Email,
				}).Return(nil, nil).Once()

				srv.repo.On("GetDetail", mock.Anything, dto.TeamMemberDetailReq{
					CustomColumn:   "id",
					UsernameGithub: input.UsernameGithub,
				}).Return(nil, nil).Once()

				record := &models.TeamMember{
					Name:           input.Name,
					Email:          input.Email,
					UsernameGithub: input.UsernameGithub,
				}
				srv.repo.On("Create", mock.Anything, record).Return(nil, errors.New("invalid")).Once()

			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			req:  params,
			mockFunc: func(input dto.TeamMemberCreateReq) {
				srv.repo.On("GetDetail", mock.Anything, dto.TeamMemberDetailReq{
					CustomColumn: "id",
					Email:        input.Email,
				}).Return(nil, nil).Once()

				srv.repo.On("GetDetail", mock.Anything, dto.TeamMemberDetailReq{
					CustomColumn:   "id",
					UsernameGithub: input.UsernameGithub,
				}).Return(nil, nil).Once()

				record := &models.TeamMember{
					Name:           input.Name,
					Email:          input.Email,
					UsernameGithub: input.UsernameGithub,
				}
				srv.repo.On("Create", mock.Anything, record).Return(&models.TeamMember{
					ID:             101,
					Name:           input.Name,
					Email:          input.Email,
					UsernameGithub: input.UsernameGithub,
				}, nil).Once()

			},
			want: &models.TeamMember{
				ID:             101,
				Name:           params.Name,
				Email:          params.Email,
				UsernameGithub: params.UsernameGithub,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		srv.T().Run(tt.name, func(t *testing.T) {
			if tt.mockFunc != nil {
				tt.mockFunc(tt.req)
			}

			got, err := srv.service.Create(srv.ctx, tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TeamMemberSrv.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TeamMemberSrv.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
