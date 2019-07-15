package models

type Metadata struct {
	Total       uint `json:"total"`
	CurrentPage uint `json:"current_page"`
	LastPage    uint `json:"last_page"`
	PerPage     uint `json:"per_page"`
}