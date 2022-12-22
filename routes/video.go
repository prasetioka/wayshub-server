package routes

import (
	"wayshub-server/handlers"
	"wayshub-server/pkg/middleware"
	"wayshub-server/pkg/mysql"
	"wayshub-server/repositories"

	"github.com/gorilla/mux"
)

func VideoRoutes(r *mux.Router) {
	videoRepository := repositories.RepositoryVideo(mysql.DB)
	h := handlers.HandlerVideo(videoRepository)

	r.HandleFunc("/videos", (h.FindVideos)).Methods("GET")
	r.HandleFunc("/video/{id}", (h.GetVideo)).Methods("GET")

	r.HandleFunc("/video", middleware.Auth(middleware.UploadVideo(middleware.UploadThumbnail(h.CreateVideo)))).Methods("POST")
	r.HandleFunc("/video/{id}", middleware.Auth(middleware.UploadVideo(middleware.UploadThumbnail(h.UpdateVideo)))).Methods("PATCH")
	
	r.HandleFunc("/video/{id}", middleware.Auth(h.DeleteVideo)).Methods("DELETE")

	r.HandleFunc("/myvideo", middleware.Auth(h.FindVideosByChannelId)).Methods("GET")
	r.HandleFunc("/FindMyVideos", middleware.Auth(h.FindMyVideos)).Methods("GET")

	r.HandleFunc("/UpdateViews/{id}", middleware.Auth(h.UpdateViews)).Methods("PATCH")
}