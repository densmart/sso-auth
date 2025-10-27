package utils

import (
	"testing"

	"github.com/densmart/sso-auth/pkg/configger"
	"github.com/densmart/sso-auth/pkg/logger"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestCreateOtpSecret(t *testing.T) {
	otpKey, err := CreateOtpSecret()
	assert.NoError(t, err)
	assert.Equal(t, len(otpKey), 16)
}

func TestCreateOtpURL(t *testing.T) {
	// initialize config
	configger.InitConfig("../../config", "config", "yaml")
	logger.InitLogger()
	// init test data
	email := "john@doe.com"
	issuer := viper.GetString("app.name")

	otpKey, err := CreateOtpSecret()
	assert.NoError(t, err)

	otpURL, err := CreateOtpURL(email, otpKey)

	assert.NoError(t, err)
	assert.Contains(t, otpURL, "otpauth://totp/"+email)
	assert.Contains(t, otpURL, "secret="+otpKey)
	assert.Contains(t, otpURL, "issuer="+issuer)
}

func TestVerifyOTP(t *testing.T) {
	// initialize config
	configger.InitConfig("../../config", "config", "yaml")
	logger.InitLogger()

	otpKey, err := CreateOtpSecret()
	assert.NoError(t, err)

	err = VerifyOTP(otpKey, "efwefewfwefwef")
	assert.Error(t, err)
}
