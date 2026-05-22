package repository

import (
	"context"
	"time"

	entity "github.com/TalantedMonkey21/GoLectures/internal/entity"
	"gorm.io/gorm"
)

// описание сущности для хранения в репозитории
type NoteModel struct {
	ID        uint   `gorm:"primaryKey"`
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

func toEntity(m NoteModel) entity.Note {
	return entity.Note{
		ID:        m.ID,
		Content:   m.Content,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

// конструктор
func New(db *gorm.DB) *NoteRepo {
	return &NoteRepo{db: db}
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&NoteModel{})
}

func (n *NoteRepo) Create(ctx context.Context, note entity.Note) (entity.Note, error) {
	model := NoteModel{
		Content: note.Content,
	}
	if err := n.db.WithContext(ctx).Create(&model).Error; err != nil {
		return entity.Note{}, err
	}

	en := entity.Note{
		ID:        model.ID,
		Content:   model.Content,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}

	return en, nil
}

// TODO
// Update
// Delete
func (n *NoteRepo) GetByID(ctx context.Context, id uint) (entity.Note, error) {
	var model NoteModel
	if err := n.db.WithContext(ctx).First(&model, id).Error; err != nil {
		return toEntity(model), err
	}

	return toEntity(model), nil
}

func (n *NoteRepo) Update(ctx context.Context, note entity.Note) (entity.Note, error) {
	model := NoteModel{
		ID:      note.ID,
		Content: note.Content,
	}
	if err := n.db.WithContext(ctx).Save(&model).Error; err != nil {
		return toEntity(model), err
	}

	return toEntity(model), nil
}

func (n *NoteRepo) Delete(ctx context.Context, id uint) (NoteModel, error) {
	var model NoteModel
	if err := n.db.WithContext(ctx).Delete(&model, id).Error; err != nil {
		return model, err
	}

	return model, nil
}
