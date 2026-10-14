package message

import (
	"os"
	"testing"
	"time"

	baseUtils "github.com/vucongthanh92/go-base-utils/helper/utils"
)

func TestMain(m *testing.M) {
	baseUtils.InitSnowflake(baseUtils.SnowflakeConfig{
		MachineID:   1,
		CustomEpoch: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	})
	os.Exit(m.Run())
}
