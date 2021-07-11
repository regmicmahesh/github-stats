package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type RequestBody struct {
	Repository string `json:"repository"`
}

type ErrorBody struct {
	ErrorMessage string `json:"error"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var p RequestBody

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		msg := ErrorBody{
			ErrorMessage: "Invalid JSON Body",
		}
		bytes, _ := json.Marshal(msg)
		fmt.Fprintln(w, string(bytes))
		return
	}

	processRepository(w, p.Repository)
}

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	http.HandleFunc("/", handler)

	fmt.Println("Server running at 0.0.0.0:"+port);

	http.ListenAndServe(":"+port, nil)

}
