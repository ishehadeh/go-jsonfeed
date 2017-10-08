package jsonfeed

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"
	"strconv"
	"time"

	"reflect"
)

func (jf *JSONFeed) Read(r io.Reader) error {
	dec := json.NewDecoder(r)
	dec.Decode(&jf)
	missing, err := jf.GetMissing()
	if err != nil {
		return err
	}
	if len(missing) != 0 {
		return fmt.Errorf("Feed missing field(s): %v", missing)
	}
	return nil
}
func (jf *JSONFeed) Write(w io.Writer) error {
	missing, err := jf.GetMissing()
	if err != nil {
		return err
	}
	if len(missing) > 0 {
		return fmt.Errorf("Feed missing field(s): %v", missing)
	}
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(true)
	return enc.Encode(jf)
}

//getMissingRecursive Checks if struct has all required fields recursively
func getMissingRecursive(s interface{}) (missing []string, err error) {
	st := reflect.ValueOf(s)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		if feed, ok := st.Type().Field(i).Tag.Lookup("feed"); ok {
			if field.Type().Kind() == reflect.Struct {
				m, err := getMissingRecursive(field)
				if err != nil {
					return missing, err
				}
				missing = append(missing, m...)
			} else {
				if feed == "required" && field == reflect.Zero(field.Type()).Interface() {
					missing = append(missing, st.Type().Field(i).Name)
				}
			}
		}
	}
	return
}

//SetVersion set the version JSON Feed to use
func (jf *JSONFeed) SetVersion(v int) {
	jf.Version = "https://jsonfeed.org/version/" + strconv.Itoa(v)
}

//GetMissing checks if the feed has all the required fields
func (jf JSONFeed) GetMissing() (missing []string, err error) {
	return getMissingRecursive(jf)
}

func (jf *JSONFeed) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	jf.Write(rw)
}

//OpenFeed read from url into a JSONFeed
func OpenFeed(url string) (feed *JSONFeed, err error) {
	p, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	feed = &JSONFeed{}
	err = feed.Read(p.Body)
	if err != nil {
		return nil, err
	}
	p.Body.Close()
	return feed, nil
}

//CreateJSONFeed creates a new json feed (v1)
func CreateJSONFeed(title string, description string) (feed *JSONFeed) {
	feed = &JSONFeed{
		Title:       title,
		Description: description,
	}
	feed.SetVersion(1)
	return
}

//PublishText publishes a new plain text item to a JSON Feed, the publish date is time.Now()
func (jf *JSONFeed) PublishText(id string, title string, text string, attachments ...*Attachment) (i *Item) {
	jf.Items = append(jf.Items, &Item{
		ID:            id,
		Title:         title,
		ContentText:   text,
		Attachments:   attachments,
		DatePublished: time.Now().Format("2006/01/02|15:04:05"),
	})
	return jf.Items[len(jf.Items)-1]
}

//PublishHTML publishes a new html text item to a JSON Feed, the publish date is time.Now()
func (jf *JSONFeed) PublishHTML(id string, title string, textHTML string, attachments ...*Attachment) (i *Item) {
	jf.Items = append(jf.Items, &Item{
		ID:            id,
		Title:         title,
		ContentHTML:   textHTML,
		Attachments:   attachments,
		DatePublished: time.Now().Format("2006/01/02|15:04:05"),
	})
	return jf.Items[len(jf.Items)-1]
}

//NewImage creates a new image attachment
func NewImage(title string, url string) (at *Attachment) {
	imgType := path.Ext(url)
	return &Attachment{
		Title:    title,
		URL:      url,
		MimeType: "image/" + imgType[1:],
	}
}

//NewVideo creates a new video attachment
func NewVideo(title string, url string) (at *Attachment) {
	vidType := path.Ext(url)
	return &Attachment{
		Title:    title,
		URL:      url,
		MimeType: "video/" + vidType[1:],
	}
}

//NewAudio creates a new adio attachment
func NewAudio(title string, url string) (at *Attachment) {
	audType := path.Ext(url)
	return &Attachment{
		Title:    title,
		URL:      url,
		MimeType: "audio/" + audType[1:],
	}
}
