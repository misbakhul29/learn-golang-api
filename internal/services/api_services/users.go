package api_services

import (
	"belajargo/internal/dto"
	"belajargo/internal/models"
	"belajargo/internal/repositories"
	"belajargo/internal/services/redis"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type UserServiceImpl struct {
	repo repositories.UserRepository
	rdb  *redis.Store
}

func NewUserService(repo repositories.UserRepository, rdb *redis.Store) *UserServiceImpl {
	return &UserServiceImpl{repo: repo, rdb: rdb}
}

func (s *UserServiceImpl) GetAll() ([]models.Users, error) {
	var users []models.Users
	err := s.rdb.GetObject(redis.KEY_ALL_USER, &users)
	if err == nil {
		return users, nil
	}

	users, err = s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	_ = s.rdb.SetObject(redis.KEY_ALL_USER, users, 15*time.Minute)
	return users, nil
}

func (s *UserServiceImpl) GetByID(id uuid.UUID) (*models.Users, error) {
	var user models.Users
	key := fmt.Sprintf(redis.KEY_USER_BY_ID, id.String())
	err := s.rdb.GetObject(key, &user)
	if err == nil {
		return &user, nil
	}

	u, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	_ = s.rdb.SetObject(key, u, time.Hour)
	return u, nil
}

func (s *UserServiceImpl) Create(input *dto.InputUser) (*models.Users, error) {
	u, err := s.repo.Create(input)
	if err != nil {
		return nil, err
	}

	_ = s.rdb.Del(redis.KEY_ALL_USER)
	return u, nil
}

func (s *UserServiceImpl) Update(id uuid.UUID, input *dto.InputUser) (*models.Users, error) {
	u, err := s.repo.Update(id, input)
	if err != nil {
		return nil, err
	}

	_ = s.rdb.Del(fmt.Sprintf(redis.KEY_USER_BY_ID, id.String()))
	_ = s.rdb.Del(redis.KEY_ALL_USER)
	return u, nil
}

func (s *UserServiceImpl) Delete(id uuid.UUID) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	_ = s.rdb.Del(fmt.Sprintf(redis.KEY_USER_BY_ID, id.String()))
	_ = s.rdb.Del(redis.KEY_ALL_USER)
	return nil
}
