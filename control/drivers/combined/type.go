package combined

import (
	"control/drivers/kernel"
	"control/drivers/telegram"
	pb "control/grpc/control"
	"database/sql"
)

type Server struct {
	pb.UnimplementedControlServer
	db       *sql.DB
	telegram *telegram.Telegram
	kernel   *kernel.Kernel
}

func NewServer(db *sql.DB, telegram *telegram.Telegram, kernel *kernel.Kernel) *Server {
	s := &Server{db: db, telegram: telegram, kernel: kernel}
	return s
}
