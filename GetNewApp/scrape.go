package main

import (
	"errors"
	"net/http"
	"strings"

	"github.com/kimulaco/daily-steam-core/app"
	"github.com/kimulaco/daily-steam-core/date"

	"github.com/PuerkitoBio/goquery"
)

func Get(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return res, err
	}
	if res.StatusCode != 200 {
		return res, errors.New("failed http request")
	}

	return res, nil
}

func ParseNewApps(res *http.Response, d date.Date) ([]app.App, error) {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return []app.App{}, err
	}

	var apps []app.App
	doc.Find("a.search_result_row").Each(func(i int, s *goquery.Selection) {
		releasedDate, err := date.ParseDate(s.Find(".search_released").Text())

		if err != nil || releasedDate.IsEmpty() || !releasedDate.Equal(d) {
			return
		}

		_p := s.Find(".search_price")
		salePrice := strings.TrimSpace(_p.Find("strike").Text())
		var price string
		if salePrice == "" {
			price = strings.TrimSpace(_p.Text())
		} else {
			_p = _p.Children().RemoveFiltered("span").Parent()
			price = strings.TrimSpace(_p.Text())
		}

		if price == "" {
			return
		}

		id, _ := s.Attr("data-ds-appid")
		url, _ := s.Attr("href")
		title := s.Find(".title").Text()
		thumbUrl, _ := s.Find(".search_capsule img").Attr("src")

		apps = append(apps, app.App{
			Id:         id,
			Url:        url,
			Title:      title,
			ThumbUrl:   thumbUrl,
			ReleasedAt: releasedDate.ToString(),
			Price:      price,
			SalePrice:  salePrice,
		})
	})

	return apps, nil
}
