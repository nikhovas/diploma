package telegram

import (
	kernel2 "control/drivers/kernel"
	"database/sql"
	ctrlProto "github.com/nikhovas/diploma/go/lib/proto/controller"
)

func generateBadChatTypeResponse() (*ctrlProto.DefaultResponse, error) {
	return &ctrlProto.DefaultResponse{Resp: &ctrlProto.DefaultResponse_BadChatType{BadChatType: &ctrlProto.BadChatType{}}}, nil
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
