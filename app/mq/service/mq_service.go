package service

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"strings"

	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	pb "github.com/Andrew-M-C/trpc-go-demo/proto/mq"
	"github.com/Andrew-M-C/trpc-go-utils/log"
	"github.com/Andrew-M-C/trpc-go-utils/tracelog"
	"github.com/IBM/sarama"
	"trpc.group/trpc-go/trpc-database/kafka"
	"trpc.group/trpc-go/trpc-go/server"
)

// 注册 MQ demo 服务
func RegisterMQService(s *server.Server, d Dependency) error {
	impl := &impl{Dependency: d}
	return impl.initialize(s)
}

type Dependency struct {
	StringKafka func(context.Context) (kafka.Client, error)
	JSONKafka   func(context.Context) (kafka.Client, error)
}

type impl struct {
	Dependency
}

func (impl *impl) initialize(s *server.Server) error {
	// kafka.RegisterBatchHandlerService(s.Service("kafka.amc.demo.mq"), impl.batchConsumeMQ)
	kafka.RegisterKafkaHandlerService(s.Service("kafka.amc.demo.mq"), impl.onceConsumeMQ)
	pb.RegisterMQService(s, impl)
	return nil
}

// 这个方法不知道为什么收不到
func (impl *impl) batchConsumeMQ(
	ctx context.Context, msgList []*sarama.ConsumerMessage,
) error {
	log.New(ctx).Text("收到批量 Kafka 消息").
		Int("count", len(msgList)).
		Stringer("data", tracelog.ToJSON(msgList)).
		Info()
	return nil
}

func (impl *impl) onceConsumeMQ(
	ctx context.Context, msg *sarama.ConsumerMessage,
) error {
	log.New(ctx).Text("收到一条 Kafka 消息").
		Stringer("data", tracelog.ToJSON(msg)).
		Info()
	return nil
}

func (impl *impl) TestMQAdd(
	ctx context.Context, req *pb.TestMQAddRequest,
) (*pb.TestMQAddResponse, error) {
	var cli kafka.Client
	var err error
	var genData func() []byte

	// 参数检查
	switch strings.ToLower(req.GetType()) {
	case "json":
		cli, err = impl.JSONKafka(ctx)
		genData = func() []byte {
			v := jsonvalue.NewObject()
			v.At("int").Set(rand.Int63n(999999999999))
			return v.MustMarshal()
		}

	case "string", "str":
		cli, err = impl.StringKafka(ctx)
		genData = func() []byte {
			buff := bytes.Buffer{}
			buff.WriteString(fmt.Sprint(rand.Int63()))
			return buff.Bytes()
		}

	default:
		err = fmt.Errorf("非法的类型值 '%s'", req.GetType())
	}
	if err != nil {
		return nil, err
	}

	// 数据生成
	for i := 0; i < int(req.GetCount()); i++ {
		if err := cli.Produce(ctx, nil, genData()); err != nil {
			return nil, fmt.Errorf("写入 DB 失败")
		}
	}

	log.New(ctx).Text("写入成功").Int("count", int(req.GetCount())).Debug()
	return &pb.TestMQAddResponse{}, nil
}
