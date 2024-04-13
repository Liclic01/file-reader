package api

import (
	"context"
	"database/sql"
	"github.com/heroiclabs/nakama-common/runtime"
	"zepto-lab.com/file-reader/internal/repository"
	"zepto-lab.com/file-reader/internal/service"
)

func ProcessRPCRequest(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	return service.Reader(logger, db, &service.DefaultFileReader{}, &repository.SqlJsonRepository{}, payload)
}
