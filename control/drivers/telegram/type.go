package telegram

import (
	kernel2 "control/drivers/kernel"
	pb "control/grpc/control"
	"database/sql"
)

func generateBadChatTypeResponse() (*pb.DefaultResponse, error) {
	return &pb.DefaultResponse{Resp: &pb.DefaultResponse_BadChatType{BadChatType: &pb.BadChatType{}}}, nil
}

type Telegram struct {
	db     *sql.DB
	kernel *kernel2.Kernel
}

func NewTelegram(db *sql.DB, kernel *kernel2.Kernel) Telegram {
	return Telegram{
		db:     db,
		kernel: kernel,
	}
}
