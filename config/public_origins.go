package config

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

const publicOriginsEnv = "COURIER_PUBLIC_ORIGINS"

// ApplyPublicOrigins merges the comma-separated public browser origins from the
// runtime environment with the configured CORS and WebSocket allow-list.
func ApplyPublicOrigins(cfg *AppConfig) error {
	if cfg == nil || cfg.Http == nil {
		return nil
	}

	origins, err := ParsePublicOrigins(os.Getenv(publicOriginsEnv))
	if err != nil {
		return err
	}

	cfg.Http.AllowOrigins = mergeOrigins(cfg.Http.AllowOrigins, origins)
	return nil
}

// ParsePublicOrigins parses COURIER_PUBLIC_ORIGINS as a comma-separated list
// of browser origins, for example "https://courier.io.vn,http://localhost:5173".
func ParsePublicOrigins(raw string) ([]string, error) {
	if strings.TrimSpace(raw) == "" {
		return nil, nil
	}

	origins := make([]string, 0)
	for _, value := range strings.Split(raw, ",") {
		origin := strings.TrimSpace(value)
		if origin == "" {
			continue
		}

		parsed, err := url.ParseRequestURI(origin)
		if err != nil || parsed.Scheme == "" || parsed.Host == "" ||
			parsed.Path != "" || parsed.RawQuery != "" || parsed.Fragment != "" {
			return nil, fmt.Errorf("%s contains invalid origin %q", publicOriginsEnv, origin)
		}
		if parsed.Scheme != "http" && parsed.Scheme != "https" {
			return nil, fmt.Errorf("%s contains unsupported origin scheme %q", publicOriginsEnv, parsed.Scheme)
		}

		origins = append(origins, parsed.Scheme+"://"+parsed.Host)
	}

	return mergeOrigins(nil, origins), nil
}

func mergeOrigins(existing []string, additions []string) []string {
	seen := make(map[string]struct{}, len(existing)+len(additions))
	merged := make([]string, 0, len(existing)+len(additions))
	for _, origin := range append(existing, additions...) {
		if _, ok := seen[origin]; ok {
			continue
		}
		seen[origin] = struct{}{}
		merged = append(merged, origin)
	}
	return merged
}
