package routes

import (
	"api-go-rest/controllers"
	"api-go-rest/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("GET")
	r.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("POST")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaPersonalidade).Methods("DELETE")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaPersonalidade).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
