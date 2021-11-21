// This sample program demonstrates how to decode a JSON response
// using the json package and NewDecoder function.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type (
	// gResult maps to the result document received from the search.
	//gResult struct {
	//	GsearchResultClass string `json:"GsearchResultClass"`
	//	UnescapedURL       string `json:"unescapedUrl"`
	//	URL                string `json:"url"`
	//	VisibleURL         string `json:"visibleUrl"`
	//	CacheURL           string `json:"cacheUrl"`
	//	Title              string `json:"title"`
	//	TitleNoFormatting  string `json:"titleNoFormatting"`
	//	Content            string `json:"content"`
	//}

	// gResponse contains the top level document.
	gResponse struct {
		Error struct {
			Code   int64 `json:"code"`
			Errors []struct {
				Domain  string `json:"domain"`
				Message string `json:"message"`
				Reason  string `json:"reason"`
			} `json:"errors"`
			Message string `json:"message"`
			Status  string `json:"status"`
		} `json:"error"`
	}
)

func main() {
	uri := "https://www.googleapis.com/customsearch/v1?q=golang"

	// Issue the search against Google.
	resp, err := http.Get(uri)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("Code: %d\n", resp.StatusCode)
	//fmt.Printf("Body: %s\n", body)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// Decode the JSON response into our struct type.
	var gr gResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println("Non-Pretty Format: \n", gr)

	// Marshal the struct type into a pretty print
	// version of the JSON document.
	pretty, err := json.MarshalIndent(gr, "", "    ")
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println("\nPretty Format: \n", string(pretty))
}
