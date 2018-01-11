package sdk

// Message represents a standard response.
type Message struct {
	Status  int64  `json:"status"`
	Message string `json:"message"`
}
