package main

import(
	"net/http"
	"github.com/gorilla/mux"
	"github.com/mycodesmells/mongo-go-api/api"


)

func main(){
	r :=mux.NewRouter()
	r.HandleFunc("/api/items", api.GetUsers).Methods("GET")
	r.HandleFunc("/api/items/{id}", api.CreateUser).Methods("POST")
	r.HandleFunc("/api/items/{id}", api.UpdateUser).Methods("PUT")
	r.HandleFunc("/api/items/{id}", api.DeleteUser).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}

