package response

type IvanError struct {
	Message    string `json:"message"`
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
}

type IvanSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

