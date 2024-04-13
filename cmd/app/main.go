package main

import (
	"context"
	"database/sql"
	"github.com/heroiclabs/nakama-common/runtime"
	"zepto-lab.com/file-reader/internal/api"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	logger.Info("Start init")
	err := initializer.RegisterRpc("file-reader", api.ProcessRPCRequest)
	if err != nil {
		return err
	}
	return nil
}
