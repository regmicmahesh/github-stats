package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestBody struct {
	Repository string `json:"repository"`
}

type ErrorBody struct{
	ErrorMessage string `json:"error"`
}


func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var p RequestBody

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil{
		msg := ErrorBody{
			ErrorMessage: "Invalid JSON Body",
		}
		bytes, _ := json.Marshal(msg)
		fmt.Fprintln(w, string(bytes))
		return;
	}

	fmt.Println(p)

	//bytes, _ := json.Marshal(p)
	processRepository(w, p.Repository)

}

func main() {

	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)


}
