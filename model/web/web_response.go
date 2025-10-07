package web

// @Description Web Response
type WebResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}
