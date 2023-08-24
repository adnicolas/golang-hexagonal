package pg

const (
	sqlUserTable = "public.user"
)

type SqlUser struct {
	// TODO: Usar tipo uuid para robustecer validaciones
	UUID     string/*uuid.UUID*/ `db:"uuid"`
	Name     string `db:"name"`
	Surname  string `db:"surname"`
	Password string `db:"password"`
	Email    string `db:"email"`
}
