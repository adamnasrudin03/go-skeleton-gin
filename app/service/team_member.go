package service

import (
	"context"
	"time"

	"github.com/adamnasrudin03/go-skeleton-gin/app/configs"
	"github.com/adamnasrudin03/go-skeleton-gin/app/dto"
	"github.com/adamnasrudin03/go-skeleton-gin/app/models"
	"github.com/adamnasrudin03/go-skeleton-gin/app/repository"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
	"github.com/sirupsen/logrus"
)

type TeamMemberService interface {
	Create(ctx context.Context, req *dto.TeamMemberCreateReq) (*models.TeamMember, error)
	GetByID(ctx context.Context, id uint64) (*models.TeamMember, error)
	DeleteByID(ctx context.Context, id uint64) error
	Update(ctx context.Context, req *dto.TeamMemberUpdateReq) error
}

type TeamMemberSrv struct {
	Repo   repository.TeamMemberRepository
	Cfg    *configs.Configs
	Logger *logrus.Logger
}

func NewTeamMemberService(
	tmRepo repository.TeamMemberRepository,
	cfg *configs.Configs,
	logger *logrus.Logger,
) TeamMemberService {
	return TeamMemberSrv{
		Repo:   tmRepo,
		Cfg:    cfg,
		Logger: logger,
	}
}

func (s TeamMemberSrv) Create(ctx context.Context, req *dto.TeamMemberCreateReq) (*models.TeamMember, error) {
	var (
		opName = "TeamMemberService-Create"
		err    error
		resp   *models.TeamMember
	)

	err = s.checkDuplicate(ctx, dto.TeamMemberDetailReq{
		Email:          req.Email,
		UsernameGithub: req.UsernameGithub,
	})
	if err != nil {
		return nil, err
	}

	resp, err = s.Repo.Create(ctx, &models.TeamMember{
		Name:           req.Name,
		Email:          req.Email,
		UsernameGithub: req.UsernameGithub,
	})
	if err != nil {
		s.Logger.Errorf("%s, failed create db: %v", opName, err)
		return nil, helpers.ErrCreatedDB()
	}

	return resp, nil
}

func (s TeamMemberSrv) GetByID(ctx context.Context, id uint64) (*models.TeamMember, error) {
	var (
		opName = "TeamMemberService-GetByID"
		err    error
		resp   models.TeamMember
		key    = models.KeyCacheTeamMemberDetail(id)
	)

	ok := s.Repo.GetCache(ctx, key, &resp)
	if ok && resp.ID > 0 {
		return &resp, nil
	}

	detail, err := s.Repo.GetDetail(ctx, dto.TeamMemberDetailReq{
		ID: id,
	})
	if err != nil {
		s.Logger.Errorf("%s, failed get detail: %v", opName, err)
		return nil, helpers.ErrDB()
	}

	isExist := detail != nil && detail.ID > 0
	if !isExist {
		return nil, helpers.ErrNotFound()
	}

	go s.Repo.CreateCache(context.Background(), key, detail, time.Minute)

	return detail, nil
}

func (s TeamMemberSrv) DeleteByID(ctx context.Context, id uint64) error {
	var (
		opName = "TeamMemberService-DeleteByID"
		key    = models.KeyCacheTeamMemberDetail(id)
		err    error
	)

	_, err = s.GetByID(ctx, id)
	if err != nil {
		s.Logger.Errorf("%s, failed get detail: %v", opName, err)
		return err
	}

	err = s.Repo.Delete(ctx, &models.TeamMember{ID: id})
	if err != nil {
		s.Logger.Errorf("%s, failed delete db: %v", opName, err)
		return helpers.ErrDB()
	}

	go s.Repo.DeleteCache(context.Background(), key)

	return nil
}

func (s TeamMemberSrv) Update(ctx context.Context, req *dto.TeamMemberUpdateReq) error {
	var (
		opName = "TeamMemberService-Update"
		err    error
	)

	_, err = s.GetByID(ctx, req.ID)
	if err != nil {
		s.Logger.Errorf("%s, failed get detail: %v", opName, err)
		return err
	}

	err = s.checkDuplicate(ctx, dto.TeamMemberDetailReq{
		Email:          req.Email,
		UsernameGithub: req.UsernameGithub,
		NotID:          req.ID,
	})
	if err != nil {
		return err
	}

	err = s.Repo.Update(ctx, &models.TeamMember{
		ID:             req.ID,
		Name:           req.Name,
		Email:          req.Email,
		UsernameGithub: req.UsernameGithub,
	})
	if err != nil {
		s.Logger.Errorf("%s, failed update db: %v", opName, err)
		return helpers.ErrUpdatedDB()
	}

	return nil
}
