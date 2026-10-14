package config

import "github.com/vucongthanh92/go-base-utils/models"

type AppConfig struct {
	ServiceName      string                         `mapstructure:"serviceName"`
	Development      bool                           `mapstructure:"development"`
	Logger           *models.LoggerConfig           `mapstructure:"logger"`
	Http             *models.HttpConfig             `mapstructure:"http"`
	GRPC             *models.GrpcConfig             `mapstructure:"grpc"`
	Client           *models.GrpcClientConfig       `mapstructure:"client"`
	Database         *models.DatabaseConfig         `mapstructure:"database"`
	Tracing          *models.TracingConfig          `mapstructure:"tracing"`
	Redis            *models.RedisConfig            `mapstructure:"redis"`
	Kafka            *models.KafkaConfigV2          `mapstructure:"kafka"`
	MessageRateLimit *models.MessageRateLimitConfig `mapstructure:"messageRateLimit"`
	Heathcheck       *models.HeathcheckConfig       `mapstructure:"heathcheck"`
	Metrics          *models.MetricsConfig          `mapstructure:"metrics"`
	CronJob          *models.CronJob                `mapstructure:"cronjob"`
	Loki             *models.LokiConfig             `mapstructure:"loki"`
}
