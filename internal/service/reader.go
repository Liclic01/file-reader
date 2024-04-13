package service

import (
	"database/sql"
	"encoding/json"
	"github.com/heroiclabs/nakama-common/runtime"
	"zepto-lab.com/file-reader/internal/model"
	"zepto-lab.com/file-reader/internal/repository"
	"zepto-lab.com/file-reader/internal/util"
)

func Reader(logger runtime.Logger, db *sql.DB, fileReader FileReader, repository repository.JsonRepository, payload string) (string, error) {
	var request model.Request
	if err := util.UnmarshalAndSetDefaults(payload, &request); err != nil {
		logger.Error("Failed to unmarshal request payload: %v", err)
		return "", runtime.NewError("Failed to unmarshal request payload", util.INVALID_ARGUMENT)
	}
	fileContent, err := fileReader.ReadFileContent("/path/inside/container", request.Type, request.Version)
	if err != nil {
		logger.Error("Failed to read file content: %v", err)
		return "", runtime.NewError("Failed to unmarshal request payload", util.UNAVAILABLE)
	}

	contentHash := util.CalculateHash(fileContent)

	err = repository.SaveJson(db, "file_content", request.Type+"_"+request.Version, map[string]interface{}{"content": fileContent})
	if err != nil {
		logger.Error("Failed to save data to repository: %v", err)
		return "", runtime.NewError("failed to save data to repository", util.UNAVAILABLE)
	}

	response := model.Response{
		Type:    request.Type,
		Version: request.Version,
		Hash:    contentHash,
		Content: fileContent,
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		logger.Error("Failed to marshal response: %v", err)
		return "", runtime.NewError("Failed to marshal response", util.UNAVAILABLE)
	}

	return string(responseJSON), nil
}
