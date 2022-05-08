package domain

type User struct {
	Id             uint   `db:"id"`
	Fullname       string `db:"fullname"`
	Mobile         string `db:"mobile"`
	ActiveCode     uint   `db:"active_code"`
	ExpireDateCode string `db:"expire_date_code"`
	Active         uint   `db:"active"`
	CreatedAt      string `db:"created_at"`
	TypeUser       uint   `db:"type_user"`
}
