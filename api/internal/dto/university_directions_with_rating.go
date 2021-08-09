package dto

import "sort"

type UniversityDirectionsWithRating struct {
	UniversityID       uint                  `json:"university_id"`
	UniversityName     string                `json:"university_name"`
	UniversityFullName string                `json:"university_full_name"`
	Directions         []DirectionWithRating `json:"directions"`
}

func (u *UniversityDirectionsWithRating) SortDirections() {
	sort.SliceStable(u.Directions, func(i, j int) bool {
		return u.Directions[i].ID < u.Directions[j].ID
	})
}
