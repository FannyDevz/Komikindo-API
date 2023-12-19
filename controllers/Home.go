package controllers

import (
	helpers "baca-manga/helpers"
	responses "baca-manga/response"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {

	var HomeResp responses.MenuResponses

	resp := helpers.Get("https://komikindo.id/")
	if resp.StatusCode != 200 {
		c.JSON(500, responses.Error{
			Status:     "error",
			Message:    "Something went wrong",
			StatusCode: resp.StatusCode,
		})
		return
	}

	data, errs := goquery.NewDocumentFromReader(resp.Body)
	if errs != nil {
		c.JSON(500, responses.Error{
			Status:     "error",
			Message:    errs.Error(),
			StatusCode: resp.StatusCode,
		})
		return
	}
	data.Find("#menu-second-menu").Children().Each(func(i int, s *goquery.Selection) {
		name := s.Find("a").Text()
		url := s.Find("a").AttrOr("href", "No url")
		endpoint := strings.Replace(url, helpers.GetKeyEnv("BASE_URL"), "", -1)
		HomeResp.Menu = append(HomeResp.Menu, responses.DetailResponses{
			Name:     name,
			Url:      url,
			Endpoint: endpoint,
		})
	})
	data.Find(".mangapopuler").Find(".animepost").Each(func(i int, s *goquery.Selection) {
		name := s.Find("[itemprop='url']").AttrOr("title", "No title")
		thumnail := strings.Split(s.Find("img").AttrOr("src", "No src"), "?")[0]
		url := s.Find("[itemprop='url']").AttrOr("href", "No url")
		endpoint := strings.Replace(url, helpers.GetKeyEnv("BASE_URL"), "", -1)
		lastUpload := s.Find(".datech").Text()

		lastChapter := s.Find(".lsch").Find("a").Text()
		lastChapterUrl := s.Find(".lsch").Find("a").AttrOr("href", "No url")
		lastChapterEndpoint := strings.Replace(lastChapterUrl, helpers.GetKeyEnv("BASE_URL"), "", -1)

		HomeResp.Populars = append(HomeResp.Populars, responses.PopularManga{
			Name:       name,
			Thumbnail:  strings.Replace(thumnail, "https://", "https://cdn.statically.io/img/", -1),
			Url:        url,
			Endpoint:   endpoint,
			LastUpload: lastUpload,
			LastChapter: []responses.DetailResponses{
				{
					Name:     lastChapter,
					Url:      lastChapterUrl,
					Endpoint: lastChapterEndpoint,
				},
			},
		})
	})

	data.Find(".latestupdate-v2").Find(".animepost").Each(func(i int, s *goquery.Selection) {
		name := s.Find("[itemprop='url']").AttrOr("title", "No title")
		thumnail := strings.Split(s.Find("img").AttrOr("src", "No src"), "?")[0]
		url := s.Find("[itemprop='url']").AttrOr("href", "No url")
		endpoint := strings.Replace(url, helpers.GetKeyEnv("BASE_URL"), "", -1)

		HomeResp.Latest = append(HomeResp.Latest, responses.LatesUpdate{
			Name:      name,
			Thumbnail: strings.Replace(thumnail, "https://", "https://cdn.statically.io/img/", -1),
			Url:       url,
			Endpoint:  endpoint,
		})
	})
	defer resp.Body.Close()

	c.JSON(200, responses.Success{
		Message: "success",
		Data:    HomeResp,
	})
}
