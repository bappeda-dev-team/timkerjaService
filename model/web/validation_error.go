package web

type ValidationError struct {
	Message string            `json:"message"`
	Fields  map[string]string `json:"fields"`
}

func (e *ValidationError) Error() string {
	return e.Message
}
