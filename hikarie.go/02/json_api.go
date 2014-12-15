package main

import (
	"fmt"
	"strconv"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		name := params.Get("name")
		age, err := strconv.Atoi(params.Get("age"))
		if err != nil {
			fmt.Errorf("age must be a integer value")
			return
		}
		usr := &User{name, age}

		jsn, err := json.Marshal(usr)
		if err != nil {
			fmt.Errorf("could not marshal")
			return
		}
		ioutil.WriteFile("./data.txt", jsn, 0644)

		fmt.Printf("WRITE: %s\n", string(jsn))

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", string(jsn))
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		txt, err := ioutil.ReadFile("./data.txt")
		if err != nil {
			fmt.Errorf("%s\n", err)
			return
		}

		fmt.Printf("READ : %s\n", txt)

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", txt)
	})
	http.ListenAndServe(":4000", nil)

}

