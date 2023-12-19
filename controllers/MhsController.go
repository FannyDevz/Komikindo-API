package controllers

import (
	helpers "baca-manga/helpers"
	responses "baca-manga/response"

	"github.com/gin-gonic/gin"
)

func Mahsiswa(c *gin.Context) {

	// var HomeResp responses.DataMhs

	resp := helpers.Get("https://dinus.ac.id/mahasiswa/A11.2019.12096")

	if resp.StatusCode != 200 {
		c.JSON(500, gin.H{
			"status":      "error",
			"message":     "Something went wrong",
			"status_code": resp.StatusCode,
		})
		return
	}

	defer resp.Body.Close()

	c.JSON(200, responses.PingSuccess{
		Status:     "success",
		Message:    "Website is up",
		StatusCode: resp.StatusCode,
	})
}
