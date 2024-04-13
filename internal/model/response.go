package model

type Response struct {
	Type    string      `json:"type"`
	Version string      `json:"version"`
	Hash    string      `json:"hash"`
	Content interface{} `json:"content"`
}
