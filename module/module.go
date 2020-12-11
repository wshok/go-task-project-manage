package module

import (
	// "fmt"
	// "flag"
	"time"
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

type Model struct {
  Id       uint  `json:"id" gorm:"primary_key,AUTO_INCREMENT"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt *time.Time
}

type User struct {
	gorm.Model

	Username   string `json:"username,omitempty"`
	Realname   string `json:"realname,omitempty"`
	Password   string `json:"password,omitempty"`
	Email      string `json:"email,omitempty"`
	Phone      string `json:"phone,omitempty"`
	Qq         string `json:"qq,omitempty"`
	Gender     int    `json:"gender,omitempty"`
	Department string `json:"department,omitempty"`
	Role       string `json:"role,omitempty"`
}

type Doc struct {
	gorm.Model

	Title      string `json:"title,omitempty"`
	Content    string `json:"content,omitempty"`
	Category   string `json:"category,omitempty"`
	Uid        int    `json:"uid,omitempty"`
	User       User   `gorm:"ForeignKey:Uid;AssociationForeignKey:id"`
}

type Task struct {
	gorm.Model

	Title      string `form:"title" json:"title,omitempty" binding:"required"`
	Content    string `form:"content" json:"content,omitempty" binding:"required"`
	Uid        string `form:"uid" json:"uid,omitempty" binding:"required"`
	User       User   `gorm:"ForeignKey:Uid;AssociationForeignKey:id"`
	Status     string `form:"status" json:"status,omitempty" binding:"required"`
	Progress   string `form:"progress" json:"progress,omitempty" binding:"required"`
	Project    string `form:"project" json:"project,omitempty" binding:"required"`
	Type       string `form:"type" json:"type,omitempty" binding:"required"`
	Accessory  string `form:"accessory" json:"accessory,omitempty"`
	StartTime  time.Time  `form:"start_time" json:"start_time,omitempty" binding:"required"`
	EndTime    time.Time  `form:"end_time" json:"end_time,omitempty" binding:"required"`
	BeginTime  time.Time  `json:"begin_time,omitempty"`
	FinishTime time.Time  `json:"finish_time,omitempty"`
}

var (
	db  *gorm.DB
	dsn string = "data/j8rtiEF10ysQY.db"
)

func init() {
	db, _ = opendb()
}

func opendb() (*gorm.DB, error) {

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "hd_" + defaultTableName
	}

	db, err := gorm.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	db.SingularTable(true) // 禁用表名复数

	// db.DropTable(&User{})
	// db.AutoMigrate(&User{})

	// defer db.Close()

	return db, err
}

//
//@@@@@@@@@@@@@@@@@@@@@@@@@@@user
//

func UserList() []User {
	var val []User
	// todo  page

	db.Model(&User{}).Order("id desc").Scan(&val)

	return val
}

func UserAdd(data User) bool {

	db.Create(&data)

	if db.RowsAffected > 0 || db.Error == nil {
		return true
	}

	return false
}

func UserEdit(uid string, data User) int {
	var user User
	db.First(&user, uid)
	if user == (User{}) {
		return -1
	}

	db.Model(&User{}).Where("id = ?", uid).Updates(data)

	if db.RowsAffected > 0 || db.Error == nil {
		return 1
	}

	return 0
}


func UserInfo(tid string) Task {
	var task Task
	db.First(&task, tid)

	return task
}

func UserDelete(uid string) bool {
	var user User
	db.First(&user, uid)

	if user == (User{}) {
		return false
	}

	db.Delete(&user)

	if db.RowsAffected > 0 || db.Error == nil {
		return true
	}

	return false
}


//
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@task
//

func TaskList() []Task {
	var val []Task

	db.Model(&Task{}).Order("id desc").Preload("User").Find(&val)

	return val
}

func TaskAdd(data Task) bool {

	db.Create(&data)

	if db.RowsAffected > 0 || db.Error == nil {
		return true
	}

	return false
}

func TaskEdit(tid string, data Task) int {
	var task Task
	db.First(&task, tid)
	if task == (Task{}) {
		return -1
	}

	db.Model(&Task{}).Where("id = ?", tid).Updates(data)

	if db.RowsAffected > 0 || db.Error == nil {
		return 1
	}

	return 0
}

func TaskInfo(tid string) Task {
	var task Task
	db.First(&task, tid)

	return task
}

func TaskModify(tid, status string) bool {
	var task Task
	db.First(&task, tid)

	if task == (Task{}) {
		return false
	}

	if "doing" == status {
		var data = Task{Status: "doing", BeginTime: time.Now()}
		db.Model(&Task{}).Where("id = ?", tid).Updates(data)
	} else if "done" == status {
		var data = Task{Status: "done", FinishTime: time.Now()}
		db.Model(&Task{}).Where("id = ?", tid).Updates(data)
	}

	if db.RowsAffected > 0 || db.Error == nil {
		return true
	}

	return false
}

func TaskDelete(tid string) bool {
	var task Task
	db.First(&task, tid)

	if task == (Task{}) {
		return false
	}

	db.Delete(&task)

	if db.RowsAffected > 0 || db.Error == nil {
		return true
	}

	return false
}


//
//@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@doc
//

func DocList() []Doc {
	var val []Doc

	db.Model(&Doc{}).Order("id desc").Preload("User").Find(&val)

	return val
}

func DocAdd(data Doc) bool {

	db.Create(&data)

	if db.RowsAffected > 0 || db.Error == nil {
		return true
	}

	return false
}

func DocEdit(id string, data Doc) int {
	var doc Doc
	db.First(&doc, id)
	if doc == (Doc{}) {
		return -1
	}

	db.Model(&Doc{}).Where("id = ?", id).Updates(data)

	if db.RowsAffected > 0 || db.Error == nil {
		return 1
	}

	return 0
}


func DocInfo(id string) Doc {
	var doc Doc
	db.First(&doc, id)

	return doc
}

func DocDelete(id string) bool {
	var doc Doc
	db.First(&doc, id)

	if doc == (Doc{}) {
		return false
	}

	db.Delete(&doc)

	if db.RowsAffected > 0 || db.Error == nil {
		return true
	}

	return false
}







//
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
