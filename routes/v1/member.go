package v1

import (
	"go-api/controllers"

	"github.com/gorilla/mux"
)

func MemberRoute(route *mux.Router) {
	route.HandleFunc("/member", controllers.GetAllMember).Methods("GET")
	route.HandleFunc("/member/{id}", controllers.GetOneMember).Methods("GET")
	route.HandleFunc("/member", controllers.PostMember).Methods("POST")
	route.HandleFunc("/member/{id}", controllers.PutMember).Methods("PUT")
	route.HandleFunc("/member/{id}", controllers.DelMember).Methods("DELETE")
}
