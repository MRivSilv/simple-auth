package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"simple-auth/handlers"
	"simple-auth/store"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("failed to open database: ", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("failed to ping database: ", err)
	}

	userStore := store.NewStore(db)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /register-user", handlers.RegisterUser(userStore))
	mux.HandleFunc("POST /login", handlers.Login(userStore))
	mux.HandleFunc("GET /protected", handlers.RequireAuth(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You are in"))
	}))
	mux.HandleFunc("DELETE /account", handlers.RequireAuth(handlers.DeleteUser(userStore)))

	log.Println("RUNNING on localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
