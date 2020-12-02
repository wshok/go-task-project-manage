package module

import (

	// "fmt"
	// "flag"
	// "time"
	// "strconv"
	// "database/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	// "github.com/cihub/seelog"
)

type Page struct {
	Offset int
	Size   int
}

type Doc struct {
	id           int
	Title        string
	Content      string
	Category     string
	Author       int
	Create_time  int
	Update_time  int
	Delete_time  int
}

type Task struct {
	Id           int
	Title        string
	Content      string
	Owner        int
	Status       int
	Progress     int
	Project      int
	Type         string
	Accessory    string
	Start_time   int
	End_time     int
	Finish_time  int
	Create_time  int
	Update_time  int
	Delete_time  int
}


type User struct {
	Id           int
	Username     string
	Realname     string
	Password     string
	Email        string
	Phone        string
	Qq           string
	Gender       int
	Department   string
	Role         int
	Create_time  int
	Update_time  int
	Delete_time  int
}


var (
	db   *gorm.DB
	dsn string = "data/j8rtiEF10ysQY.db"
)


func init() {
	db, _ = opendb()
}


func opendb() (*gorm.DB, error) {

	// logConfigPath := flag.String("L", "conf/seelog.xml", "log config file path")

	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
	    return "hd_" + defaultTableName;
	}

	db, err := gorm.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	db.SingularTable(true)	// 禁用表名复数

  	// defer db.Close()

	return db, err
}


func UserList() []User {
	var val []User

	// var p Page {
	// 	Start: (page - 1) * p.ListRows,
	// 	Size:  p.ListRows,
	// }

	db.Model(&User{}).Where("delete_time = 0").Order("id desc").Scan(&val)

	return val
}

func TaskList() []Task {
	var val []Task

	db.Model(&Task{}).Where("delete_time = 0").Order("id desc").Scan(&val)

	return val
}

// func ArticleCount(f *Filter) int {
// 	var (
// 		val      int
// 	)

// 	if (f.Page == "index") {
// 		DB.Model(&Content{}).Where("status=? AND type=?", "publish", "post").Count(&val)

// 	} else if (f.Page == "category") {

// 		DB.Model(&Content{}).Where("status=? AND type=? ", "publish", "post").Count(&val)

// 		DB.Model(&Content{}).Joins("left join typecho_relationships on typecho_relationships.cid = typecho_contents.cid") .Joins("left join typecho_metas on typecho_metas.mid = typecho_relationships.mid") .Where("typecho_contents.status=? AND typecho_contents.type=? AND typecho_metas.type = ?", "publish", "post", "category") .Where("typecho_metas.slug = ?", f.Category) .Count(&val)
// 	} else if (f.Page == "archive") {

// 		DB.Model(&Content{}).Where("status=? AND type=?", "publish", "post") .Where("strftime('%Y/%m',datetime(created, 'unixepoch')) = ?", fmt.Sprintf("%s/%s", f.Year, f.Month)).Count(&val)
// 	}

// 	return val
// }


// func ArticleList(l *LimitRows, f *Filter) []Content {
// 	var (
// 		val      []Content
// 	)

// 	if (f.Page == "index") {
// 		DB.Model(&Content{}) .Select("typecho_contents.*, typecho_metas.name AS category_name, typecho_metas.slug AS category_slug, typecho_users.screenName as author") .Joins("left join typecho_relationships on typecho_relationships.cid = typecho_contents.cid") .Joins("left join typecho_metas on typecho_metas.mid = typecho_relationships.mid") .Joins("left join typecho_users on typecho_contents.authorId = typecho_users.uid") .Where("typecho_contents.status=? AND typecho_contents.type=? AND typecho_metas.type = ?", "publish", "post", "category") .Order("typecho_contents.cid desc").Offset(l.Start) .Limit(l.Size).Scan(&val)
// 	} else if (f.Page == "category") {

// 		DB.Model(&Content{}) .Select("typecho_contents.*, typecho_metas.name AS category_name, typecho_metas.slug AS category_slug, typecho_users.screenName as author") .Joins("left join typecho_relationships on typecho_relationships.cid = typecho_contents.cid") .Joins("left join typecho_metas on typecho_metas.mid = typecho_relationships.mid") .Joins("left join typecho_users on typecho_contents.authorId = typecho_users.uid") .Where("typecho_contents.status=? AND typecho_contents.type=? AND typecho_metas.type = ?", "publish", "post", "category") .Where("typecho_metas.slug = ?", f.Category) .Order("typecho_contents.cid desc").Offset(l.Start) .Limit(l.Size).Scan(&val)
// 	} else if (f.Page == "archive") {

// 		DB.Model(&Content{}) .Select("typecho_contents.*, typecho_metas.name AS category_name, typecho_metas.slug AS category_slug, typecho_users.screenName as author") .Joins("left join typecho_relationships on typecho_relationships.cid = typecho_contents.cid") .Joins("left join typecho_metas on typecho_metas.mid = typecho_relationships.mid") .Joins("left join typecho_users on typecho_contents.authorId = typecho_users.uid") .Where("typecho_contents.status=? AND typecho_contents.type=? AND typecho_metas.type = ?", "publish", "post", "category") .Where("strftime('%Y/%m',datetime(typecho_contents.created, 'unixepoch')) = ?", fmt.Sprintf("%s/%s", f.Year, f.Month)) .Order("typecho_contents.cid desc").Offset(l.Start).Limit(l.Size).Scan(&val)
// 	}

// 	return val
// }


// func LeastPosted() []Content {
// 	var (
// 		val      []Content
// 	)

// 	DB.Model(&Content{}).Select("cid,title,slug").Where("status=? AND type=?", "publish", "post").Order("created desc").Offset(0).Limit(8).Scan(&val)

// 	return val
// }


// func Category() []Meta {
// 	var (
// 		val      []Meta
// 	)

// 	DB.Table("typecho_metas").Where("type=?", "category").Order("order").Scan(&val)

// 	return val
// }

// type Result struct {
//     Yearmonth string
//     Count  int
// }

// func Archive() []Result{

// 	var val []Result

// 	DB.Table("typecho_contents").Select("strftime('%Y/%m',datetime(created, 'unixepoch')) AS yearmonth, COUNT(1) AS count") .Where("status=? AND type=?", "publish", "post") .Group("yearmonth") .Order("created desc") .Scan(&val)
		
// 	return val
// }

// func Detail(url string) Content {
// 	var val Content

//     DB.Model(&Content{}).Select("typecho_contents.*, typecho_metas.name AS category_name, typecho_metas.slug AS category_slug, typecho_users.screenName as author") .Joins("left join typecho_relationships on typecho_relationships.cid = typecho_contents.cid") .Joins("left join typecho_metas on typecho_metas.mid = typecho_relationships.mid") .Joins("left join typecho_users on typecho_contents.authorId = typecho_users.uid") .Where("typecho_contents.status=? AND typecho_contents.type=? AND typecho_metas.type = ? and typecho_contents.slug = ?", "publish", "post", "category", url) .Scan(&val)

// 	return val
// }

// func Page(url string) Content {
// 	var val Content

//     DB.Model(&Content{}).Where("status=? AND type=? and slug = ?", "publish", "page", url) .Scan(&val)

// 	return val
// }

