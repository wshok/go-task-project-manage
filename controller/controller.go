package controller

import (
	"github.com/gin-gonic/gin"

	"app/helper"
	"app/module"
	// "html/template"
	"fmt"
	"crypto/md5"
	"strings"
	"time"
	"strconv"
)


func Login(c *gin.Context) {
	
	if helper.IsAjax(c) {

		var user module.User

		if err := c.ShouldBind(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var username = c.PostForm("username")
		var password = c.PostForm("password")
		password = fmt.Sprintf("%x", md5.Sum([]byte(password)))

		user = module.UserInfoByName(username)

		if user == (module.User{}) || password != user.Password {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "",
			})
		} else {
			var host =c.GetHeader("Host")
			c.SetCookie("_token_", strconv.FormatUint(uint64(user.Id), 10), 0, "/", host, false, true)

			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "成功",
				"data": "",
			})
		}
	} else {

		c.HTML(200, "login.html", gin.H{
			"controller": "login",
			"action":     "index",
			"data": "",
			"captcha": 0,
		})
	}
}

//
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@user
//

func UserList(c *gin.Context) {

	if helper.IsAjax(c) {

		c.JSON(200, gin.H{
			"code":  1,
			"msg":   "",
			"count": 10,
			"data":  module.UserList(),
		})

	} else {

		c.HTML(200, "user/index.html", gin.H{
			"controller": "user",
			"action":     "index",
		})
	}
}

func UserAdd(c *gin.Context) {

	if helper.IsAjax(c) {

		var user module.User

		if err := c.ShouldBind(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		user.Password = fmt.Sprintf("%x", md5.Sum([]byte(user.Password)))

		if module.UserAdd(user) {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "添加成功",
				"data": "",
			})
		} else {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "添加失败",
				"data": "",
			})
		}
	} else {

		c.HTML(200, "user/add.html", gin.H{
			"controller": "user",
			"action":     "add",
		})
	}
}

func UserEdit(c *gin.Context) {
	var uid = c.Query("id")

	if helper.IsAjax(c) {
		var user module.User

		if err := c.ShouldBind(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		switch module.UserEdit(uid, user) {
		case -1:
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "记录不存在",
				"data": "",
			})
		case 0:
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "修改失败",
				"data": "",
			})
		case 1:
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "修改成功",
				"data": "",
			})
		}

	} else {

		c.HTML(200, "user/edit.html", gin.H{
			"controller": "user",
			"action":     "edit",
			"data":       module.UserInfo(uid),
		})
	}
}

func UserModify(c *gin.Context) {
	var field = c.PostForm("field")
	// value, _ := strconv.Atoi(c.PostForm("value"))
	var value = c.PostForm("value")
	var uid = c.PostForm("id")

	if module.UserModify(uid, field, value) {

		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "修改成功",
			"data": "",
		})

	} else {

		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "修改失败",
			"data": "",
		})
	}
}

func UserPassword(c *gin.Context) {
	var uid = c.Query("id")
	var password = c.PostForm("password")
    password = fmt.Sprintf("%x", md5.Sum([]byte(password)))

	if helper.IsAjax(c) {
		if module.UserModify(uid, "password", password) {

			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "修改成功",
				"data": "",
			})

		} else {

			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "修改失败",
				"data": "",
			})
		}
	} else {

		c.HTML(200, "user/password.html", gin.H{
			"controller": "user",
			"action":     "edit",
			"data":       module.UserInfo(uid),
		})
	}
}

func UserDelete(c *gin.Context) {
	var uid = c.Query("id")

	if module.UserDelete(uid) {

		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "删除成功",
			"data": "",
		})

	} else {

		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "删除失败",
			"data": "",
		})
	}
}

//
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@task
//

func TaskList(c *gin.Context) {
	if helper.IsAjax(c) {

		c.JSON(200, gin.H{
			"code":  1,
			"msg":   "",
			"count": 10,
			"data":  module.TaskList(),
		})

	} else {

		c.HTML(200, "task/index.html", gin.H{
			"controller": "task",
			"action":     "index",
		})
	}
}

func TaskAdd(c *gin.Context) {

	if helper.IsAjax(c) {

		var task module.Task

		up_date := strings.Split(c.PostForm("up_date"), " ~ ")

		startTime, _ := time.Parse("2006-01-02", up_date[0])
		task.StartTime = startTime.Unix()

		endTime, _ := time.Parse("2006-01-02", up_date[1])
		task.EndTime = endTime.Unix()

		if err := c.ShouldBind(&task); err != nil {
			c.JSON(200, gin.H{"error": err.Error()})
			return
		}

		if module.TaskAdd(task) {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "添加成功",
				"data": "",
			})
		} else {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "添加失败",
				"data": "",
			})
		}
	} else {

		c.HTML(200, "task/add.html", gin.H{
			"controller": "task",
			"action":     "add",
		})
	}
}

