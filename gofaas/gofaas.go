package gofaas

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

// GoFaaS is an HTTP Cloud Function
// - JSON Body: { name: string }
func GoFaaS(w http.ResponseWriter, r *http.Request) {
	var rDTO struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&rDTO); err != nil {
		fmt.Fprint(w, "Wrong json format!")
		return
	}
	if rDTO.Name == "" {
		fmt.Fprint(w, "Hello, World!")
		return
	}
	fmt.Fprintf(w, "Hello, %s!", html.EscapeString(rDTO.Name))
}
