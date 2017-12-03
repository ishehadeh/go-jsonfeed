package jsonfeed

import "time"

//Author Basic information on the feed's author
type Author struct {
	Name   string `json:"name,omitempty"`
	URL    string `json:"url,omitempty"`
	Avatar string `json:"avatar,omitempty"`
}

//Hub An endpoint that can be used to subscibe to the real-time notifications from the publisher of a feed
type Hub struct {
	Type string `json:"type,omitempty"`
	URL  string `json:"url,omitempty"`
}

//Attachment a list of resources related to a file
type Attachment struct {
	URL               string `json:"url,omitempty" feed:"required"`
	MimeType          string `json:"mime_type,omitempty" feed:"required"`
	Title             string `json:"title,omitempty" feed:"optional"`
	SizeInBytes       int64  `json:"size_in_bytes,omitempty" feed:"optional"`
	DurationInSeconds int64  `json:"duration_in_seconds,omitempty" feed:"optional"`
}

//Item a single item in a feed
type Item struct {
	ID            string        `json:"id,omitempty" feed:"required"`
	URL           string        `json:"url,omitempty" feed:"optional"`
	ExternalURL   string        `json:"external_url,omitempty" feed:"optional"`
	Title         string        `json:"title,omitempty" feed:"optional"`
	ContentHTML   string        `json:"content_html,omitempty" feed:"optional"`
	ContentText   string        `json:"content_text,omitempty" feed:"optional"`
	Summary       string        `json:"summary,omitempty" feed:"optional"`
	Image         string        `json:"image,omitempty" feed:"optional"`
	BannerImage   string        `json:"banner_image,omitempty" feed:"optional"`
	DatePublished time.Time     `json:"date_published,omitempty" feed:"optional"`
	DateModified  time.Time     `json:"date_modified,omitempty" feed:"optional"`
	Author        Author        `json:"author,omitempty" feed:"optional"`
	Tags          []string      `json:"tags,omitempty" feed:"optional"`
	Attachments   []*Attachment `json:"attachments,omitempty" feed:"optional"`
}

//JSONFeed The structure of a JSON Feed
type JSONFeed struct {
	Version     string   `json:"version,omitempty" feed:"required"`
	Title       string   `json:"title,omitempty" feed:"required"`
	HomePageURL string   `json:"home_page_url,omitempty" feed:"optional"`
	FeedURL     string   `json:"feed_url,omitempty" feed:"optional"`
	Description string   `json:"description,omitempty" feed:"optional"`
	UserComment string   `json:"user_comment,omitempty" feed:"optional"`
	NextURL     string   `json:"next_url,omitempty" feed:"optional"`
	Icon        string   `json:"icon,omitempty" feed:"optional"`
	Favicon     string   `json:"favicon,omitempty" feed:"optional"`
	Author      Author   `json:"author,omitempty" feed:"optional"`
	Items       []*Item  `json:"items,omitempty" feed:"optional"`
	Expired     bool     `json:"expired,omitempty" feed:"optional"`
	Hubs        []string `json:"hubs,omitempty" feed:"optional"`
}
