package types

// Articles is a single article of news
type Articles struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// News is the list of articles
type News struct {
	Status       string     `json:"status"`
	TotalResults int        `json:"totalResults"`
	Articles     []Articles `json:"articles"`
}
