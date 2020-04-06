package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"caesartools"
)

type payload struct {
	Text  string
	Shift string
}

func handleEncrypt(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is supported", http.StatusNotFound)
		return
	}

	// Request
	decoder := json.NewDecoder(r.Body)
	var request payload
	if err := decoder.Decode(&request); err != nil {
		fmt.Fprintf(w, "Unable to parse request payload | err: %v", err)
		return
	}
	fmt.Printf("ENCRYPTING: plaintext received | payload = %v\n", request)
	shift, _ := strconv.Atoi(request.Shift)
	var ciphertext string = caesartools.Encrypt(request.Text, shift)

	// Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"response": "` + ciphertext + `"}`))
}

func handleDecrypt(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is supported", http.StatusNotFound)
		return
	}

	// Request
	decoder := json.NewDecoder(r.Body)
	var request payload
	if err := decoder.Decode(&request); err != nil {
		fmt.Fprintf(w, "Unable to parse request payload | err: %v", err)
		return
	}
	fmt.Printf("DECRYPTING: plaintext received | payload = %v\n", request)
	shift, _ := strconv.Atoi(request.Shift)
	var plaintext string = caesartools.Decrypt(request.Text, shift)

	// Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"response": "` + plaintext + `"}`))
}

func handleSolve(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is supported", http.StatusNotFound)
		return
	}

	// Request
	decoder := json.NewDecoder(r.Body)
	var request payload
	if err := decoder.Decode(&request); err != nil {
		fmt.Fprintf(w, "Unable to parse request payload | err: %v", err)
		return
	}
	fmt.Printf("SOLVING: plaintext received | payload = %v\n", request)
	var plaintext string = caesartools.Solve(request.Text)

	// Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"response": "` + plaintext + `"}`))
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./website")))
	http.HandleFunc("/encrypt", handleEncrypt)
	http.HandleFunc("/decrypt", handleDecrypt)
	http.HandleFunc("/solve", handleSolve)

	log.Fatal(http.ListenAndServe(":8090", nil))
}
