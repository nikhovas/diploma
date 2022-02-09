package grpcServer

import (
	"control/modules/combined"
	"database/sql"
	ctrlProto "github.com/nikhovas/diploma/go/lib/proto/controller"
)

type Server struct {
	ctrlProto.UnimplementedControllerServer
	Combined *combined.Combined
	Database *sql.DB
}
