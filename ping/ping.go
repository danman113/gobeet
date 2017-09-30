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
	resChan <- nil
}

func PingWebsite(website *site.Website) {
	for {
		fmt.Println(website.Url)
		for i, _ := range website.Pages {
			ret := make(chan error)
			go PingStatus(&(website.Pages[i]), ret)
			for e := range ret {
				if e != nil {
					handleError(e, &(website.Pages[i]))
				} else {
					website.Pages[i].DownSince = nil
				}
			}
		}
		time.Sleep(time.Duration(website.Interval) * time.Second)
	}
}

func handleError(e error, pg *site.Page) {
	fmt.Println("Error: ")
	fmt.Println(e)
	fmt.Println(*pg)
	if pg.DownSince != nil {
		since := time.Since(*pg.DownSince)
		if since.Seconds() > float64(pg.Duration) {
			fmt.Println("Sent Email")
			pg.DownSince = nil
			email.SendAlert(e, pg)
		}
	} else {
		now := time.Now()
		pg.DownSince = &now
	}
}
