package config

import "github.com/vucongthanh92/go-base-utils/models"

type AppConfig struct {
	ServiceName      string                  `mapstructure:"serviceName"`
	Development      bool                    `mapstructure:"development"`
	Logger           *LoggerConfig           `mapstructure:"logger"`
	Http             *HttpConfig             `mapstructure:"http"`
	GRPC             *GrpcConfig             `mapstructure:"grpc"`
	Client           *ClientConfig           `mapstructure:"client"`
	Database         *DatabaseConfig         `mapstructure:"database"`
	Tracing          *TracingConfig          `mapstructure:"tracing"`
	Redis            *RedisConfig            `mapstructure:"redis"`
	Kafka            *KafkaConfig            `mapstructure:"kafka"`
	MessageRateLimit *MessageRateLimitConfig `mapstructure:"messageRateLimit"`
	Heathcheck       *HeathcheckConfig       `mapstructure:"heathcheck"`
	Metrics          *MetricsConfig          `mapstructure:"metrics"`
	CronJob          *CronJob                `mapstructure:"cronjob"`
	Loki             *LokiConfig             `mapstructure:"loki"`
}

type LoggerConfig = models.LoggerConfig
type HttpConfig = models.HttpConfig
type GrpcConfig = models.GrpcConfig
type ClientConfig = models.GrpcClientConfig
type DatabaseConfig = models.DatabaseConfig
type ReadDbConfig = models.ReadDbConfig
type WriteDbConfig = models.WriteDbConfig
type TracingConfig = models.TracingConfig
type RedisConfig = models.RedisConfig

type KafkaConfig struct {
	Brokers []string          `mapstructure:"brokers"`
	GroupID string            `mapstructure:"groupID"`
	Topics  KafkaTopicsConfig `mapstructure:"topics"`
}

type KafkaTopicsConfig struct {
	UserEvents         string `mapstructure:"userEvents"`
	ChatEvents         string `mapstructure:"chatEvents"`
	AssistantRequested string `mapstructure:"assistantRequested"`
	AssistantResponded string `mapstructure:"assistantResponded"`
}

type RateLimitWindowConfig struct {
	Limit         int `mapstructure:"limit"`
	WindowSeconds int `mapstructure:"windowSeconds"`
}

type MessageRateLimitConfig struct {
	Enabled      bool                  `mapstructure:"enabled"`
	Burst        RateLimitWindowConfig `mapstructure:"burst"`
	Conversation RateLimitWindowConfig `mapstructure:"conversation"`
	User         RateLimitWindowConfig `mapstructure:"user"`
}

type HeathcheckConfig = models.HeathcheckConfig
type MetricsConfig = models.MetricsConfig
type CronJob = models.CronJob
type LokiConfig = models.LokiConfig
