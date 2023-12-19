package controllers

import (
	helpers "baca-manga/helpers"
	responses "baca-manga/response"

	"github.com/PuerkitoBio/goquery"

	"github.com/gin-gonic/gin"
)

func BejjoController(c *gin.Context) {

	var PokeRes responses.PokedexList

	resp := helpers.Get("https://pokemondb.net/pokedex/all/")
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

	data.Find("#pokedex").Find("tr").Each(func(i int, s *goquery.Selection) {
		td := s.Find("td")
		PokeRes.List = append(PokeRes.List, responses.Pokedex{
			Id:        td.Eq(0).Text(),
			Name:      td.Eq(1).Text(),
			Total:     td.Eq(3).Text(),
			HP:        td.Eq(4).Text(),
			Attack:    td.Eq(5).Text(),
			Defense:   td.Eq(6).Text(),
			Speed:     td.Eq(9).Text(),
			SpAttack:  td.Eq(7).Text(),
			SpDefense: td.Eq(8).Text(),
		})
	})

	defer resp.Body.Close()

	c.JSON(200, responses.Success{
		Message: "success",
		Data:    PokeRes,
	})
}
