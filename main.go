package main

import (
	"fmt"
	"path/filepath"
	"os"
	"html/template"

	"github.com/gin-gonic/gin"
)


func main() {

	gin.SetMode(gin.ReleaseMode)

	g := gin.New()

	g.SetFuncMap(helperFuncs)

	g.LoadHTMLGlob(filepath.Join("", "./view/**/*.html"))

	g.Static("/static", filepath.Join("", "./static"))
	g.Static("/plugs", filepath.Join("", "./static/plugs"))

	g.Static("/api", filepath.Join("", "./api"))

	g.GET("/", func(c *gin.Context) {
		c.HTML(200, "index/index.html", gin.H{
			"controller": "index",
			"action": "index",
		})
	})
	g.GET("/index/welcome.html", func(c *gin.Context) {
	    c.HTML(200, "index/welcome.html", gin.H{
	    	"controller": "index",
			"action": "welcome",
	    })
	})

	admin := g.Group("/admin")
	{
		admin.GET("/index.html", func(c *gin.Context) {
		    c.HTML(200, "admin/index.html", gin.H{
		    	"controller": "admin",
				"action": "index",
		    })
		})
	}

	auth := g.Group("/auth")
	{
		auth.GET("/index.html", func(c *gin.Context) {
		    c.HTML(200, "auth/index.html", gin.H{
		    	"controller": "auth",
				"action": "index",
		    })
		})
	}

	task := g.Group("/task")
	{
		task.GET("/index.html", func(c *gin.Context) {
		    c.HTML(200, "task/index.html", gin.H{
		    	"controller": "task",
				"action": "index",
		    })
		})
		task.GET("/add.html", func(c *gin.Context) {
		    c.HTML(200, "task/add.html", gin.H{
		    	"controller": "task",
				"action": "add",
		    })
		})
	}

	g.Run(":8090")
}


// func fileExists(fpath string) bool {
// 	if _, err := os.Stat(fpath); err == nil {
// 		return true
// 	}
// 	return false
// }

var helperFuncs = template.FuncMap {
	"jsExists": func (fpath string) bool {
		jspath := fmt.Sprintf("./static/js/%s.js", fpath)
		if _, err := os.Stat(jspath); err == nil {
			return true
		}
		return false
	},
}