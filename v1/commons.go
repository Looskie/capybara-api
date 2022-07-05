package v1

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty" `
	Data    interface{} `json:"data"`
}
