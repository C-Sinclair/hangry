package core

type Repository interface {
	Find(id string) Article
	FindAll() Articles
}
