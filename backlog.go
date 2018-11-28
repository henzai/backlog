package backlog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	baseURLf = "https://%v.backlog.jp/api/v2/%v?apiKey=%v"
)

type BackLog struct {
	baseURL string
	key     string
	space   string
}

func New(space, key string) *BackLog {
	baseURL := fmt.Sprintf(baseURLf, space, "%v", key)
	backLog := BackLog{baseURL, key, space}
	return &backLog
}

func (b *BackLog) GetBaseURL() string {
	return b.baseURL
}

func (b *BackLog) GetKey() string {
	return b.key
}

func Get(url string, i interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	v, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	json.Unmarshal(v, i)
	return nil
}

func getMappedStruct(url string, i interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	v, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	json.Unmarshal(v, i)
	return nil
}

func delete(url string) (bool, error) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return false, err
	}
	_, err = client.Do(req)
	if err != nil {
		return false, err
	}
	return true, nil
}

type BacklogTime struct {
	*time.Time
}

func (bt *BacklogTime) UnmarshalJSON(data []byte) error {
	str := string(data)
	t, err := time.Parse("\"2006-01-02T15:04:05Z07:00\"", str)
	*bt = BacklogTime{&t}
	return err
}

func (bt *BacklogTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(bt.Format("2018-11-12T15:04:05Z"))
}
