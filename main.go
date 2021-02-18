package main

import (
	"fmt"
	"html/template"
	"regexp"
	"net/http"
	"path/filepath"

	"app/controller"
	"app/helper"

	"github.com/gin-gonic/gin"
)

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/login" {
			return
		}

		match, _ := regexp.MatchString(".*.(js|css|jpg|png|ico)", c.Request.URL.Path)
		if match {
			return
		}

		// Set example variable
		cookie, err := c.Cookie("_token_")

        if err == nil {
        	uid,_ := helper.Decrypt([]byte(cookie))
        	if len(uid) < 1 {
        		c.Redirect(http.StatusFound, "/login")
        	}
        } else {
        	c.Redirect(http.StatusFound, "/login")
        }

		c.Next()
	}
}

func main() {

	gin.SetMode(gin.ReleaseMode)

	g := gin.New()

	g.Use(gin.Recovery())
	g.Use(auth())

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

	g.Any("/login", controller.Login)

	user := g.Group("/user")
	{
		user.GET("/index", controller.UserList)

		user.Any("/add", controller.UserAdd)

		user.Any("/edit", controller.UserEdit)

		user.POST("/modify", controller.UserModify)

		user.Any("/password", controller.UserPassword)

		user.POST("/delete", controller.UserDelete)
	}

	task := g.Group("/task")
	{
		task.GET("/index", controller.TaskList)

		task.Any("/add", controller.TaskAdd)

		task.Any("/edit", controller.TaskEdit)

		task.POST("/delete", controller.TaskDelete)

		task.POST("/modify/:id", controller.TaskModify)

		task.GET("/card", controller.CardList)

		task.GET("/calendar", controller.Calendar) // todo
	}

	doc := g.Group("/doc")
	{
		doc.GET("/index", controller.DocList)

		doc.Any("/add", controller.DocAdd)

		doc.Any("/edit", controller.DocEdit)

		doc.POST("/delete", controller.DocDelete)

		doc.GET("/view/:id", controller.DocView)
	}

	g.Run(":8090")
}

// built-in: role:: administrator/projector/employee/master.
// department?/doc-category/task-type :: direct save chineses-name, can add.
// task-status: todo/doing/done
// project:  can add

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
