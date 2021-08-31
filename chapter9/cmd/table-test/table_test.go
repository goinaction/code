package table_test

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestDownload(t *testing.T) {

	var urls = []struct {
		url        string
		statusCode int
	}{
		{
			"http://www.google.com",
			http.StatusOK,
		},
		{
			"http://rss.cnn.com/rss/cnn_topstbadurl.rss",
			http.StatusNotFound,
		},
	}
	//url := "http://www.google.com"
	//
	//statusOK := 200

	t.Log("Given the need to test connect .")
	{
		for _, u := range urls {
			t.Logf("\tWhen checking \"%s\" for status code \"%d\"", u.url, u.statusCode)
			{
				response, err := http.Get(u.url)
				if err != nil {
					t.Fatal("\t\tShould be able to make Get call.", ballotX, err)
				}
				t.Log("\t\tShould be able to make the Get call.", checkMark)

				defer response.Body.Close()

				if response.StatusCode == u.statusCode {
					t.Logf("\t\tShould receive a \"%d\" status. %v %v", u.statusCode, checkMark, response.StatusCode)
				} else {
					t.Errorf("\t\tShould receive a \"%d\" status. %v %v", u.statusCode, ballotX, response.StatusCode)
				}
			}

		}
	}
}
