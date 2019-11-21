package gofaas

import (
	"net/http/httptest"
	"strings"
	"testing"
)

// TestGoFaaS test GoFaaS function
func TestGoFaaS(t *testing.T) {
	tests := []struct {
		body string
		want string
	}{
		{body: ``, want: "Wrong json format!"},
		{body: `{"name": ""}`, want: "Hello, World!"},
		{body: `{"name": "CTG"}`, want: "Hello, CTG!"},
	}

	for _, test := range tests {
		req := httptest.NewRequest("GET", "/", strings.NewReader(test.body))
		req.Header.Add("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		GoFaaS(rr, req)

		if got := rr.Body.String(); got != test.want {
			t.Errorf("Failed > Hello(\"%v\"): %v => %v", test.body, test.want, got)
		} else {
			t.Logf("Success > Hello(\"%v\"): %v => %v", test.body, test.want, got)
		}
	}
}
