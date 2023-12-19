package response

type Villaneuvo struct {
	Status     string `json:"status"`
	StatusCode string `json:"status_code"`
	Message    string `json:"message"`
	Name       string `json:"name"`
}
