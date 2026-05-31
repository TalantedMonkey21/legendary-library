package repository

import (
	"context"
	"errors"
	"time"

	"github.com/TalantedMonkey21/GoLectures/internal/apperrors"
	entity "github.com/TalantedMonkey21/GoLectures/internal/entity"
	"gorm.io/gorm"
)

// описание сущности для хранения в репозитории
type NoteModel struct {
	ID        uint   `gorm:"primaryKey"`
	UserID	  uint   `gorm:"not null"`
	Content   string `gorm:"type:text; not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// даем название таблицы
func (NoteModel) TableName() string {
	return "notes"
}

// слой репозиторий
type NoteRepo struct {
	db *gorm.DB
}

func toEntityNote(m NoteModel) entity.Note {
	return entity.Note{
		ID:        m.ID,
		UserID:    m.UserID,
		Content:   m.Content,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func toModelNote(n entity.Note) NoteModel {
	return NoteModel{
		ID:  	   n.ID,
		UserID:    n.UserID,
		Content:   n.Content,
		CreatedAt: n.CreatedAt,
		UpdatedAt: n.UpdatedAt,
	}
}

// конструктор
func NewNoteRepo(db *gorm.DB) *NoteRepo {
	return &NoteRepo{db: db}
}

func (n *NoteRepo) Create(ctx context.Context, note entity.Note) (entity.Note, error) {
	model := toModelNote(note)
	if err := n.db.WithContext(ctx).Create(&model).Error; err != nil {
		return entity.Note{}, err
	}

	return toEntityNote(model), nil
}

func (n *NoteRepo) GetByID(ctx context.Context, id uint, userId uint) (entity.Note, error) {
	var model NoteModel
	if err := n.db.WithContext(ctx).Where("id = ? and user_id = ?", id, userId).First(&model).Error; err != nil {
		return entity.Note{}, err
	}

	return toEntityNote(model), nil
}

func (n *NoteRepo) Update(ctx context.Context, note entity.Note) (entity.Note, error) {
	var model NoteModel
	if err := n.db.WithContext(ctx).Where("id = ? and user_id = ?", note.ID, note.UserID).First(&model, note.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Note{}, apperrors.ErrNoteNotFound
		}
		return entity.Note{}, err
	}
	model.Content = note.Content
	if err := n.db.WithContext(ctx).Save(&model).Error; err != nil {
		return entity.Note{}, err
	}
	return toEntityNote(model), nil
}

func (n *NoteRepo) Delete(ctx context.Context, id uint, userId uint) error {
	del := n.db.WithContext(ctx).Where("id = ? and user_id = ?", id, userId).Delete(&NoteModel{})
	if del.Error != nil {
		return del.Error
	}
	if del.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
