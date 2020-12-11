package main

import (
	"fmt"
	"html/template"
	"path/filepath"

	"app/controller"
	"app/helper"

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
			"action":     "index",
		})
	})

	g.GET("/index/welcome", func(c *gin.Context) {
		c.HTML(200, "index/welcome.html", gin.H{
			"controller": "index",
			"action":     "welcome",
		})
	})

	user := g.Group("/user")
	{
		user.GET("/index", controller.UserList)

		user.Any("/add", controller.UserAdd)

		user.Any("/edit", controller.UserEdit)

		user.POST("/del", controller.UserDelete)
	}

	task := g.Group("/task")
	{
		task.GET("/index", controller.TaskList)

		task.Any("/add", controller.TaskAdd)

		task.Any("/edit", controller.TaskEdit)

		task.POST("/del", controller.TaskDelete)

		task.POST("/modify/:id", controller.TaskModify)

		task.GET("/card", controller.CardList)

		task.GET("/calendar", controller.Calendar) // todo
	}

	doc := g.Group("/doc")
	{
		doc.GET("/index", controller.DocList)

		doc.GET("/add", controller.DocAdd)

		doc.Any("/edit", controller.DocEdit)

		doc.POST("/del", controller.DocDelete)
	}

	g.Run(":8090")
}

// built-in: role:: administrator/projector/employee/master.
// department/doc-category/task-type :: direct save chineses-name, can add.
// task-status: todo/doing/done
//

var helperFuncs = template.FuncMap{
	"jsExists": func(fpath string) bool {
		jspath := fmt.Sprintf("./static/js/%s.js", fpath)
		if _, err := helper.PathExists(jspath); err == nil {
			return true
		}
		return false
	},
	"timeFormat": helper.TimeFormat,
}
