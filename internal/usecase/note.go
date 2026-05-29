package usecase

import (
	"context"
	"strings"

	"github.com/TalantedMonkey21/GoLectures/internal/apperrors"
	entity "github.com/TalantedMonkey21/GoLectures/internal/entity"
)

// требования для слоя репозитория
type NoteRepositorier interface {
	Create(ctx context.Context, note entity.Note) (entity.Note, error)
	GetByID(ctx context.Context, id uint, userId uint) (entity.Note, error)
	Update(ctx context.Context, note entity.Note) (entity.Note, error)
	Delete(ctx context.Context, id uint, userId uint) error
}

// сам слой бизнес логики
type NoteUseCase struct {
	repo NoteRepositorier
}

// функция конструктор возвращающая слой бизнес логики (принимает в себя слой репоззитория)
func New(repo NoteRepositorier) *NoteUseCase {
	return &NoteUseCase{repo: repo}
}

func checkLen(content string) error {
	if len(content) < 5 {
		return apperrors.ErrTooShort
	}
	return nil
}

func (uc *NoteUseCase) Create(ctx context.Context, content string) (entity.Note, error) {
	content = strings.TrimSpace(content)
	if err := checkLen(content); err != nil {
		return entity.Note{}, err
	}

	note := entity.Note{Content: content}

	return uc.repo.Create(ctx, note)
}

func (uc *NoteUseCase) GetByID(ctx context.Context, id uint, userId uint) (entity.Note, error) {
	if id == 0 {
		return entity.Note{}, apperrors.ErrInvalidID
	}
	return uc.repo.GetByID(ctx, id, userId)
}

func (uc *NoteUseCase) Update(ctx context.Context, note entity.Note) (entity.Note, error) {
	content := strings.TrimSpace(note.Content)
	if err := checkLen(content); err != nil {
		return entity.Note{}, err
	}
	note.Content = content
	return uc.repo.Update(ctx, note)
}

func (uc *NoteUseCase) Delete(ctx context.Context, id uint, userId uint) error {
	if id == 0 {
		return apperrors.ErrInvalidID
	}
	return uc.repo.Delete(ctx, id, userId)
}


