package rss

import (
	"encoding/xml"
	"errors"
	"net/http"
	"regexp"
)

type (
	// Item defines the fields associated with the item tag in the rss document.
	Item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}

	// Image defines the fields associated with the image tag in the rss document.
	Image struct {
		XMLName xml.Name `xml:"image"`
		Url     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	// Channel defines the fields associated with the channel tag in the rss document.
	Channel struct {
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
		Item           []Item   `xml:"item"`
	}

	// Document defines the fields associated with the rss document.
	Document struct {
		XMLName xml.Name `xml:"rss"`
		Channel Channel  `xml:"channel"`
		Uri     string
	}
)

type (
	// SearchResult contains the result of a search.
	SearchResult struct {
		Field    string
		Document Item
	}
)

// Retrieve performs a HTTP Get request for the rss feed and unmarshals the results.
func Retrieve(uri string, document *Document) error {
	if uri == "" {
		return errors.New("No RSS Feed Uri Provided")
	}

	// Retrieve the rss feed document from the web.
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}

	// Close the response once we return from the function.
	defer resp.Body.Close()

	// Unmarshal the document into our struct type.
	err = xml.NewDecoder(resp.Body).Decode(document)
	if err != nil {
		return err
	}

	// Save the uri we used to retrieve this document.
	document.Uri = uri

	return err
}

// Search looks at the document for the specified search term.
func Search(document Document, searchTerm string) ([]SearchResult, error) {
	var searchResults []SearchResult

	for _, item := range document.Channel.Item {
		// Check the title for the search term.
		matched, err := regexp.MatchString(searchTerm, item.Title)
		if err != nil {
			return nil, err
		}

		// If we found as match save the result.
		if matched {
			searchResults = append(searchResults, SearchResult{
				Field:    "Title",
				Document: item,
			})
		}

		// Check the description for the search term.
		matched, err = regexp.MatchString(searchTerm, item.Description)
		if err != nil {
			return nil, err
		}

		// If we found a match so save the result.
		if matched {
			searchResults = append(searchResults, SearchResult{
				Field:    "Description",
				Document: item,
			})
		}
	}

	return searchResults, nil
}
