package response

type MenuResponses struct {
	Menu     []DetailResponses `json:"menu"`
	Populars []PopularManga    `json:"populars"`
	Latest   []LatesUpdate     `json:"latest"`
}

type DetailResponses struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	Endpoint string `json:"endpoint"`
}

type LatesUpdate struct {
	Name      string `json:"name"`
	Url       string `json:"url"`
	Endpoint  string `json:"endpoint"`
	Thumbnail string `json:"thumbnail"`
}

type ListManga struct {
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
	Url       string `json:"url"`
	Endpoint  string `json:"endpoint"`
}

type ListResponsesManga struct {
	Manga      []ListManga       `json:"manga"`
	Pagination []DetailResponses `json:"pagination"`
}

type PopularManga struct {
	Name        string            `json:"name"`
	Thumbnail   string            `json:"thumbnail"`
	Url         string            `json:"url"`
	Endpoint    string            `json:"endpoint"`
	LastUpload  string            `json:"last_upload"`
	LastChapter []DetailResponses `json:"last_chapter"`
}

type SearchResponses struct {
	Manga      []ListManga       `json:"manga"`
	Pagination []DetailResponses `json:"pagination"`
}
