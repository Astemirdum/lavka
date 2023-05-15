package app

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Astemirdum/lavka/pkg/cache"

	"github.com/Astemirdum/lavka/pkg/logger"

	"github.com/Astemirdum/lavka/pkg/postgres"

	"github.com/Astemirdum/lavka/internal/config"
	"github.com/Astemirdum/lavka/internal/handler"
	"github.com/Astemirdum/lavka/internal/repository"
	"github.com/Astemirdum/lavka/internal/server"
	"github.com/Astemirdum/lavka/internal/service"
	"go.uber.org/zap"
)

func Run(cfg *config.Config) {
	log := logger.NewLogger(cfg.Log, "lavka")
	db, err := postgres.NewPostgresDB(&cfg.Database)
	if err != nil {
		log.Fatal("db init", zap.Error(err))
	}
	rs, err := cache.New()
	if err != nil {
		log.Fatal("rs cache", zap.Error(err))
	}

	repo, err := repository.NewRepository(db, log)
	if err != nil {
		log.Fatal("repo", zap.Error(err))
	}
	services := service.NewService(repo, log)
	h := handler.New(services, log, rs)

	srv := server.NewServer(cfg.Server, h.NewRouter())
	log.Info("http server start on: ",
		zap.String("addr",
			net.JoinHostPort(cfg.Server.Host, cfg.Server.Port)))
	go func() {
		if err := srv.Run(); err != nil {
			log.Fatal("server run", zap.Error(err))
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	termSig := <-sig

	log.Debug("Graceful shutdown", zap.Any("signal", termSig))

	closeCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err = srv.Stop(closeCtx); err != nil {
		log.DPanic("srv.Stop", zap.Error(err))
	}
	if err = db.Close(); err != nil {
		log.DPanic(" db.Close()", zap.Error(err))
	}
}
