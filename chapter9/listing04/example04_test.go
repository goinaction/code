// Sample test to show how to test the execution of an internal endpoint.
package main_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	ex "github.com/goinaction/code/chapter9/listing04"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func init() {
	ex.Routes()
}

// TestSendJSON testing the sendjson internal endpoint.
func TestSendJSON(t *testing.T) {
	t.Log("Given the need to test the SendJSON endpoint.")
	{
		r, _ := http.NewRequest("GET", "/sendjson", nil)
		w := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(w, r)

		if w.Code != 200 {
			t.Fatalf("\tShould received a status code of 200 for the response. Received[%d] %s", w.Code, ballotX)
		}
		t.Log("\tShould received a status code of 200 for the response.", checkMark)

		u := struct {
			Name  string
			Email string
		}{}
		if err := json.NewDecoder(w.Body).Decode(&u); err != nil {
			t.Fatal("\tShould be able to unmarshal the response.", ballotX)
		}
		t.Log("\tShould be able to unmarshal the response.", checkMark)

		if u.Name == "Bill" {
			t.Log("\tShould have \"Bill\" for Name in the response.", checkMark)
		} else {
			t.Error("\tShould have \"Bill\" for Name in the response.", ballotX, u.Name)
		}

		if u.Email == "bill@ardanstudios.com" {
			t.Log("\tShould have \"bill@ardanstudios.com\" for Email in the response.", checkMark)
		} else {
			t.Error("\tShould have \"bill@ardanstudios.com\" for Email in the response.", ballotX, u.Email)
		}
	}
}
