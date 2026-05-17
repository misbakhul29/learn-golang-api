package repositories

import (
	"belajargo/internal/dto"
	"belajargo/internal/models"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]models.Users, error)
	FindByID(id uuid.UUID) (*models.Users, error)
	Create(input *dto.InputUser) (*models.Users, error)
	Update(id uuid.UUID, input *dto.InputUser) (*models.Users, error)
	Delete(id uuid.UUID) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindAll() ([]models.Users, error) {
	var users []models.Users
	if err := r.db.Find(&users).Error; err != nil {
		return nil, huma.NewError(http.StatusInternalServerError, "Failed to get all users", err)
	}
	return users, nil
}

func (r *userRepository) FindByID(id uuid.UUID) (*models.Users, error) {
	var user models.Users
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, huma.NewError(http.StatusNotFound, "User not found", err)
		}
		return nil, huma.NewError(http.StatusInternalServerError, "Failed to get user", err)
	}
	return &user, nil
}

func (r *userRepository) Create(input *dto.InputUser) (*models.Users, error) {
	name := ""
	if input.Name != nil {
		name = *input.Name
	}
	email := ""
	if input.Email != nil {
		email = *input.Email
	}

	user := models.Users{
		Name:  name,
		Email: email,
	}
	if err := r.db.Create(&user).Error; err != nil {
		return nil, huma.NewError(http.StatusInternalServerError, "Failed to create user", err)
	}
	return &user, nil
}

func (r *userRepository) Update(id uuid.UUID, input *dto.InputUser) (*models.Users, error) {
	var user models.Users
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, huma.NewError(http.StatusNotFound, "User not found", err)
	}
	if err := r.db.Model(&user).Updates(input).Error; err != nil {
		return nil, huma.NewError(http.StatusInternalServerError, "Failed to update user", err)
	}
	return &user, nil
}

func (r *userRepository) Delete(id uuid.UUID) error {
	var user models.Users
	if err := r.db.Where("id = ?", id).Delete(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return huma.NewError(http.StatusNotFound, "User not found", err)
		}
		return huma.NewError(http.StatusInternalServerError, "Failed to delete user", err)
	}
	return nil
}
