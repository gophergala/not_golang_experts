package worker

import (
	"time"
	"fmt"
	"net/http"
	"io/ioutil"
)

type Page struct {
	Url					string
	CheckedAt		time.Time
	HTMLString	string
}

var stopchannel		chan bool
var pagestocheck	[]*Page
var ticker				*time.Ticker

func StartObserving(stopped chan bool) {
	stopchannel = stopped
	pagestocheck = append(pagestocheck, &Page{"http://swanros.com/about/", time.Now(), ""}, &Page{"http://swanros.com/contact/", time.Now(), ""} )

	fmt.Printf("%v\n\n", pagestocheck)

	ticker = time.NewTicker(time.Millisecond * 1500) // 1.5 secs

	go observe()
}

func StopObserving() {
	ticker.Stop()
	stopchannel <- true
}

func observe() {
	for t := range ticker.C {
		for _, page := range pagestocheck {
			fmt.Printf("Checking page: %v - %v\n", page.Url, t)
			resultchan := make(chan string)
			go requestHTML(*page, resultchan)

			resultString := <-resultchan
			if page.HTMLString != resultString {
				fmt.Printf("\n\n%v has a change!\n\n", page.Url)
				page.HTMLString = resultString
			}
		}
	}
}

func requestHTML(p Page, result chan string) {
	res, err := http.Get(p.Url)

	if err!=nil {
		panic(err)
	}else{
		defer res.Body.Close()
		html, err := ioutil.ReadAll(res.Body)
		if err!=nil {
			panic(err)
		}
		result <- string(html)
	}
}
