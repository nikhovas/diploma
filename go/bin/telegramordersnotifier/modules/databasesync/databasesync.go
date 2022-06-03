package databasesync

import (
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"github.com/nikhovas/diploma/go/lib/utils/clients"
	"github.com/nikhovas/diploma/go/lib/utils/foundation"
	"log"
	"sync"
	"time"
)

type DatabaseSync struct {
	db *sql.DB
	updateTime time.Duration
	storage map[int]int
}

func FromConfig(config Config) *DatabaseSync {
	ds := &DatabaseSync{}

	ds.db = clients.CreateSqlConn()
	ds.updateTime = time.Duration(config.UpdateTimeInterval) * time.Second
	ds.storage = make(map[int]int)

	return ds
}

func (ds *DatabaseSync) Run(ctx context.Context, wg *sync.WaitGroup) {
	foundation.RunPeriodic(ctx, wg, ds.worker, ds.updateTime)
}

func (ds *DatabaseSync) worker(ctx context.Context) {
	rows, err := squirrel.Select("shop_id", "telegram_group_id").From("telegram_staff_shop_group").
		PlaceholderFormat(squirrel.Dollar).RunWith(ds.db).QueryContext(ctx)
	if err != nil {
		log.Fatal(err)
	}

	newStorage := make(map[int]int)

	for rows.Next() {
		var shopId int
		var telegramId int
		_ = rows.Scan(&shopId, &telegramId)
		newStorage[shopId] = telegramId
	}

	ds.storage = newStorage
}

func (ds *DatabaseSync) GetTelegramGroup(shopId int) (int, bool) {
	groupId, found := ds.storage[shopId]
	return groupId, found
}

func (ds *DatabaseSync) Stop() {
	err := ds.db.Close()
	if err != nil {}
}
