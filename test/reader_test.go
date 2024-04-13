package test

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"zepto-lab.com/file-reader/internal/model"
	"zepto-lab.com/file-reader/internal/service"
)

func TestReader(t *testing.T) {
	mockLogger := new(MockLogger)
	mockFileReader := new(MockFileReader)
	mockRepository := new(MockRepository)
	mockDB := new(sql.DB)

	payload := `{"type": "test_type", "version": "1.0.0"}`

	mockLogger.On("Debug", mock.Anything, mock.Anything).Once()

	mockFileReader.On("ReadFileContent", "/path/inside/container", "test_type", "1.0.0").
		Return("file_content", nil).Once()

	mockRepository.On("SaveJson", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil).Once()
	result, err := service.Reader(mockLogger, mockDB, mockFileReader, mockRepository, payload)

	assert.NoError(t, err)

	var response model.Response
	err = json.Unmarshal([]byte(result), &response)
	assert.NoError(t, err)

	assert.Equal(t, "test_type", response.Type)
	assert.Equal(t, "1.0.0", response.Version)
	assert.NotEmpty(t, response.Hash)
	assert.NotEmpty(t, response.Content)

	mockFileReader.AssertExpectations(t)
	mockRepository.AssertExpectations(t)
}

func TestReader_FileReadError(t *testing.T) {
	mockLogger := new(MockLogger)
	mockFileReader := new(MockFileReader)
	mockRepository := new(MockRepository)
	mockDB := new(sql.DB)

	payload := `{"type": "test_type", "version": "1.0.0"}`

	mockLogger.On("Error", mock.Anything, mock.Anything).Once()

	mockFileReader.On("ReadFileContent", "/path/inside/container", "test_type", "1.0.0").
		Return("", fmt.Errorf("failed to read file")).Once()

	_, err := service.Reader(mockLogger, mockDB, mockFileReader, mockRepository, payload)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Failed to unmarshal request payload")

	mockLogger.AssertExpectations(t)
	mockFileReader.AssertExpectations(t)
	mockRepository.AssertExpectations(t)
}

func TestReader_SaveJsonError(t *testing.T) {
	mockLogger := new(MockLogger)
	mockFileReader := new(MockFileReader)
	mockRepository := new(MockRepository)
	mockDB := new(sql.DB)

	payload := `{"type": "test_type", "version": "1.0.0"}`

	mockLogger.On("Error", mock.Anything, mock.Anything).Once()

	mockFileReader.On("ReadFileContent", "/path/inside/container", "test_type", "1.0.0").
		Return("file_content", nil).Once()

	mockRepository.On("SaveJson", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(fmt.Errorf("failed to save data")).Once()

	_, err := service.Reader(mockLogger, mockDB, mockFileReader, mockRepository, payload)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to save data")

	mockLogger.AssertExpectations(t)
	mockFileReader.AssertExpectations(t)
	mockRepository.AssertExpectations(t)
}
