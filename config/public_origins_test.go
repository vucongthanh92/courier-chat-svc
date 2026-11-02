package config

import (
	"reflect"
	"testing"

	"github.com/vucongthanh92/go-base-utils/models"
)

func TestApplyPublicOriginsMergesRuntimeOrigins(t *testing.T) {
	t.Setenv(publicOriginsEnv, "https://courier.io.vn, http://localhost:5173,https://courier.io.vn")
	cfg := &AppConfig{Http: &models.HttpConfig{AllowOrigins: []string{"http://localhost:8080"}}}

	if err := ApplyPublicOrigins(cfg); err != nil {
		t.Fatalf("ApplyPublicOrigins() error = %v", err)
	}

	want := []string{"http://localhost:8080", "https://courier.io.vn", "http://localhost:5173"}
	if !reflect.DeepEqual(cfg.Http.AllowOrigins, want) {
		t.Fatalf("AllowOrigins = %#v, want %#v", cfg.Http.AllowOrigins, want)
	}
}

func TestParsePublicOriginsRejectsPath(t *testing.T) {
	_, err := ParsePublicOrigins("https://courier.io.vn/conversa")
	if err == nil {
		t.Fatal("ParsePublicOrigins() error = nil, want invalid origin error")
	}
}
