package main

import (
	"fmt"
	"path/filepath"
	"os"
	"html/template"

	"app/controller"

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
	g.GET("/index/welcome", func(c *gin.Context) {
	    c.HTML(200, "index/welcome.html", gin.H{
	    	"controller": "index",
			"action": "welcome",
	    })
	})

	user := g.Group("/user")
	{
		user.GET("/index", controller.UserList)
	}

	task := g.Group("/task")
	{
		task.GET("/index", controller.TaskList)

		task.GET("/add", func(c *gin.Context) {
		    c.HTML(200, "task/add.html", gin.H{
		    	"controller": "task",
				"action": "add",
		    })
		})

		task.GET("/calendar", controller.Calendar)

		task.GET("/card", controller.CardList)
	}

	doc := g.Group("/doc")
	{
		doc.GET("/index", controller.DocList)

		doc.GET("/add", func(c *gin.Context) {
		    c.HTML(200, "doc/add.html", gin.H{
		    	"controller": "doc",
				"action": "add",
		    })
		})
	}

	g.Run(":8090")
}

// built-in: role:: administrator/projector/employee/master.
// department/doc-category/task-type :: direct save chineses-name, can add.
// task-status: todo/doing/done
//
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
