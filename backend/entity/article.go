package entity

type Article struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Tags      []*Tag `json:"tags"`
	Viewers   uint   `json:"viewer"`
	Favorites uint   `json:"favorite"`
}
