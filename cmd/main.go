package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/densmart/sso-auth/internal/domain/repo"
	"github.com/densmart/sso-auth/pkg/configger"
	"github.com/densmart/sso-auth/pkg/logger"
)

func main() {
	configger.InitConfig(configger.DefaultCfgPath, "config", "yaml")
	logger.InitLogger()
	logger.Infof("logger initialized. level: %s", logger.GetLevel())

	appCtx, cancel := context.WithCancel(context.Background())

	logger.Infof("starting DB connection...")
	repository, err := repo.NewRepo(appCtx, "mockdb")
	if err != nil {
		logger.Fatalf("error starting Repo: %s", err.Error())
	}
	logger.Infof("DB connection established")

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	// wait for shutdown
	<-quit

	logger.Infof("stopping app...")
	// cancel context
	cancel()
	logger.Infof("...closing DB connections...")
	repository.Close()
	logger.Infof("...app stopped!")
}
