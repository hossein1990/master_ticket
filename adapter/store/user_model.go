package store

import "github.com/hossein1990/master_ticket/domain"

type User struct {
	Id             uint   `json:"id" gorm:"primary_key"`
	Fullname       string `json:"fullname"`
	Mobile         string `json:"mobile"`
	ActiveCode     uint   `json:"active_code"`
	ExpireDateCode string `json:"expire_date_code"`
	Active         uint   `json:"active"`
	CreatedAt      string `json:"created_at"`
	TypeUser       uint   `json:"type_user"`
}

func mapFromUserEntity(user domain.User) User {
	return User{
		Id:             user.Id,
		Fullname:       user.Fullname,
		Mobile:         user.Mobile,
		ActiveCode:     user.ActiveCode,
		ExpireDateCode: user.ExpireDateCode,
		Active:         user.Active,
		CreatedAt:      user.CreatedAt,
		TypeUser:       user.TypeUser,
	}
}

func mapToUserEntity(user User) domain.User {
	return domain.User{
		Id:             user.Id,
		Fullname:       user.Fullname,
		Mobile:         user.Mobile,
		ActiveCode:     user.ActiveCode,
		ExpireDateCode: user.ExpireDateCode,
		Active:         user.Active,
		CreatedAt:      user.CreatedAt,
		TypeUser:       user.TypeUser,
	}
}
