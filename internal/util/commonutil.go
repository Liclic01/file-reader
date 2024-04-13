package util

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"zepto-lab.com/file-reader/internal/model"
)

func UnmarshalAndSetDefaults(payload string, request *model.Request) error {
	if err := json.Unmarshal([]byte(payload), request); err != nil {
		return err
	}

	request.SetDefaultValues()

	return nil
}

func CalculateHash(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes)
}
