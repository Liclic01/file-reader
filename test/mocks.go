package test

import (
	"database/sql"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/stretchr/testify/mock"
)

type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Debug(format string, v ...interface{}) {
	m.Called(format, v)
}

func (m *MockLogger) Info(format string, v ...interface{}) {
	m.Called(format, v)
}

func (m *MockLogger) Warn(format string, v ...interface{}) {
	m.Called(format, v)
}

func (m *MockLogger) Error(format string, v ...interface{}) {
	m.Called(format, v)
}

func (m *MockLogger) WithField(key string, v interface{}) runtime.Logger {
	args := m.Called(key, v)
	return args.Get(0).(runtime.Logger)
}

func (m *MockLogger) WithFields(fields map[string]interface{}) runtime.Logger {
	args := m.Called(fields)
	return args.Get(0).(runtime.Logger)
}

func (m *MockLogger) Fields() map[string]interface{} {
	args := m.Called()
	return args.Get(0).(map[string]interface{})
}

type MockFileReader struct {
	mock.Mock
}

func (m *MockFileReader) ReadFileContent(path string, fileType string, fileVersion string) (string, error) {
	args := m.Called(path, fileType, fileVersion)
	return args.String(0), args.Error(1)
}

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) SaveJson(db *sql.DB, id string, hash string, content interface{}) error {
	args := m.Called(db, id, hash, content)
	return args.Error(0)
}

type MockJsonRepository struct {
	mock.Mock
}

func (m *MockJsonRepository) SaveJson(id string, hash string, content interface{}) error {
	args := m.Called(id, hash, content)
	return args.Error(0)
}
