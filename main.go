package main

import ( 
	"log"
	"net/http"

	"simple-auth/handlers"
	"simple-auth/store"

	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()
	userStore := store.NewUserStore()

	mux := http.NewServeMux()
	mux.HandleFunc("POST /register", handlers.Register(userStore))
	mux.HandleFunc("POST /login", handlers.Login(userStore))
	mux.HandleFunc("GET /protected", handlers.RequireAuth(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You are in"))
	}))

	log.Println("RUNNING on localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", mux))
}