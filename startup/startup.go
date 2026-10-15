package startup

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gammazero/workerpool"
	"github.com/vucongthanh92/courier-chat-svc/config"
	"github.com/vucongthanh92/courier-chat-svc/database"
	healthcheck "github.com/vucongthanh92/courier-chat-svc/helper/healthcheck"
	"github.com/vucongthanh92/courier-chat-svc/internal"
	"github.com/vucongthanh92/courier-chat-svc/internal/api"
	"github.com/vucongthanh92/courier-chat-svc/redis"
	"github.com/vucongthanh92/go-base-utils/command"
	baseUtils "github.com/vucongthanh92/go-base-utils/helper/utils"
	"github.com/vucongthanh92/go-base-utils/logger"
	"github.com/vucongthanh92/go-base-utils/tracing"
	"go.uber.org/zap"
)

var cfg *config.AppConfig

func Execute() {
	cfg = &config.AppConfig{}

	command.UseCommands(
		command.WithStartCommand(start, cfg),
	)
}

func start() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	tracing.UseOpenTelemetry(tracing.Config(*cfg.Tracing))
	baseUtils.InitSnowflake(baseUtils.SnowflakeConfig{
		MachineID:   1,
		CustomEpoch: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	})

	container, readDb, writeDb := registerDependencies(ctx)
	runServer(ctx, container, readDb, writeDb)
}

func registerDependencies(_ context.Context) (*api.ApiContainer, database.GormReadDb, database.GormWriteDb) {
	redisClient := redis.Open(cfg.Redis)
	readDb, writeDb := database.GetConnectByGorm(cfg.Database)

	return internal.InitializeContainer(
		cfg,
		&readDb,
		&writeDb,
		redisClient,
	), readDb, writeDb
}

func runServer(
	ctx context.Context,
	container *api.ApiContainer,
	readDb database.GormReadDb,
	writeDb database.GormWriteDb,
) {
	wp := workerpool.New(5)

	wp.Submit(container.GrpcServer.Run)
	wp.Submit(container.HttpServer.Run)
	wp.Submit(container.CronServer.Run)
	wp.Submit(healthcheck.RunHealthCheck(ctx, cfg, readDb, writeDb))
	wp.Submit(func() {
		if err := container.UserEventConsumer.Start(ctx); err != nil && !errors.Is(err, context.Canceled) {
			logger.Error("user event consumer stopped", zap.Error(err))
		}
	})
	wp.Submit(func() {
		if err := container.ChatEventConsumer.Start(ctx); err != nil && !errors.Is(err, context.Canceled) {
			logger.Error("chat event consumer stopped", zap.Error(err))
		}
	})
	wp.Submit(func() {
		if err := container.AssistantResponseConsumer.Start(ctx); err != nil && !errors.Is(err, context.Canceled) {
			logger.Error("assistant response consumer stopped", zap.Error(err))
		}
	})

	wp.StopWait()
	if container.UserEventConsumer != nil {
		_ = container.UserEventConsumer.Close()
	}
	if container.ChatEventConsumer != nil {
		_ = container.ChatEventConsumer.Close()
	}
	if container.AssistantResponseConsumer != nil {
		_ = container.AssistantResponseConsumer.Close()
	}

	logger.Info("chat-service stopped")
}
