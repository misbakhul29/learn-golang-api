package repositories

import (
	"belajargo/internal/dto"
	"belajargo/internal/models"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostRepository interface {
	FindAll() ([]models.Posts, error)
	FindAllWithUsers() ([]models.Posts, error)
	FindByID(id uuid.UUID) (*models.Posts, error)
	Create(input *dto.InputPosts) (*models.Posts, error)
	Update(id uuid.UUID, input *dto.InputPosts) (*models.Posts, error)
	Delete(id uuid.UUID) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) FindAll() ([]models.Posts, error) {
	var posts []models.Posts
	if err := r.db.Find(&posts).Error; err != nil {
		return nil, huma.NewError(http.StatusInternalServerError, "Failed to get all posts", err)
	}
	return posts, nil
}

func (r *postRepository) FindAllWithUsers() ([]models.Posts, error) {
	var posts []models.Posts
	if err := r.db.Preload("User").Find(&posts).Error; err != nil {
		return nil, huma.NewError(http.StatusInternalServerError, "Failed to get all posts", err)
	}
	return posts, nil
}

func (r *postRepository) FindByID(id uuid.UUID) (*models.Posts, error) {
	var post models.Posts
	if err := r.db.Where("id = ?", id).First(&post).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, huma.NewError(http.StatusNotFound, "Post not found", err)
		}
		return nil, huma.NewError(http.StatusInternalServerError, "Failed to get post", err)
	}
	return &post, nil
}

func (r *postRepository) Create(input *dto.InputPosts) (*models.Posts, error) {
	title := ""
	if input.Title != nil {
		title = *input.Title
	}
	content := ""
	if input.Content != nil {
		content = *input.Content
	}
	var userID uuid.UUID
	if input.UserID != nil {
		userID, _ = uuid.Parse(*input.UserID)
	}

	post := models.Posts{
		Title:   title,
		Content: content,
		UserID:  userID,
	}
	if err := r.db.Create(&post).Error; err != nil {
		return nil, huma.NewError(http.StatusInternalServerError, "Failed to create post", err)
	}
	return &post, nil
}

func (r *postRepository) Update(id uuid.UUID, input *dto.InputPosts) (*models.Posts, error) {
	var post models.Posts
	if err := r.db.First(&post, "id = ?", id).Error; err != nil {
		return nil, huma.NewError(http.StatusNotFound, "Post not found", err)
	}
	if err := r.db.Model(&post).Updates(input).Error; err != nil {
		return nil, huma.NewError(http.StatusInternalServerError, "Failed to update post", err)
	}
	return &post, nil
}

func (r *postRepository) Delete(id uuid.UUID) error {
	var post models.Posts
	if err := r.db.Where("id = ?", id).Delete(&post).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return huma.NewError(http.StatusNotFound, "Post not found", err)
		}
		return huma.NewError(http.StatusInternalServerError, "Failed to delete post", err)
	}
	return nil
}
