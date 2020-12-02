package controller

import (
	"github.com/gin-gonic/gin"

	"app/module"
	"app/helper"

	// "html/template"
	// "fmt"
	// "strings"
	// "strconv"
)


func UserList(c *gin.Context) {

	if helper.IsAjax(c) {

		c.JSON(200, gin.H{
			"code": 0,
			"msg": "",
			"count": 10,
			"data": module.UserList(),
		})

	} else {

		c.HTML(200, "user/index.html", gin.H{
	    	"controller": "user",
			"action": "index",
	    })
	}
}

func TaskList(c *gin.Context) {
	if helper.IsAjax(c) {

		c.JSON(200, gin.H{
			"code": 0,
			"msg": "",
			"count": 10,
			"data": module.TaskList(),
		})

	} else {

		c.HTML(200, "task/index.html", gin.H{
	    	"controller": "task",
			"action": "index",
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


// func Category (c *gin.Context) {
// 	var url string = c.Param("name")

// 	curPage, _ := strconv.Atoi(c.Param("p"))

// 	f := &module.Filter {
// 		Page: "category",
// 		Category: url,
// 	}

// 	p := &helper.Pager{
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

// func Archive (c *gin.Context) {
// 	var y string = c.Param("y")
// 	var m string = c.Param("m")

// 	curPage, _ := strconv.Atoi(c.Param("p"))

// 	f := &module.Filter {
// 		Page: "archive",
// 		Year: y,
// 		Month: m,
// 	}

// 	p := &helper.Pager{
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


// func Article(c *gin.Context) {

// 	var url string = c.Param("url")
	
// 	url = strings.TrimSuffix(url, ".html")

// 	c.HTML(200, "article.html", gin.H{
// 		"Article": module.Detail(url),

// 		"LastPost": module.LeastPosted(),
// 		"Category": module.Category(),
// 		"Archive": module.Archive(),
// 	})
// }


// func Page(c *gin.Context) {

// 	var url string = c.FullPath()
	
// 	url = strings.TrimPrefix(url, "/")
// 	url = strings.TrimSuffix(url, ".html")

// 	c.HTML(200, "article.html", gin.H{
// 		"Article": module.Page(url),

// 		"LastPost": module.LeastPosted(),
// 		"Category": module.Category(),
// 		"Archive": module.Archive(),
// 	})
// }



// func NotFound(c *gin.Context) {
// 	c.HTML(200, "404.html", gin.H{})
// }



