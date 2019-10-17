package database

import (
	"github.com/braulio94/base_project/models"
	"github.com/jinzhu/gorm"
)

type ProjectRepository struct {
	Conn *gorm.DB
}

func NewRepository(conn *gorm.DB) ProjectRepository {
	return ProjectRepository{conn}
}

func fetch(limit int64) ([]*models.Project, error) {

	return nil, nil
}
