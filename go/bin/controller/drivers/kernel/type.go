package kernel

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v8"
	qw "github.com/nikhovas/diploma/go/lib/proto/question_worker"
	tsb "github.com/nikhovas/diploma/go/lib/proto/staff_bot"
)

type NoSuchItemError struct {
}

func (e *NoSuchItemError) Error() string {
	return fmt.Sprintf("NoSuchItemError")
}

type Kernel struct {
	db               *sql.DB
	rdb              *redis.Client
	qwClient         qw.QuestionWorkerClient
	telegramStaffBot tsb.TelegramStaffBotClient
}

func NewKernel(db *sql.DB, rdb *redis.Client, qwClient qw.QuestionWorkerClient, telegramStaffBot tsb.TelegramStaffBotClient) Kernel {
	return Kernel{
		db:               db,
		rdb:              rdb,
		qwClient:         qwClient,
		telegramStaffBot: telegramStaffBot,
	}
}