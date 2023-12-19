package response

type RezaMessage struct {
	Message    string `json:"message"`
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
}

type RezaSucess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
