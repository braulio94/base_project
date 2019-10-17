package models

type Project struct {
	ID             int64  `json:"id"`
	ProjectName    string `json:"project_name"`
	Description    string `json:"description"`
	LastAccessed   string `json:"last_accessed"`
	LastAccessedBy string `json:"last_accessed_by"`
	Status         Status `json:"status"`
}
