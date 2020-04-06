package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"caesartools"
)

func parsePlaintext(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is supported", http.StatusNotFound)
		return
	}

	// Request
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Printf("plaintext received | r.PostFrom = %v\n", r.PostForm)
	text := r.FormValue("text")
	shift, _ := strconv.Atoi(r.FormValue("shift"))
	var ciphertext string = caesartools.Encrypt(text, shift)

	// Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"ciphertext": "` + ciphertext + `"}`))
}

func parseCiphertext(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is supported", http.StatusNotFound)
		return
	}

	// Request
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Printf("ciphertext received | r.PostFrom = %v\n", r.PostForm)
	mode := r.FormValue("action")

	var plaintext string
	switch mode {
	case "DECRYPT":
		text := r.FormValue("text")
		shift, _ := strconv.Atoi(r.FormValue("shift"))
		plaintext = caesartools.Decrypt(text, shift)
	case "SOLVE":
		text := r.FormValue("text")
		plaintext = caesartools.Solve(text)
	default:
		plaintext = ""
	}

	// Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"plaintext": "` + plaintext + `"}`))
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./website")))
	http.HandleFunc("/plaintext", parsePlaintext)
	http.HandleFunc("/ciphertext", parseCiphertext)

	log.Fatal(http.ListenAndServe(":8090", nil))
}
