package domain

type User struct {
	Id             uint   `json:"id"`
	Fullname       string `json:"fullname"`
	Mobile         string `json:"mobile"`
	ActiveCode     uint   `json:"active_code"`
	ExpireDateCode string `json:"expire_date_code"`
	Active         uint   `json:"active"`
	CreatedAt      string `json:"created_at"`
	TypeUser       uint   `json:"type_user"`
}
