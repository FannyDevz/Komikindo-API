package controllers

import (
	helper "baca-manga/helpers"
	"baca-manga/response"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

var MenuResp response.MenuResponses

func ControllerVillaneuvo(ctx *gin.Context) {
	resp := helper.Get("https://komiku.id/")

	if resp.StatusCode != 200 {
		ctx.JSON(500, gin.H{
			"status":      "error",
			"status_code": resp.StatusCode,
			"message":     "Something went wrong",
		})
		return
	}

	data, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		ctx.JSON(500, gin.H{
			"status":      "error",
			"status_code": resp.StatusCode,
			"message":     err.Error(),
		})
	}

	data.Find(".second_nav").Children().Each(func(i int, s *goquery.Selection) {
		name := s.Find("a").Text()
		url := s.Find("a").AttrOr("href", "No URL")
		endPoint := strings.Replace(url, "https://komiku.id/", "", -1)
		MenuResp.Menu = append(MenuResp.Menu, response.DetailResponses{
			Name:     name,
			Url:      url,
			Endpoint: "https://komiku.id" + endPoint,
		})
	})

	defer resp.Body.Close()

	ctx.JSON(200, gin.H{
		"status":      "success",
		"status_code": resp.StatusCode,
		"message":     MenuResp,
	})
}
