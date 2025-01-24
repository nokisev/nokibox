package main

import (
	"log"
	"net/http"
)

// controller home
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Привет из nokibox"))
}

// springboot app starter
func main() {
	mux := http.NewServeMux()
	// обработка controller
	mux.HandleFunc("/", home)

	log.Println("Запуск веб-сервера на http://localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
