package dto

type UniversityDirectionsWithRating struct {
	UniversityID       uint                  `json:"university_id"`
	UniversityName     string                `json:"university_name"`
	UniversityFullName string                `json:"university_full_name"`
	Directions         []DirectionWithRating `json:"directions"`
}
