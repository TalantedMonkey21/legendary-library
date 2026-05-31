package repository

import (
	"context"
	"time"

	"github.com/TalantedMonkey21/GoLectures/internal/entity"
	"gorm.io/gorm"
)

//TODO: Сделать ВСЕ!!! По аналогии с note.go
type UserModel struct {
	ID 			 uint   `gorm:"primaryKey"`
	Name 		 string `gorm:"not null; unique"`
	PasswordHash string	`gorm:"not null"`
	Email 		 string `gorm:"not null; unique"`
	CreatedAt 	 time.Time
	UpdatedAt 	 time.Time
}

func (UserModel) TableName() string {
	return "users"
}

type UserRepo struct {
	db *gorm.DB
}

func toEntityUser(m UserModel) entity.User {
	return entity.User{
		ID: 		  m.ID,
		Name: 		  m.Name,
		PasswordHash: m.PasswordHash,
		Email: 		  m.Email,
		CreatedAt: 	  m.CreatedAt,
		UpdatedAt: 	  m.UpdatedAt,
	}
}

func toModelUser(u entity.User) UserModel {
	return UserModel{
		ID: 		  u.ID,
		Name: 		  u.Name,
		PasswordHash: u.PasswordHash,
		Email: 		  u.Email,
		CreatedAt: 	  u.CreatedAt,
		UpdatedAt: 	  u.UpdatedAt,
	}
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) Create(ctx context.Context, user entity.User) (entity.User, error) {
	model := toModelUser(user)
	if err := u.db.WithContext(ctx).Create(&model).Error; err != nil {
		return entity.User{}, err
	}

	return toEntityUser(model), nil
}

func (u *UserRepo) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	var model UserModel
	if err := u.db.WithContext(ctx).Where("email = ?", email).First(&model).Error; err != nil {
		return entity.User{}, err
	}

	return toEntityUser(model), nil
}

func (u *UserRepo) Update(ctx context.Context, user entity.User) (entity.User, error) {
	var model UserModel
	
}