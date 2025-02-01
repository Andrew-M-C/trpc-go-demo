package account

import (
	"time"

	"github.com/Andrew-M-C/go.util/ids/base36"
	dbutil "github.com/Andrew-M-C/go.util/mysql"
	pb "github.com/Andrew-M-C/trpc-go-demo/proto/user"
)

const userAccountTableName = "t_trpc_demo_user_account"

type userAccountItem struct {
	ID           uint32    `db:"id"`
	Username     string    `db:"username"`
	PasswordHash string    `db:"password_hash"`
	CreateAt     time.Time `db:"create_at"`
	Nickname     string    `db:"nickname"`
}

func (u userAccountItem) toPB(tzCorrector dbutil.TimezoneCorrector) *pb.UserInfo {
	return &pb.UserInfo{
		UserId:       base36.QuirkyItoa32(u.ID),
		Username:     u.Username,
		PasswordHash: u.PasswordHash,
		Nickname:     u.Nickname,
		CreateTsSec:  tzCorrector.CorrectTimezoneFromDB(u.CreateAt).Unix(),
	}
}
