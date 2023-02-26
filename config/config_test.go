package config_test

import (
	"os"
	"sync"
	"testing"

	"github.com/andresxlp/backend-twitter/config"
	"github.com/stretchr/testify/assert"
)

func TestEnvironments_WhenAllVarsAreSet(t *testing.T) {
	config.Environments()

	assert.Equal(t, "default", config.Cfg.SecretJWT)
}

func TestEnvironments_WhenSomeVarIsMissed(t *testing.T) {
	expectedError := "Error parsing environment vars &errors.errorString{s:\"required key SERVER_HOST missing value\"}"
	config.Once = sync.Once{}

	os.Clearenv()

	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, expectedError, r)
		}
	}()

	config.Environments()
}

//
//func TestIsDevEnvironment_WhenTrue(t *testing.T) {
//	_ = os.Setenv("APP_POSTFIX", "dev")
//
//	assert.True(t, IsDevEnvironment())
//}
