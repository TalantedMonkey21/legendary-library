package usecase

import (
	"context"
	"strings"

	"github.com/TalantedMonkey21/GoLectures/internal/apperrors"
	"github.com/TalantedMonkey21/GoLectures/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

//TODO: interface

type UserRepositorier interface {
	Create(ctx context.Context, user entity.User) (entity.User, error)
	GetByEmail(ctx context.Context, email string) (entity.User, error)
	Update(ctx context.Context, user entity.User) (entity.User, error)
	Delete(ctx context.Context, id uint) error
}

type UserUseCase struct {
	repo UserRepositorier
}

func (uc *UserUseCase) Register(ctx context.Context, name, email, password string) (entity.User, error) {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(strings.ToLower(email))
	password = strings.TrimSpace(password)
	//TODO: По настроению отдельный файл ошибок
	if len(name) < 2 {
		return entity.User{}, apperrors.ErrTooShort
	}
	if !strings.Contains(email, "@") {
		return entity.User{}, apperrors.ErrInvalidEmail
	}
	if len(password) < 6 {
		return entity.User{}, apperrors.ErrWeakPassword
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return entity.User{}, err
	}
	user := entity.User{
		Name: name,
		Email: email,
		PasswordHash: string(hash), 
	}
	return user, nil
}