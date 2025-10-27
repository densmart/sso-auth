package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePasswordHash(t *testing.T) {
	password := "my-secure-password"
	hash, err := GeneratePasswordHash(password)
	// assert no error occurred
	assert.NoError(t, err)
	assert.NotNil(t, hash)
	assert.NotEmpty(t, hash)

	verified := CheckPasswordHash(hash, password)
	assert.True(t, verified)
}

func TestCheckPasswordHash_InvalidPassword(t *testing.T) {
	password := "my-secure-password"
	wrongPassword := "wrong-password"
	hash, err := GeneratePasswordHash(password)
	// assert no error occurred
	assert.NoError(t, err)
	assert.NotNil(t, hash)
	assert.NotEmpty(t, hash)

	verified := CheckPasswordHash(hash, wrongPassword)
	assert.False(t, verified)
}

func TestCheckPasswordHash_InvalidHash(t *testing.T) {
	password := "my-secure-password"
	invalidHash := "invalid-hash-string"

	verified := CheckPasswordHash(invalidHash, password)
	assert.False(t, verified)
}

func TestGeneratePasswordHash_EmptyPassword(t *testing.T) {
	password := ""
	hash, err := GeneratePasswordHash(password)
	// assert no error occurred
	assert.NoError(t, err)
	assert.NotNil(t, hash)
	assert.NotEmpty(t, hash)

	verified := CheckPasswordHash(hash, password)
	assert.True(t, verified)
}

func TestGeneratePasswordHash_SpecialCharacters(t *testing.T) {
	password := "p@$$w0rd!#%"
	hash, err := GeneratePasswordHash(password)
	// assert no error occurred
	assert.NoError(t, err)
	assert.NotNil(t, hash)
	assert.NotEmpty(t, hash)

	verified := CheckPasswordHash(hash, password)
	assert.True(t, verified)
}

func TestGeneratePasswordHash_LongPassword(t *testing.T) {
	password := "thisisaverylongpasswordthatexceedstheusualcharacterlimitsetbymostapplications1234567890!@#$%^&*()"
	hash, err := GeneratePasswordHash(password)
	// assert no error occurred
	assert.Error(t, err)
	assert.Empty(t, hash)

	password = "thisisaverylongpasswordthatexceedstheusualcharacterlimitsetbymostapplica"
	hash, err = GeneratePasswordHash(password)
	// assert no error occurred
	assert.NoError(t, err)
	assert.NotNil(t, hash)
	assert.NotEmpty(t, hash)

	verified := CheckPasswordHash(hash, password)
	assert.True(t, verified)
}

func TestGeneratePasswordHash_UnicodePassword(t *testing.T) {
	password := "pässwörd😊漢字"
	hash, err := GeneratePasswordHash(password)
	// assert no error occurred
	assert.NoError(t, err)
	assert.NotNil(t, hash)
	assert.NotEmpty(t, hash)

	verified := CheckPasswordHash(hash, password)
	assert.True(t, verified)
}

func TestGeneratePasswordHash_RepeatedCalls(t *testing.T) {
	password := "consistent-password"
	hash1, err1 := GeneratePasswordHash(password)
	hash2, err2 := GeneratePasswordHash(password)
	// assert no error occurred
	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.NotNil(t, hash1)
	assert.NotNil(t, hash2)
	assert.NotEmpty(t, hash1)
	assert.NotEmpty(t, hash2)
	// hashes should be different due to salting
	assert.NotEqual(t, hash1, hash2)

	verified1 := CheckPasswordHash(hash1, password)
	verified2 := CheckPasswordHash(hash2, password)
	assert.True(t, verified1)
	assert.True(t, verified2)
}
