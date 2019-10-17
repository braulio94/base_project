package models

type SubMenu struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Path      string `json:"path"`
	Menu      Menu   `json:"menu"`
	IsVisible bool   `json:"is_visible"`
	IsEnabled bool   `json:"is_enabled"`
}
