package routes

import (
	v1 "go-api/routes/v1"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func IndexRoutes() {
	r := mux.NewRouter()

	v1.MemberRoute(r.PathPrefix("/v1").Subrouter())

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(headers, methods, origins)(r))

}
