package routes

import (
	"github.com/gin-gonic/gin"

	"fita-test-02/controllers/doctorsController"
	homepageController "fita-test-02/controllers/homepageController"
)

type routes struct {
	router *gin.Engine
}

func (r routes) Run(addr ...string) error {
	return r.router.Run()
}

func InitializeRoute() routes {
	r := routes{
		router: gin.Default(),
	}

	r.router.Use(gin.Recovery())

	r.router.GET("/", homepageController.Index)

	doctor := r.router.Group("/doctors")
	r.routesDoctors(doctor)

	return r
}

func (r routes) routesDoctors(rg *gin.RouterGroup) {
	rg.GET("/index", doctorsController.Index)
	rg.GET("/localtimezone", doctorsController.LocalTime)
}
