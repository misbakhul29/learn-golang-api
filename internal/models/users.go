package models

import (
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;not null" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"not null;unique" json:"email"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type InputUser struct {
	Name  *string `json:"name,omitempty" doc:"Name of the user"`
	Email *string `json:"email,omitempty" doc:"Email of the user"`
}

func GetAllUsers(db *gorm.DB) ([]Users, error) {
	users := []Users{}
	if err := db.Find(&users).Error; err != nil {
		return nil, huma.NewError(http.StatusInternalServerError, "Failed to get all users", err)
	}
	return users, nil
}

func GetUserById(db *gorm.DB, id uuid.UUID) (*Users, error) {
	users := Users{}
	if err := db.Where("id = ?", id).First(&users).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, huma.NewError(http.StatusNotFound, "User not found", err)
		}
		return nil, huma.NewError(http.StatusInternalServerError, "Failed to get user", err)
	}
	return &users, nil
}

func CreateUser(db *gorm.DB, input *InputUser) (*Users, error) {
	name := ""
	if input.Name != nil {
		name = *input.Name
	}
	email := ""
	if input.Email != nil {
		email = *input.Email
	}

	user := Users{
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := db.Create(&user).Error; err != nil {
		return nil, huma.NewError(http.StatusInternalServerError, "Failed to create user", err)
	}
	return &user, nil
}

func UpdateUserById(db *gorm.DB, input *InputUser, id uuid.UUID) (*Users, error) {
	user := Users{}
	if err := db.First(&user, "id = ?", id).Error; err != nil {
		return nil, huma.NewError(http.StatusNotFound, "User not found", err)
	}
	if err := db.Model(&user).Updates(input).Error; err != nil {
		return nil, huma.NewError(http.StatusInternalServerError, "Failed to update user", err)
	}
	return &user, nil
}

func DeleteUserById(db *gorm.DB, id uuid.UUID) (*Users, error) {
	users := Users{}
	if err := db.Where("id = ?", id).Delete(&users).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, huma.NewError(http.StatusNotFound, "User not found", err)
		}
		if err == gorm.ErrInvalidData {
			return nil, huma.NewError(http.StatusBadRequest, "Invalid data", err)
		}
		return nil, huma.NewError(http.StatusInternalServerError, "Failed to delete user", err)
	}
	return nil, nil
}
