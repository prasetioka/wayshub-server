package routes

import (
  "github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
  AuthRoutes(r)
  ChannelRoutes(r)
  VideoRoutes(r)
  CommentRoutes(r)
  SubscribeRoutes(r)
}