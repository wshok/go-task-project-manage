package main

import (
	// "fmt"
	"path/filepath"
	// "html/template"

	"github.com/gin-gonic/gin"
)


func main() {

	gin.SetMode(gin.ReleaseMode)

	g := gin.New()

	g.LoadHTMLGlob(filepath.Join("", "./view/**/*.html"))

	g.Static("/static", filepath.Join("", "./static"))
	g.Static("/plugs", filepath.Join("", "./static/plugs"))

	g.Static("/api", filepath.Join("", "./api"))

	g.GET("/", func(c *gin.Context) {
		c.HTML(200, "index/index.html", gin.H{})
	})
	g.GET("/index/welcome.html", func(c *gin.Context) {
	    c.HTML(200, "index/welcome.html", gin.H{})
	})

	g.Run(":8090")
}



