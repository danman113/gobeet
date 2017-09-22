package ping

import (
	"fmt"
	"github.com/danman113/gobeet/site"
	"net/http"
	"time"
)

var (
	client = &http.Client{}
)

func PingStatus(site *site.Page, res chan error) {
	client.Timeout = time.Duration(site.Timeout) * time.Millisecond
	req, err := http.NewRequest(site.Method, site.Url, nil)
	if err != nil {
		res <- error
	}
	res, err := client.Do(req)
	if err != nil {
		res <- error
	}
	fmt.Println(res)
}
