package dto

type UniversityDirections struct {
	UniversityID       uint        `json:"university_id"`
	UniversityName     string      `json:"university_name"`
	UniversityFullName string      `json:"university_full_name"`
	Directions         []Direction `json:"directions"`
}
