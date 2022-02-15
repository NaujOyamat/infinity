package models

// RequestResult encapsulates result values of a request
type RequestResult struct {
	Ok      bool        `json:"ok"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}
