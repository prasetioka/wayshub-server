package main

import (
	// "fmt"
	"net/http"

	// "os"
	"wayshub-server/database"
	"wayshub-server/pkg/mysql"
	"wayshub-server/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	// inisiasi database mysql
	mysql.DatabaseInit()

	// untuk running miggrasi
	database.RunMigration()

	// untuk routing
	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	var AllowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	var AllowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"})
	var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})

	// var port = "5000"
	// fmt.Println("server running localhost:" + port)
	// http.ListenAndServe("localhost:"+port, handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))

	// var port = os.Getenv("PORT")
	http.ListenAndServe(":", handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))
}
