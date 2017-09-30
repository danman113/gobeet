package ping

import (
	"errors"
	"fmt"
	"github.com/danman113/gobeet/email"
	"github.com/danman113/gobeet/site"
	"net/http"
	"time"
)

var (
	client = &http.Client{}
)

func GobeetError(msg string) error {
	return errors.New("Gobeet: " + msg)
}

func PingStatus(site *site.Page, resChan chan error) {
	defer close(resChan)
	client.Timeout = time.Duration(site.Timeout) * time.Millisecond
	req, errReq := http.NewRequest(site.Method, site.Url, nil)
	if errReq != nil {
		resChan <- errReq
		return
	}
	resHttp, errHttp := client.Do(req)
	if errHttp != nil {
		resChan <- errHttp
		return
	}
	if resHttp.StatusCode != site.Status {
		resChan <- GobeetError(fmt.Sprintf("%s: Method %d != %d", site.Url, resHttp.StatusCode, site.Status))
		return
	}
	fmt.Printf("\t %s = %d\n", site.Url, site.Status)
}

func PingWebsite(website *site.Website) {
	for {
		fmt.Println(website.Url)
		for _, page := range website.Pages {
			ret := make(chan error)
			go PingStatus(&page, ret)
			for e := range ret {
				handleError(e, &page)
			}
		}
		time.Sleep(time.Duration(website.Interval) * time.Millisecond)
	}
}

func handleError(e error, pg *site.Page) {
	fmt.Println("Error: ")
	fmt.Println(e)
	fmt.Println(*pg)
	email.SendAlert(e, pg)
}
