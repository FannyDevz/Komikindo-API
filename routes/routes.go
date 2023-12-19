package routes

import (
	"baca-manga/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	router.GET("/ping", controllers.Ping)
	router.GET("/home", controllers.Home)
	router.GET("/ivan", controllers.ControllerIvan)
	router.GET("/pokedex", controllers.BejjoController)
	router.GET("/mahasiswa", controllers.Mahsiswa)
	router.GET("/ping1", controllers.Ping)
	router.GET("/enzd", controllers.EnzdController)
	router.GET("/wagyuA5", controllers.WagyuController)
	router.GET("/villaneuvo", controllers.ControllerVillaneuvo)
	router.GET("/itsuka", controllers.ControllerHandoko)
}
