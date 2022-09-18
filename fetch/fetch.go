package fetch

import (
	"io/ioutil"
	"net/http"
	"time"
)

func Fetch(url string) ([]byte, error) {
	client := http.Client{Timeout: time.Second * 50}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	r, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return r, nil

}
