package controllers

import (
	Api "baca-manga/helpers"
	responses "baca-manga/response"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func ListKomik(c *gin.Context) {
	var ListKomikResponse responses.ListResponsesManga

	page := c.Param("page")
	if page == "" {
		page = "1"
	}

	resp := Api.Get("https://komikindo.id/daftar-manga/page/" + page)

	if resp.StatusCode != 200 {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Terjadi kesalahan",
			"code":    resp.StatusCode,
		})
		return
	}

	data, errs := goquery.NewDocumentFromReader(resp.Body)

	if errs != nil {
		fmt.Println(errs)
	}

	data.Find(".animepost").Find("li").Each(func(i int, s *goquery.Selection) {

		name := s.Find("[itemprop=url]").AttrOr("title", "No Title")
		thumbnail := strings.Split(s.Find("img").AttrOr("src", "No Image"), "?")[0]
		link := s.Find("[itemprop=url]").AttrOr("href", "No Link")
		endpoint := strings.Replace(link, "https://komikindo.id/", "", -1)

		ListKomikResponse.Manga = append(ListKomikResponse.Manga, responses.ListManga{
			Name:      name,
			Thumbnail: strings.Replace(thumbnail, "https://", "https://cdn.statically.io/img/", -1),
			Url:       link,
			Endpoint:  endpoint,
		})

	})

}
