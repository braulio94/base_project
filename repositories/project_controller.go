package repositories

import "github.com/braulio94/base_project/models"

type ProjectController interface {
	Fetch(limit int64) ([]*models.Project, error)
	GetById(id int64) (*models.Project, error)
	Update(project *models.Project) error
	Store(project *models.Project) error
	Delete(project *models.Project) error
}
