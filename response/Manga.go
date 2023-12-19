package response

type Manga struct {
	Title      string            `json:"title"`
	Thumbnail  string            `json:"thumbnail"`
	Alter      string            `json:"alter"`
	Status     string            `json:"status"`
	Author     string            `json:"author"`
	Ilustrator string            `json:"ilustrator"`
	Grafis     string            `json:"grafis"`
	Scores     string            `json:"scores"`
	Genre      []string          `json:"genre"`
	Synopsis   string            `json:"synopsis"`
	Chapters   []DetailResponses `json:"chapters"`
}

type MangaChapter struct {
	Title     string   `json:"title"`
	Thumbnail string   `json:"thumbnail"`
	Images    []string `json:"images"`
}
