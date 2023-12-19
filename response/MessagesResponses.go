package response

type Error struct {
	Message    string `json:"message"`
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
}

type Success struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PingSuccess struct {
	Status     string `json:"status"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}
