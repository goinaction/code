package test_demo

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestDownload(t *testing.T) {
	url := "http://www.google.com"
	statusOK := 200

	t.Log("Given the need to test connect .")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", url, statusOK)
		{
			response, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\tShould be able to make Get call.", ballotX, err)
			}
			t.Log("\t\tShould be able to make the Get call.", checkMark)

			defer response.Body.Close()

			if response.StatusCode == statusOK {
				t.Logf("\t\tShould receive a \"%d\" status. %v %v", statusOK, checkMark, response.StatusCode)
			} else {
				t.Errorf("\t\tShould receive a \"%d\" status. %v %v", statusOK, ballotX, response.StatusCode)
			}
		}
	}
}
