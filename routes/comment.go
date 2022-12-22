package routes

import (
	"wayshub-server/handlers"
	"wayshub-server/pkg/middleware"
	"wayshub-server/pkg/mysql"
	"wayshub-server/repositories"

	"github.com/gorilla/mux"
)

func CommentRoutes(r *mux.Router) {
	commentRepository := repositories.RepositoryComment(mysql.DB)
	h := handlers.HandlerComment(commentRepository)

	r.HandleFunc("/comments", (h.FindComments)).Methods("GET")
	r.HandleFunc("/comment/{id}", (h.GetComment)).Methods("GET")
	r.HandleFunc("/comment/{id}", middleware.Auth(h.CreateComment)).Methods("POST")
	r.HandleFunc("/comment/{id}", middleware.Auth(h.UpdateComment)).Methods("PATCH")
	r.HandleFunc("/comment/{id}", middleware.Auth(h.DeleteComment)).Methods("DELETE")
}