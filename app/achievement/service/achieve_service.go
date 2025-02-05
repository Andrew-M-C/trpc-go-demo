package service

import (
	"context"

	"github.com/Andrew-M-C/trpc-go-demo/app/achievement/repo"
	pb "github.com/Andrew-M-C/trpc-go-demo/proto/achieve"
	"github.com/go-playground/validator/v10"
)

type Dependency struct {
	BadgeRepo      repo.Badge      `validate:"required"`
	ReputationRepo repo.Reputation `validate:"required"`
}

type impl struct {
	Dependency
}

func New(d Dependency) (pb.AchievementService, error) {
	if err := validator.New().Struct(d); err != nil {
		return nil, err
	}
	return &impl{d}, nil
}

var rankMapping = []string{
	"白板",
	"青铜",
	"白银",
	"黄金",
	"翡翠",
	"璧玉",
}

func (impl *impl) GetUserReputation(
	ctx context.Context, req *pb.GetUserReputationRequest,
) (*pb.GetUserReputationResponse, error) {
	// DB 查询
	res, err := impl.ReputationRepo.GetUserReputation(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}

	// 映射等级名称
	rankName := func() string {
		if int(res.Rank) > len(rankMapping)-1 {
			return rankMapping[len(rankMapping)-1]
		}
		return rankMapping[int(res.Rank)]
	}()

	// 打包返回
	rsp := &pb.GetUserReputationResponse{
		Data: &pb.GetUserReputationResponse_Data{},
	}
	rsp.Data.Reputation = &pb.Reputation{
		ExpPoints: res.ExpPoints,
		Rank:      res.Rank,
		RankName:  rankName,
	}
	return rsp, nil
}

func (impl *impl) GetUserBadges(
	ctx context.Context, req *pb.GetUserBadgesRequest,
) (*pb.GetUserBadgesResponse, error) {
	// DB 查询
	res, err := impl.BadgeRepo.GetUserBadges(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}

	// 聚合打包返回
	m := map[string]*pb.Badge{}
	rsp := &pb.GetUserBadgesResponse{
		Data: &pb.GetUserBadgesResponse_Data{},
	}
	for _, b := range res {
		if item, exist := m[b.Name]; exist {
			item.Count++
			continue
		}
		item := &pb.Badge{
			Name:              b.Name,
			Count:             1,
			FirstAcquireTsSec: b.CreateTime.Unix(),
		}
		m[b.Name] = item
		rsp.Data.Badges = append(rsp.Data.Badges, item)
	}
	return rsp, nil
}
