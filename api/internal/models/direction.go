package models

type Direction struct {
	ID           uint   `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	URL          string `json:"url" db:"url"`
	UniversityID uint   `json:"university_id" db:"university_id"`
}
