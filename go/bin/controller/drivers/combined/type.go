package combined

import (
	"control/drivers/kernel"
	"control/drivers/telegram"
	"database/sql"
	ctrlProto "github.com/nikhovas/diploma/go/lib/proto/controller"
)

type Server struct {
	ctrlProto.UnimplementedControllerServer
	db       *sql.DB
	telegram *telegram.Telegram
	kernel   *kernel.Kernel
}

func NewServer(db *sql.DB, telegram *telegram.Telegram, kernel *kernel.Kernel) *Server {
	s := &Server{db: db, telegram: telegram, kernel: kernel}
	return s
}
