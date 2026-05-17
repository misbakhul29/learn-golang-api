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

type PostServiceImpl struct {
	repo repositories.PostRepository
	rdb  *redis.Store
}

func NewPostService(repo repositories.PostRepository, rdb *redis.Store) *PostServiceImpl {
	return &PostServiceImpl{repo: repo, rdb: rdb}
}

func (s *PostServiceImpl) GetAll(includeUsers bool) ([]models.Posts, error) {
	var posts []models.Posts
	var cacheKey string
	if includeUsers {
		cacheKey = redis.KEY_ALL_POST_WITH_USERS
	} else {
		cacheKey = redis.KEY_ALL_POST
	}

	err := s.rdb.GetObject(cacheKey, &posts)
	if err == nil {
		return posts, nil
	}

	if includeUsers {
		posts, err = s.repo.FindAllWithUsers()
		if err != nil {
			return nil, err
		}
		_ = s.rdb.SetObject(cacheKey, posts, 10*time.Minute)
		return posts, nil
	}

	posts, err = s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	_ = s.rdb.SetObject(cacheKey, posts, 10*time.Minute)
	return posts, nil
}

func (s *PostServiceImpl) GetByID(id uuid.UUID) (*models.Posts, error) {
	var post models.Posts
	key := fmt.Sprintf(redis.KEY_POST_BY_ID, id.String())
	err := s.rdb.GetObject(key, &post)
	if err == nil {
		return &post, nil
	}

	postVal, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	_ = s.rdb.SetObject(key, postVal, 10*time.Minute)
	return postVal, nil
}

func (s *PostServiceImpl) Create(input *dto.InputPosts) (*models.Posts, error) {
	p, err := s.repo.Create(input)
	if err != nil {
		return nil, err
	}

	_ = s.rdb.Del(redis.KEY_ALL_POST)
	_ = s.rdb.Del(redis.KEY_ALL_POST_WITH_USERS)
	return p, nil
}

func (s *PostServiceImpl) Update(id uuid.UUID, input *dto.InputPosts) (*models.Posts, error) {
	p, err := s.repo.Update(id, input)
	if err != nil {
		return nil, err
	}

	_ = s.rdb.Del(fmt.Sprintf(redis.KEY_POST_BY_ID, id.String()))
	_ = s.rdb.Del(redis.KEY_ALL_POST)
	_ = s.rdb.Del(redis.KEY_ALL_POST_WITH_USERS)
	return p, nil
}

func (s *PostServiceImpl) Delete(id uuid.UUID) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	_ = s.rdb.Del(fmt.Sprintf(redis.KEY_POST_BY_ID, id.String()))
	_ = s.rdb.Del(redis.KEY_ALL_POST)
	_ = s.rdb.Del(redis.KEY_ALL_POST_WITH_USERS)
	return nil
}
