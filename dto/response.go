package dto

type Response struct {
	Data  any `json:"data,omitempty"`
	Error any `json:"error,omitempty"`
}
