package controllers

import (
	Api "baca-manga/helpers"
	responses "baca-manga/response"

	"github.com/gin-gonic/gin"
)

func EnzdController(c *gin.Context) {
	resp := Api.Get("http://komikindo.id/")
	if resp.StatusCode != 200 {
		c.JSON(500, gin.H{
			"status":      "error",
			"message":     "Something went wrong",
			"status_code": resp.StatusCode,
		})
		return
	}

	defer resp.Body.Close()

	c.JSON(404, responses.IvanError{
		Status:     "error",
		Message:    "something missing",
		StatusCode: resp.StatusCode,
	})

}
