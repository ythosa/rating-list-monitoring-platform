package dto

import "sort"

type UniversityDirections struct {
	UniversityID       uint        `json:"university_id"`
	UniversityName     string      `json:"university_name"`
	UniversityFullName string      `json:"university_full_name"`
	Directions         []Direction `json:"directions"`
}

func (u *UniversityDirections) SortDirections() {
	sort.SliceStable(u.Directions, func(i, j int) bool {
		return u.Directions[i].ID < u.Directions[j].ID
	})
}
