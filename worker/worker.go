package worker

import (
	"time"
	"fmt"
	"net/http"
	"io/ioutil"
	"crypto/md5"
	"encoding/hex"
	"github.com/gophergala/not_golang_experts/model"
)

var stopchannel		chan bool
var ticker				*time.Ticker

func StartObserving(stopped chan bool) {
	stopchannel = stopped

	ticker = time.NewTicker(time.Millisecond * 1500) // 1.5 secs

	go observe()
}

func StopObserving() {
	ticker.Stop()
	stopchannel <- true
}

func observe() {
	for t := range ticker.C {
		pagestocheck := model.PagesToCheck()
		for _, page := range pagestocheck {
			fmt.Printf("Checking page: %v - %v\n", page.Url, t)
			resultchan := make(chan string)
			go requestHTML(*page, resultchan)

			resultString := <-resultchan
			if page.HtmlString != resultString {
				fmt.Printf("%v has a change!\n\n", page.Url)
				page.HtmlString = resultString
			}else{
				page.LastCheckedAt = time.Now()
			}

			page.Save()
		}
	}
}

func requestHTML(p model.Page, result chan string) {
	res, err := http.Get(p.Url)

	if err!=nil {
		panic(err)
	}else{
		defer res.Body.Close()
		html, err := ioutil.ReadAll(res.Body)
		if err!=nil {
			panic(err)
		}

		hasher := md5.New()
		hasher.Write([]byte(html))
		result <- hex.EncodeToString(hasher.Sum(nil))
	}
}
