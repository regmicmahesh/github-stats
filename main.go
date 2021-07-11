package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/patrickmn/go-cache"
)

type RequestBody struct {
	Repository string `json:"repository"`
}

type ErrorBody struct {
	ErrorMessage string `json:"error"`
}

var c *cache.Cache

func handler(w http.ResponseWriter, r *http.Request) {
	languageCountMap = map[string]int{}
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

	if resp, found := c.Get(p.Repository); found {
		jsonString := resp.(string)
		fmt.Println("found cache")
		fmt.Fprintln(w, jsonString)
		return;
	}
	processRepository(w, p.Repository)
}

func main() {

	c = cache.New(5*time.Minute, 10*time.Minute)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	http.HandleFunc("/", handler)

	fmt.Println("Server running at 0.0.0.0:" + port)

	http.ListenAndServe(":"+port, nil)

}
