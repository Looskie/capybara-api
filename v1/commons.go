package v1

const (
	NUMBER_OF_IMAGES = 738
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ImageStruct struct {
	URL    string `json:"url"`
	Index  int    `json:"index"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
