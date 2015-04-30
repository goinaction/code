// This sample test demonstrates the basic usage of the
// testing package.
package listing01

import (
	"encoding/xml"
	"net/http"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

// Item defines the fields associated with the item tag in
// the buoy RSS document.
type Item struct {
	XMLName     xml.Name `xml:"item"`
	PubDate     string   `xml:"pubDate"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	GUID        string   `xml:"guid"`
	GeoRssPoint string   `xml:"georss:point"`
}

// Image defines the fields associated with the image tag in
// the buoy RSS document.
type Image struct {
	XMLName xml.Name `xml:"image"`
	URL     string   `xml:"url"`
	Title   string   `xml:"title"`
	Link    string   `xml:"link"`
}

// Channel defines the fields associated with the channel tag in
// the buoy RSS document.
type Channel struct {
	XMLName        xml.Name `xml:"channel"`
	Title          string   `xml:"title"`
	Description    string   `xml:"description"`
	Link           string   `xml:"link"`
	PubDate        string   `xml:"pubDate"`
	LastBuildDate  string   `xml:"lastBuildDate"`
	TTL            string   `xml:"ttl"`
	Language       string   `xml:"language"`
	ManagingEditor string   `xml:"managingEditor"`
	WebMaster      string   `xml:"webMaster"`
	Image          Image    `xml:"image"`
	Items          []Item   `xml:"item"`
}

// Document defines the fields associated with the buoy RSS document.
type Document struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
	URI     string
}

// TestDownload tests if download web content is working.
func TestDownload(t *testing.T) {
	URL := "http://www.goinggo.net/feeds/posts/default?alt=rss"

	t.Log("Given the need to test downloading content.")
	{
		resp, err := http.Get(URL)
		if err == nil {
			t.Log("\tShould be able to make the Get call.",
				succeed)
		} else {
			t.Fatal("\tShould be able to make the Get call.",
				failed, err)
		}

		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			t.Log("\tShould receive a \"200\" status code.",
				succeed)
		} else {
			t.Error("\tShould receive a \"200\" status code.",
				failed, resp.StatusCode)
		}

		var d Document
		if err := xml.NewDecoder(resp.Body).Decode(&d); err == nil {
			t.Log("\tShould be able to unmarshal the response.",
				succeed)
		} else {
			t.Fatal("\tShould be able to unmarshal the response.",
				failed, err)
		}

		if len(d.Channel.Items) == 1 {
			t.Log("\tShould have \"1\" item in the feed.",
				succeed)
		} else {
			t.Fatal("\tShould have \"1\" item in the feed.",
				failed, len(d.Channel.Items))
		}
	}
}
