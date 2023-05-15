package main

import (
	"log"
	"time"

	"github.com/Astemirdum/lavka/internal/app"
	"github.com/Astemirdum/lavka/internal/config"
	"github.com/joho/godotenv"
	"go.uber.org/zap/zapcore"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("load envs from .env ", err)
	}
	cfg := config.NewConfig(
		config.WithLogLevel(zapcore.DebugLevel),
		config.WithWriteTimeout(time.Second*5),
	)
	app.Run(cfg)
}
