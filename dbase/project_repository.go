package dbase

import (
	"github.com/braulio94/base_project/models"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type ProjectRepository interface {
	Fetch(limit int64) ([]*models.Project, error)
	GetById(id uuid.UUID) (*models.Project, error)
	Update(project *models.Project) error
	Store(project *models.Project) error
	Delete(project *models.Project) error
}

type repo struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) ProjectRepository {
	return &repo{DB: db}
}

func (r *repo) Fetch(limit int64) ([]*models.Project, error) {
	return nil, nil
}

func (r *repo) Store(project *models.Project) error {
	p := &project
	return r.DB.Create(p).Error
}

func (r *repo) GetById(id uuid.UUID) (*models.Project, error) {
	project := new(models.Project)
	err := r.DB.Where("id = ?", id).Find(project).Error
	return project, err
}

func (r *repo) Update(project *models.Project) error {
	return nil
}

func (r *repo) Delete(*models.Project) error {
	return nil
}
