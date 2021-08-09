package rdto

type University struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	FullName string `db:"full_name"`
}
