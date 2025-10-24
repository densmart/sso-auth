package mockdb

import (
	"time"

	"github.com/densmart/sso-auth/pkg/logger"
)

var defaultCreatedAt = time.Date(2025, 5, 25, 10, 11, 12, 13, time.UTC)

type MockDB struct{}

func NewMockDB() *MockDB {
	logger.Debugf("Using MockDB (no real connection)")
	return &MockDB{}
}

func (db *MockDB) Close() {
	logger.Debugf("MockDB closed (no-op)")
}

func (db *MockDB) MigrationUp() error {
	logger.Debugf("MockDB Migration Up (no-op)")
	return nil
}

func (db *MockDB) MigrationDown() error {
	logger.Debugf("MockDB Migration Down (no-op)")
	return nil
}