func TaskEdit(c *gin.Context) {
	var taskId = c.Query("id")

	if helper.IsAjax(c) {
		var task module.Task

		up_date := strings.Split(c.PostForm("up_date"), " ~ ")

		startTime, _ := time.Parse("2006-01-02", up_date[0])
		task.StartTime = startTime.Unix()

		endTime, _ := time.Parse("2006-01-02", up_date[1])
		task.EndTime = endTime.Unix()

		if err := c.ShouldBind(&task); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		switch module.TaskEdit(taskId, task) {
		case -1:
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "记录不存在",
				"data": "",
			})
		case 0:
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "修改失败",
				"data": "",
			})
		case 1:
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "修改成功",
				"data": "",
			})
		}

	} else {

		c.HTML(200, "task/edit.html", gin.H{
			"controller": "task",
			"action":     "edit",
			"data":       module.TaskInfo(taskId),
		})
	}
}

func TaskModify(c *gin.Context) {
	var status = c.PostForm("status")
	var taskId = c.Param("id")

	if module.TaskModify(taskId, status) {

		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "修改成功",
			"data": "",
		})

	} else {

		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "修改失败",
			"data": "",
		})
	}
}

func TaskDelete(c *gin.Context) {
	var taskId = c.Query("id")

	if module.TaskDelete(taskId) {

		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "删除成功",
			"data": "",
		})

	} else {

		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "删除失败",
			"data": "",
		})
	}
}

func CardList(c *gin.Context) {

	c.HTML(200, "task/card.html", gin.H{
		"controller": "task",
		"action":     "card",
		"data":       module.TaskList(),
	})
}

func Calendar(c *gin.Context) {
	c.HTML(200, "task/calendar.html", gin.H{
		"controller": "task",
		"action":     "card",
	})
}

//
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@doc
//

func DocList(c *gin.Context) {
	if helper.IsAjax(c) {

		c.JSON(200, gin.H{
			"code":  1,
			"msg":   "",
			"count": 10,
			"data":  module.DocList(),
		})

	} else {

		c.HTML(200, "doc/index.html", gin.H{
			"controller": "doc",
			"action":     "index",
		})
	}
}

func DocAdd(c *gin.Context) {

	if helper.IsAjax(c) {

		var doc module.Doc

		if err := c.ShouldBind(&doc); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if module.DocAdd(doc) {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "添加成功",
				"data": "",
			})
		} else {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "添加失败",
				"data": "",
			})
		}
	} else {

		c.HTML(200, "doc/add.html", gin.H{
			"controller": "doc",
			"action":     "add",
		})
	}
}

func DocEdit(c *gin.Context) {
	var id = c.Query("id")

	if helper.IsAjax(c) {
		var doc module.Doc

		if err := c.ShouldBind(&doc); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		switch module.DocEdit(id, doc) {
		case -1:
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "记录不存在",
				"data": "",
			})
		case 0:
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "修改失败",
				"data": "",
			})
		case 1:
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "修改成功",
				"data": "",
			})
		}

	} else {

		c.HTML(200, "doc/edit.html", gin.H{
			"controller": "doc",
			"action":     "edit",
			"data":       module.DocInfo(id),
		})
	}
}

func DocView(c *gin.Context) {
	var id = c.Param("id")
	
	c.HTML(200, "doc/view.html", gin.H{
		"controller": "doc",
		"action":     "view",
		"data":       module.DocInfo(id),
	})
}

func DocDelete(c *gin.Context) {
	var id = c.Query("id")

	if module.DocDelete(id) {

		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "删除成功",
			"data": "",
		})

	} else {

		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "删除失败",
			"data": "",
		})
	}
}

// func Index(c *gin.Context) {

// 	curPage, _ := strconv.Atoi(c.Param("p"))

// 	f := &module.Filter {
// 		Page: "index",
// 	}

// 	p := &helper.Pager {
// 		TotalRows: module.ArticleCount(f),
// 		Method:    "html",
// 		Parameter: "/page/?",
// 		NowPage:   curPage,
// 	}

// 	Page := helper.NewPager(p, "1")

// 	ArtList := module.ArticleList(
// 		&module.LimitRows{
// 			Start: (p.NowPage - 1) * p.ListRows,
// 			Size:  p.ListRows,
// 		}, f)

// 	c.HTML(200, "index.html", gin.H{
// 		"ArtList": ArtList,
// 		"Page": Page,

// 		"LastPost": module.LeastPosted(),
// 		"Category": module.Category(),
// 		"Archive": module.Archive(),
// 	})
// }
