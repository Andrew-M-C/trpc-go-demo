package account

import (
	"github.com/Andrew-M-C/go.util/ids/base36"
	"github.com/Andrew-M-C/trpc-go-demo/app/user/entity"
)

const userAccountTableName = "t_trpc_demo_user_account"

type userAccountItem struct {
	ID           uint32 `db:"id"`
	Username     string `db:"username"`
	PasswordHash string `db:"password_hash"`
}

func (u userAccountItem) toEntity() *entity.Account {
	return &entity.Account{
		ID:           base36.QuirkyItoa32(u.ID),
		Username:     u.Username,
		PasswordHash: u.PasswordHash,
	}
}
