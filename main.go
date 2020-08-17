/*
  Caesar Cipher Solver
  --
  This application has 3 main functions to interact with caesar ciphers:
  - Solve
  - Encrypt
  - Decrypt
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"caesarciphersolver/internal/caesartools"
	"caesarciphersolver/internal/ratelimiter"
)

type payload struct {
	Text  string
	Shift string
}

var dictionary map[string]int

func handleEncrypt(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is supported", http.StatusNotFound)
		return
	}

	// Process request
	decoder := json.NewDecoder(r.Body)
	var request payload
	if err := decoder.Decode(&request); err != nil {
		fmt.Fprintf(w, "Unable to parse request payload | err: %v", err)
		return
	}

	// Encrypt: plaintext + shift => ciphertext
	fmt.Printf("ENCRYPTING: plaintext received | payload = %v\n", request)
	shift, _ := strconv.Atoi(request.Shift)
	var ciphertext string = caesartools.Shift(request.Text, shift)

	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"response": "` + ciphertext + `"}`))
}

func handleDecrypt(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is supported", http.StatusNotFound)
		return
	}

	// Process request
	decoder := json.NewDecoder(r.Body)
	var request payload
	if err := decoder.Decode(&request); err != nil {
		fmt.Fprintf(w, "Unable to parse request payload | err: %v", err)
		return
	}

	// Decrypt: ciphertext + shift => plaintext
	fmt.Printf("DECRYPTING: ciphertext received | payload = %v\n", request)
	shift, _ := strconv.Atoi(request.Shift)
	var plaintext string = caesartools.Shift(request.Text, -1*shift)

	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"response": "` + plaintext + `"}`))
}

func handleSolve(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is supported", http.StatusNotFound)
		return
	}

	// Process request
	decoder := json.NewDecoder(r.Body)
	var request payload
	if err := decoder.Decode(&request); err != nil {
		fmt.Fprintf(w, "Unable to parse request payload | err: %v", err)
		return
	}

	// Solve: ciphertext => plaintext
	fmt.Printf("SOLVING: ciphertext received | payload = %v\n", request)
	var shift int = caesartools.Solve(request.Text, dictionary)
	var plaintext string = caesartools.Shift(request.Text, -1*shift)

	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"response": "` + plaintext + `", "shift": "` + strconv.Itoa(shift) + `"}`))
}

func main() {
	mux := http.NewServeMux()

	// Serve page
	mux.Handle("/", http.FileServer(http.Dir("./website")))

	// Serve API endpoints
	mux.HandleFunc("/encrypt", handleEncrypt)
	mux.HandleFunc("/decrypt", handleDecrypt)
	mux.HandleFunc("/solve", handleSolve)

	// Initialize dictionary map
	dictionary = make(map[string]int)
	caesartools.ParseWords(dictionary)

	// Initialize server with rate limiting
	fmt.Println("Listening on port 8090")
	log.Fatal(http.ListenAndServe(":8090", ratelimiter.Limit(mux)))
}
