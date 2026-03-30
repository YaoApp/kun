package str

import (
	"encoding/base64"
	"os"
	"strings"
)

// EnvVar resolves environment variable references in configuration values.
//
//	"$ENV.FOO"      → os.Getenv("FOO")
//	"$ENV_B64.FOO"  → base64-decode(os.Getenv("FOO"))
//	anything else   → returned as-is
func EnvVar(value string) string {
	if strings.HasPrefix(value, "$ENV_B64.") {
		raw := os.Getenv(value[9:])
		if decoded, err := base64.StdEncoding.DecodeString(raw); err == nil {
			return string(decoded)
		}
		return raw
	}
	if strings.HasPrefix(value, "$ENV.") {
		return os.Getenv(value[5:])
	}
	return value
}
