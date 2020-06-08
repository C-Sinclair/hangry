package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/c-sinclair/restic/pkg/core"
)

type Service struct {
	repo   core.Repository
	Router http.Handler
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := core.Articles{
		core.Article{
			Title:   "A great test title",
			Desc:    "Test description",
			Content: "Hello World",
		},
	}
	fmt.Println("All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint hit")
}

func New(repo core.Repository) Service {
	service := Service{
		repo: repo,
	}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/articles", allArticles).Methods("GET")

	service.Router = UseMiddleware(router)

	return service
}
