package jsonfeed

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

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
	enc.SetEscapeHTML(false)
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
