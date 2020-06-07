package core

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"description"`
	Content string `json:"content"`
}
type Articles []Article
