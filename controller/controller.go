package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"

	"app/helper"
	"app/module"

    "app/middleware/auth"

	// "html/template"
	// "fmt"
	// "log"
	"net/http"
	"strings"
	"time"
	"strconv"
)


func Index(c *gin.Context) {
	//
}

func Login(c *gin.Context) {
	
	if helper.IsAjax(c) {

		var user module.User

		if err := c.ShouldBind(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		password := helper.Md5(user.Password)

		user = module.UserInfoByName(user.Username)

		if user == (module.User{}) || password != user.Password {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "失败",
				"data": "",
			})
		} else {

			token,err := generateToken(user)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": 1,
					"msg": err.Error(),
					"data": "",
				})
				return
			}

			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "成功",
				"data": gin.H{"token": token},
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


func generateToken(user module.User) (string, error) {

	j := auth.NewJWT()
	claims := auth.CustomClaims{
		strconv.Itoa(int(user.Id)),
		user.Username,
		user.Phone,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 86400 * 7), // 过期时间一周
			Issuer:    "project-manager",                   //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	return token, err
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

		user.Password = helper.Md5(user.Password)

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

		switch module.UserEdit(user.Id, user) {
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
	// var field = c.PostForm("field")
	// var value = c.PostForm("value")
	// var uid = c.PostForm("id")

	type UM struct {
	    Id uint `form:"id" json:"id"`
	    Field string `form:"field" json:"field"`
	    Value interface{} `form:"value" json:"value"`
	}

	var um UM

	if err := c.ShouldBind(&um); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}


	if module.UserModify(um.Id, um.Field, um.Value) {

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
    password = helper.Md5(password)

	if helper.IsAjax(c) {
		if module.UserModify(1, "password", password) { // todo

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

//
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@pro
//

func ProList(c *gin.Context) {
	if helper.IsAjax(c) {

		c.JSON(200, gin.H{
			"code":  1,
			"msg":   "",
			"count": 10,
			"data":  module.ProList(),
		})

	} else {

		c.HTML(200, "pro/index.html", gin.H{
			"controller": "pro",
			"action":     "index",
		})
	}
}

func ProAdd(c *gin.Context) {

	if helper.IsAjax(c) {

		var pro module.Project

		up_date := strings.Split(c.PostForm("up_date"), " ~ ")

		startTime, _ := time.Parse("2006-01-02", up_date[0])
		pro.StartTime = startTime.Unix()

		endTime, _ := time.Parse("2006-01-02", up_date[1])
		pro.EndTime = endTime.Unix()

		if err := c.ShouldBind(&pro); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if module.ProAdd(pro) {
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

		c.HTML(200, "pro/add.html", gin.H{
			"controller": "pro",
			"action":     "add",
		})
	}
}

func ProEdit(c *gin.Context) {
	var id = c.Query("id")

	if helper.IsAjax(c) {
		var pro module.Project

		up_date := strings.Split(c.PostForm("up_date"), " ~ ")

		startTime, _ := time.Parse("2006-01-02", up_date[0])
		pro.StartTime = startTime.Unix()

		endTime, _ := time.Parse("2006-01-02", up_date[1])
		pro.EndTime = endTime.Unix()

		if err := c.ShouldBind(&pro); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		switch module.ProEdit(id, pro) {
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

		c.HTML(200, "pro/edit.html", gin.H{
			"controller": "pro",
			"action":     "edit",
			"data":       module.ProInfo(id),
		})
	}
}


func ProDelete(c *gin.Context) {
	var id = c.Query("id")

	if module.ProDelete(id) {

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
