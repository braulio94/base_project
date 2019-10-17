package models

type Menu struct {
	ID        int64   `json:"id"`
	Title     string  `json:"title"`
	Path      string  `json:"path"`
	Project   Project `json:"project"`
	IsVisible bool    `json:"is_visible"`
	IsEnabled bool    `json:"is_enabled"`
}
