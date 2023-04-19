package main

import (
	"fmt"
	// "log"
	"html/template"
	"regexp"
	"net/http"
	"path/filepath"
	"strings"

	"app/controller"
	"app/helper"
	"app/middleware/auth"

	"github.com/gin-gonic/gin"
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/login" {
			return
		}

		match, _ := regexp.MatchString(".*.(js|css|jpg|png|ico)", c.Request.URL.Path)
		if match {
			return
		}

		// Set example variable
		//cookie, err := c.Cookie("_token_")

   //      if err == nil {
   //          uid,_ := helper.Decrypt([]byte(cookie))
			// if len(uid) < 1 {
			// 	c.Redirect(http.StatusFound, "/login")
			// }
   //      } else {
			// c.Redirect(http.StatusFound, "/login")
   //      }

        authorization := c.Request.Header.Get("Authorization")

        token := strings.Split(string(authorization), " ")

		if (len(token) < 2) || (token[1] == "") {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}

		// log.Print("get token: ", token[1])

		j := auth.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token[1])
		if err != nil {
			if err == auth.TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"status": -1,
					"msg":    "授权已过期",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)

		c.Next()
	}
}

func main() {

	gin.SetMode(gin.ReleaseMode)

	g := gin.New()

	g.Use(gin.Recovery())
	g.Use(authMiddleware())

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

		task.GET("/calendar", controller.Calendar)
	}

	doc := g.Group("/doc")
	{
		doc.GET("/index", controller.DocList)

		doc.Any("/add", controller.DocAdd)

		doc.Any("/edit", controller.DocEdit)

		doc.POST("/delete", controller.DocDelete)

		doc.GET("/view/:id", controller.DocView)
	}

	pro := g.Group("/project")
	{
		pro.GET("/index", controller.ProList)

		pro.Any("/add", controller.ProAdd)

		pro.Any("/edit", controller.ProEdit)

		pro.POST("/delete", controller.ProDelete)
	}

	g.Run(":4001")
}

// built-in: role:: administrator/projector/employee/master.
// department?/doc-category/task-type :: direct save chineses-name, can add.
// task-status: todo/doing/done


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
