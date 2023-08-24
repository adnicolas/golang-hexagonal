package pg

const (
	sqlUserTable = "public.user"
)

type SqlUser struct {
	ID       string `db:"uuid"`
	Name     string `db:"name"`
	Surname  string `db:"surname"`
	Password string `db:"password"`
	Email    string `db:"email"`
}
