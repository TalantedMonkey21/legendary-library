package usecase

import (
	"context"
	"errors"
	"strings"

	entity "github.com/TalantedMonkey21/GoLectures/internal/entity"
	"github.com/TalantedMonkey21/GoLectures/internal/repository"
)

// требования для слоя репозитория
type NoteRepositorier interface {
	Create(ctx context.Context, note entity.Note) (entity.Note, error)
	GetByID(ctx context.Context, id uint) (entity.Note, error)
	Update(ctx context.Context, note entity.Note) (entity.Note, error)
	Delete(ctx context.Context, id uint) (repository.NoteModel, error)
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
		return errors.New("слишком мало символов")
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

func (uc *NoteUseCase) GetByID(ctx context.Context, id uint) (entity.Note, error) {
	if id == 0 {
		return entity.Note{}, errors.New("invalid id")
	}
	return uc.repo.GetByID(ctx, id)
}

func (uc *NoteUseCase) Update(ctx context.Context, note entity.Note) (entity.Note, error) {
	content := strings.TrimSpace(note.Content)
	if err := checkLen(content); err != nil {
		return entity.Note{}, err
	}

	return uc.repo.Update(ctx, note)
}

func (uc *NoteUseCase) Delete(ctx context.Context, id uint) (repository.NoteModel, error) {
	var model repository.NoteModel
	if id == 0 {
		return model, errors.New("invalid id")
	}
	return uc.repo.Delete(ctx, id)
}

// TODO: сделать функцию
// GetByID
// Update
// Delete
