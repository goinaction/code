// Copyright 2013. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rss

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"regexp"
)

type (
	// RSSItem defines the fields associated with the item tag in the buoy RSS document.
	RSSItem struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}

	// RSSImage defines the fields associated with the image tag in the buoy RSS document.
	RSSImage struct {
		XMLName xml.Name `xml:"image"`
		Url     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	// RSSChannel defines the fields associated with the channel tag in the buoy RSS document.
	RSSChannel struct {
		XMLName        xml.Name  `xml:"channel"`
		Title          string    `xml:"title"`
		Description    string    `xml:"description"`
		Link           string    `xml:"link"`
		PubDate        string    `xml:"pubDate"`
		LastBuildDate  string    `xml:"lastBuildDate"`
		TTL            string    `xml:"ttl"`
		Language       string    `xml:"language"`
		ManagingEditor string    `xml:"managingEditor"`
		WebMaster      string    `xml:"webMaster"`
		Image          RSSImage  `xml:"image"`
		Item           []RSSItem `xml:"item"`
	}

	// RSSDocument defines the fields associated with the buoy RSS document.
	RSSDocument struct {
		XMLName xml.Name   `xml:"rss"`
		Channel RSSChannel `xml:"channel"`
		Uri     string
	}
)

type (
	// SearchResult contains the result of a search.
	SearchResult struct {
		Field    string
		Document *RSSItem
	}
)

// Retrieve performs a HTTP Get request for the RSS feed and unmarshals the results.
func Retrieve(uri string) (*RSSDocument, error) {
	if uri == "" {
		return nil, fmt.Errorf("No RSS Feed Uri Provided")
	}

	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	defer func() {
		resp.Body.Close()
	}()

	rssDocument := &RSSDocument{}
	err = xml.NewDecoder(resp.Body).Decode(rssDocument)
	if err != nil {
		return nil, err
	}

	rssDocument.Uri = uri

	return rssDocument, err
}

// Search looks at the document for the specified search term.
func Search(rssDocument *RSSDocument, searchTerm string) ([]SearchResult, error) {
	var searchResults []SearchResult

	for index, rssItem := range rssDocument.Channel.Item {
		// Check the title
		matched, err := regexp.MatchString(searchTerm, rssItem.Title)
		if err != nil {
			return nil, err
		}

		if matched {
			searchResults = append(searchResults, SearchResult{
				Field:    "Title",
				Document: &rssDocument.Channel.Item[index],
			})
		}

		// Check the description
		matched, err = regexp.MatchString(searchTerm, rssItem.Description)
		if err != nil {
			return nil, err
		}

		if matched {
			searchResults = append(searchResults, SearchResult{
				Field:    "Description",
				Document: &rssDocument.Channel.Item[index],
			})
		}
	}

	return searchResults, nil
}
