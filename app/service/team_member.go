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
	Create(ctx context.Context, req dto.TeamMemberCreateReq) (*models.TeamMember, error)
	GetByID(ctx context.Context, id uint64) (*models.TeamMember, error)
	DeleteByID(ctx context.Context, id uint64) error
	Update(ctx context.Context, req dto.TeamMemberUpdateReq) error
	GetList(ctx context.Context, req dto.TeamMemberListReq) (*helpers.Pagination, error)
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

func (s TeamMemberSrv) Create(ctx context.Context, req dto.TeamMemberCreateReq) (*models.TeamMember, error) {
	var (
		opName = "TeamMemberService-Create"
		err    error
		resp   *models.TeamMember
	)

	req.Email = helpers.ToLower(req.Email)
	req.UsernameGithub = helpers.ToLower(req.UsernameGithub)

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

func (s TeamMemberSrv) Update(ctx context.Context, req dto.TeamMemberUpdateReq) error {
	var (
		opName = "TeamMemberService-Update"
		key    = models.KeyCacheTeamMemberDetail(req.ID)
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

	go s.Repo.DeleteCache(context.Background(), key)
	return nil
}

func (s TeamMemberSrv) GetList(ctx context.Context, req dto.TeamMemberListReq) (*helpers.Pagination, error) {
	var (
		opName = "TeamMemberService-GetList"
		err    error
		resp   *helpers.Pagination
	)
	err = req.Validate()
	if err != nil {
		return nil, err
	}

	data, err := s.Repo.GetList(ctx, req)
	if err != nil {
		s.Logger.Errorf("%s, failed get list: %v", opName, err)
		return nil, helpers.ErrDB()
	}

	totalRecords := len(data)
	resp = &helpers.Pagination{
		Data: data,
		Meta: helpers.Meta{
			Page:         req.Page,
			Limit:        req.Limit,
			TotalRecords: totalRecords,
		},
	}

	// total records in less than limit
	if totalRecords > 0 && totalRecords != req.Limit {
		return resp, nil
	}

	// get total data
	if totalRecords > 0 {
		req.CustomColumns = "id"
		req.IsNotDefaultQuery = true
		req.Offset = (req.Page - 1) * req.Limit
		req.Limit = models.DefaultLimitIsTotalDataTrue * req.Limit

		total, err := s.Repo.GetList(ctx, req)
		if err != nil {
			s.Logger.Errorf("%s, failed get total data: %v", opName, err)
			return nil, helpers.ErrDB()
		}
		totalRecords = len(total)
		resp.Meta.TotalRecords = totalRecords
	}

	return resp, nil
}
