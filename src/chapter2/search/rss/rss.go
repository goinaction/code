package rss

import (
	"encoding/xml"
	"errors"
	"log"
	"net/http"
	"regexp"

	"github.com/goinaction/code/src/chapter2/search/feeds"
	"github.com/goinaction/code/src/chapter2/search/find"
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
	}
)

type (
	// Implements the searcher interface.
	Search struct {
		Site feeds.Site
	}
)

// Match looks at the document for the specified search term.
func (s *Search) Match(searchTerm string) ([]find.Result, error) {
	var results []find.Result

	log.Printf("Search Feed Type[%s] Site[%s] For Uri[%s]\n", s.Site.Type, s.Site.Name, s.Site.Uri)

	// Retrieve the data to search.
	document, err := s.retrieve()
	if err != nil {
		return nil, err
	}

	for _, item := range document.Channel.Item {
		// Check the title for the search term.
		matched, err := regexp.MatchString(searchTerm, item.Title)
		if err != nil {
			return nil, err
		}

		// If we found a match save the result.
		if matched {
			results = append(results, find.Result{
				Field:   "Title",
				Content: item.Title,
			})
		}

		// Check the description for the search term.
		matched, err = regexp.MatchString(searchTerm, item.Description)
		if err != nil {
			return nil, err
		}

		// If we found a match save the result.
		if matched {
			results = append(results, find.Result{
				Field:   "Description",
				Content: item.Description,
			})
		}
	}

	return results, nil
}

// retrieve performs a HTTP Get request for the rss feed and unmarshals the results.
func (s *Search) retrieve() (*Document, error) {
	if s.Site.Uri == "" {
		return nil, errors.New("No RSS Feed Uri Provided")
	}

	// Retrieve the rss feed document from the web.
	resp, err := http.Get(s.Site.Uri)
	if err != nil {
		return nil, err
	}

	// Close the response once we return from the function.
	defer resp.Body.Close()

	// Unmarshal the rss feed document into our struct type.
	var document Document
	err = xml.NewDecoder(resp.Body).Decode(&document)
	if err != nil {
		return nil, err
	}

	return &document, nil
}
