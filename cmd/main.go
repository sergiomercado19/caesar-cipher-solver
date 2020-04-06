package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	fmt.Fprintf(w, "plaintext received | r.PostFrom = %v\n", r.PostForm)
	text := r.FormValue("text")
	shift, _ := strconv.Atoi(r.FormValue("shift"))
	var ciphertext string = encrypt(text, shift)

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
	fmt.Fprintf(w, "ciphertext received | r.PostFrom = %v\n", r.PostForm)
	mode := r.FormValue("action")

	var plaintext string
	switch mode {
	case "DECRYPT":
		text := r.FormValue("text")
		shift, _ := strconv.Atoi(r.FormValue("shift"))
		plaintext = decrypt(text, shift)
	case "SOLVE":
		text := r.FormValue("text")
		plaintext = solve(text)
	default:
		plaintext = ""
	}

	// Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"plaintext": "` + plaintext + `"}`))
}

func main() {
	http.HandleFunc("/plaintext", parsePlaintext)
	http.HandleFunc("/ciphertext", parseCiphertext)

	log.Fatal(http.ListenAndServe(":8090", http.FileServer(http.Dir("./website"))))
}
