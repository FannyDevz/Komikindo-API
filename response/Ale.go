package response

type Ale struct {
	Status     string `json:"status"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
	Nama       string `json:"nama"`
}
