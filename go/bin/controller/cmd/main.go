package main

import (
	"control/drivers/combined"
	"control/drivers/kernel"
	"control/drivers/telegram"
	"database/sql"
	"fmt"
	ctrlProto "github.com/nikhovas/diploma/go/lib/proto/controller"
	qw "github.com/nikhovas/diploma/go/lib/proto/question_worker"
	tsb "github.com/nikhovas/diploma/go/lib/proto/staff_bot"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

func createTelegramStaffBotClient() (*grpc.ClientConn, tsb.TelegramStaffBotClient) {
	conn, err := grpc.Dial("[::]:50060", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := tsb.NewTelegramStaffBotClient(conn)
	return conn, client
}

func createQuestionWorkerClient() (*grpc.ClientConn, qw.QuestionWorkerClient) {
	conn, err := grpc.Dial("[::]:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := qw.NewQuestionWorkerClient(conn)
	return conn, client
}

func createRedisConn() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})
	return rdb
}

func createSqlConn() *sql.DB {
	db, err := sql.Open("postgres", "postgres://user:password@rc1b-y3kpmg61fpuvucup.mdb.yandexcloud.net:6432/data")
	if err != nil {
		panic(err)
	}

	return db
}

func createServerSocket() net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:7777"))
	if err != nil {
		panic(err)
	}

	return lis
}

func main() {
	qwConn, qwClient := createQuestionWorkerClient()
	defer qwConn.Close()
	tsbConn, tsbClient := createTelegramStaffBotClient()
	defer tsbConn.Close()
	redisDb := createRedisConn()
	sqlDb := createSqlConn()
	lis := createServerSocket()

	grpcServer := grpc.NewServer()

	kernelInstance := kernel.NewKernel(sqlDb, redisDb, qwClient, tsbClient)
	telegramInstance := telegram.NewTelegram(sqlDb, &kernelInstance)
	serverInstance := combined.NewServer(sqlDb, &telegramInstance, &kernelInstance)

	ctrlProto.RegisterControllerServer(grpcServer, serverInstance)
	err := grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
