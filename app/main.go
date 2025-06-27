package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've hit %s\n", r.URL.Path)
}

func main() {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	}

	http.HandleFunc("/", handler)

	fmt.Println("Server is running at http://localhost" + os.Getenv("SERVER_ADDR"))
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
