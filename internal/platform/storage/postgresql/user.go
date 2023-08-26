package pg

const (
	sqlUserTable = "public.user"
)

// DTO that facilitates the mapping with the DB
// Data mapper pattern
type sqlUser struct {
	ID       string `db:"uuid" fieldtag:"select"`
	Name     string `db:"name" fieldtag:"select"`
	Surname  string `db:"surname" fieldtag:"select"`
	Password string `db:"password"`
	Email    string `db:"email" fieldtag:"select"`
}
