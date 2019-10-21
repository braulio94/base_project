package models

import (
	"github.com/satori/go.uuid"
)

type Project struct {
	ID             uuid.UUID `gorm:"primary_key" json:"id"`
	ProjectName    string    `gorm:"column:project_name" json:"project_name"`
	Description    string    `gorm:"column:description" json:"description"`
	LastAccessed   string    `gorm:"column:last_accessed" json:"last_accessed"`
	LastAccessedBy string    `gorm:"column:last_accessed_by" json:"last_accessed_by"`
	Status         string    `gorm:"column:status" json:"status"`
}

func (p *Project) TableName() string {
	return "project"
}
