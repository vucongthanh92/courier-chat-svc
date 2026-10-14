package v1

import (
	"os"
	"testing"
	"time"

	baseUtils "github.com/vucongthanh92/go-base-utils/helper/utils"
	"github.com/vucongthanh92/go-base-utils/tracing"
)

func TestMain(m *testing.M) {
	tracing.UseOpenTelemetry(tracing.Config{ServiceName: "chat-service-test"})
	baseUtils.InitSnowflake(baseUtils.SnowflakeConfig{
		MachineID:   1,
		CustomEpoch: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	})
	os.Exit(m.Run())
}
