package store

type User struct {
	ID           int64  `db:"id"`
	Login        string `db:"login"`
	Email        string `db:"email"`
	PasswordHash string `db:"password_hash"`
	Phone        string `db:"phone"`
}
