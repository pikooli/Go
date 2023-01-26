package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Founder struct {
	Name    string `json:"name"`
	Age     uint32 `json:"age"`
	Email   string `json:"email"`
	Company string `json:"company"`
}

var founders []Founder

func greetingsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Greeting from Go Server 👋")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var founder Founder
	json.NewDecoder(r.Body).Decode(&founder)
	founder.Age = founder.Age * 2
	founders = append(founders, founder)

	json.NewEncoder(w).Encode(founders)
}

func main() {
	r := chi.NewRouter()
	founders = append(founders, Founder{Name: "Mehul", Age: 23, Email: "random@random.com", Company: "BharatX"})

	r.Get("/", greetingsHandler)
	r.Post("/form", formHandler)
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
