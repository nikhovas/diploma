package clients

import (
	"context"
	"database/sql"
	"github.com/go-redis/redis/v8"
	consulApi "github.com/hashicorp/consul/api"
	"github.com/nikhovas/diploma/go/lib/proto/consumer_bot"
	ctrl "github.com/nikhovas/diploma/go/lib/proto/controller"
	qw "github.com/nikhovas/diploma/go/lib/proto/question_worker"
	tsb "github.com/nikhovas/diploma/go/lib/proto/staff_bot"
	"github.com/nikhovas/diploma/go/lib/utils/env"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateControllerClient() (*grpc.ClientConn, ctrl.ControllerClient) {
	conn, err := grpc.Dial(env.GetControllerGrpcHost(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := ctrl.NewControllerClient(conn)
	return conn, client
}

func CreateRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	return rdb
}

func CreateTelegramStaffBotClient() (*grpc.ClientConn, tsb.TelegramStaffBotClient) {
	conn, err := grpc.Dial("[::]:50060", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := tsb.NewTelegramStaffBotClient(conn)
	return conn, client
}

func CreateQuestionWorkerClient() (*grpc.ClientConn, qw.QuestionWorkerClient) {
	conn, err := grpc.Dial("[::]:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := qw.NewQuestionWorkerClient(conn)
	return conn, client
}

func CreateSqlConn() *sql.DB {
	db, err := sql.Open("postgres", "postgres://user:password@rc1b-y3kpmg61fpuvucup.mdb.yandexcloud.net:6432/data")
	if err != nil {
		panic(err)
	}

	return db
}

func CreateConsulClient() *consulApi.Client {
	client, _ := consulApi.NewClient(consulApi.DefaultConfig())
	return client
}

func CreateConsumerBotClient() (*grpc.ClientConn, consumer_bot.VkServerClient) {
	conn, err := grpc.Dial("localhost:5555", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := consumer_bot.NewVkServerClient(conn)
	return conn, client
}
