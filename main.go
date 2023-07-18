package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	emailverifier "github.com/moonman369/Go-Email-Verifier/email-verifier"
)

type Domain struct {
	Path string `json:"path"`
}

// type VerificationResult struct {
// 	Domain      *Domain `json:"domain"`
// 	HasMX       bool    `json:"hasMX"`
// 	HasSPF      bool    `json:"hasSPF"`
// 	SPFRecord   string  `json:"spfRecord"`
// 	HasDMARC    bool    `json:"hasDMARC"`
// 	DMARCRecord string  `json:"dmarcRecord"`
// }

func verifyDomain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var domain Domain
	_ = json.NewDecoder(r.Body).Decode(&domain)
	res := emailverifier.CheckDomain(domain.Path)
	json.NewEncoder(w).Encode(res)
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
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe("0.0.0.0:8080", r); err != nil {
		log.Fatal(err)
	}
}
