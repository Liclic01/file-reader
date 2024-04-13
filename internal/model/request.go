package model

type Request struct {
	Type    string `json:"type,omitempty"`
	Version string `json:"version,omitempty"`
	Hash    string `json:"hash,omitempty"`
}

func (r *Request) SetDefaultValues() {
	if r.Type == "" {
		r.Type = "core"
	}
	if r.Version == "" {
		r.Version = "1.0.0"
	}
	if r.Hash == "" {
		r.Hash = "null"
	}
}
