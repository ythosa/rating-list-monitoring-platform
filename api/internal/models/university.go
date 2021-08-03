package models

type University struct {
	ID                int    `json:"id" db:"id"`
	Name              string `json:"name" db:"name"`
	FullName          string `json:"full_name" db:"full_name"`
	DirectionsPageURL string `json:"directions_page_url" db:"directions_page_url"`
}
