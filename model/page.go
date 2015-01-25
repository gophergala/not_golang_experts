package model

import (
	"time"
)

type Page struct {
	Id            int64
	Url           string
	LastCheckedAt time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	HtmlString    string
}

func PagesToCheck() []*Page {
	DB.LogMode(false)

	var pages []*Page
	DB.Where("last_checked_at < ?", time.Now().Add(-time.Second*5)).Find(&pages)
	return pages
}

func FindOrCreatePageByUrl(url string) Page {
	page := Page{LastCheckedAt: time.Now()}
	DB.Where(Page{Url: url}).FirstOrCreate(&page)
	return page
}

func (p Page) Save() {
	DB.Save(&p)
}
