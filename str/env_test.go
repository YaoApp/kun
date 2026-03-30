package str

import (
	"encoding/base64"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvVar(t *testing.T) {
	os.Setenv("TEST_STR_ENVVAR_HOST", "192.168.1.1")
	os.Setenv("TEST_STR_ENVVAR_PORT", "22")
	defer os.Unsetenv("TEST_STR_ENVVAR_HOST")
	defer os.Unsetenv("TEST_STR_ENVVAR_PORT")

	assert.Equal(t, "192.168.1.1", EnvVar("$ENV.TEST_STR_ENVVAR_HOST"))
	assert.Equal(t, "22", EnvVar("$ENV.TEST_STR_ENVVAR_PORT"))
}

func TestEnvVarMissing(t *testing.T) {
	os.Unsetenv("TEST_STR_ENVVAR_MISSING")
	assert.Equal(t, "", EnvVar("$ENV.TEST_STR_ENVVAR_MISSING"))
}

func TestEnvVarPassthrough(t *testing.T) {
	assert.Equal(t, "plain-value", EnvVar("plain-value"))
	assert.Equal(t, "", EnvVar(""))
	assert.Equal(t, "https://example.com", EnvVar("https://example.com"))
}

func TestEnvVarB64(t *testing.T) {
	secret := "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEA...\n-----END RSA PRIVATE KEY-----\n"
	encoded := base64.StdEncoding.EncodeToString([]byte(secret))
	os.Setenv("TEST_STR_ENVVAR_B64_KEY", encoded)
	defer os.Unsetenv("TEST_STR_ENVVAR_B64_KEY")

	assert.Equal(t, secret, EnvVar("$ENV_B64.TEST_STR_ENVVAR_B64_KEY"))
}

func TestEnvVarB64InvalidFallback(t *testing.T) {
	os.Setenv("TEST_STR_ENVVAR_B64_BAD", "not-valid-base64!!!")
	defer os.Unsetenv("TEST_STR_ENVVAR_B64_BAD")

	assert.Equal(t, "not-valid-base64!!!", EnvVar("$ENV_B64.TEST_STR_ENVVAR_B64_BAD"))
}

func TestEnvVarB64Missing(t *testing.T) {
	os.Unsetenv("TEST_STR_ENVVAR_B64_MISSING")
	assert.Equal(t, "", EnvVar("$ENV_B64.TEST_STR_ENVVAR_B64_MISSING"))
}
