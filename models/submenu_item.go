package models

type SubmenuItem struct {
	ID        int64   `json:"id"`
	Title     string  `json:"title"`
	Path      string  `json:"path"`
	SubMenu   SubMenu `json:"sub_menu"`
	IsVisible bool    `json:"is_visible"`
	IsEnabled bool    `json:"is_enabled"`
}
