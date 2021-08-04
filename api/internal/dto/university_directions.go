package dto

type UniversityDirections struct {
	UniversityID   uint        `json:"university_id"`
	UniversityName string      `json:"university_name"`
	Directions     []Direction `json:"directions"`
}
