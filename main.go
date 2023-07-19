package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	emailverifier "github.com/moonman369/Go-Email-Verifier/email-verifier"
	"github.com/rs/cors"
)

type Domain struct {
	Name string `json:"name"`
}

func verifyDomain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var domain Domain
	_ = json.NewDecoder(r.Body).Decode(&domain)
	if domain.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res := emailverifier.CheckDomain(domain.Name)
	json.NewEncoder(w).Encode(res)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not allowed by the server", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "<html><body><h1>Welcome To Go Email Verifier<h1><br><br></body></html>")
}

func main() {
	// NATIVE APPLICATION CODE
	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Printf("domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord\n")

	// for scanner.Scan() {
	// 	emailverifier.CheckDomain(scanner.Text())
	// }

	// if err := scanner.Err(); err != nil {
	// 	log.Fatal("Error: Invalid input: %v\n", err)
	// }

	r := mux.NewRouter()

	r.HandleFunc("/verify", verifyDomain).Methods("POST")
	r.HandleFunc("/", helloHandler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://devfoliomoonman369.netlify.app"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe("0.0.0.0:8080", handler); err != nil {
		log.Fatal(err)
	}
}
