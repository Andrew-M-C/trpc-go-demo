package account

import (
	"strconv"

	"github.com/Andrew-M-C/trpc-go-demo/app/user/entity"
)

type userAccountItem struct {
	ID           int64  `db:"id"`
	Username     string `db:"username"`
	PasswordHash string `db:"password_hash"`
}

// TableName 返回表名
func (userAccountItem) TableName() string {
	return "t_user_account"
}

func (u userAccountItem) toEntity() *entity.Account {
	return &entity.Account{
		ID:           strconv.FormatInt(u.ID, 36),
		Username:     u.Username,
		PasswordHash: u.PasswordHash,
	}
}
