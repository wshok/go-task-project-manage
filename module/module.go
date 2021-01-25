package module

import (
	"fmt"
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

// type Model struct {
// 	Id        uint `gorm:"column:id"`
// 	CreatedAt time.Time `json:"create_at,omitempty"`
// 	UpdatedAt time.Time `json:"update_at,omitempty"`
// 	DeletedAt time.Time `json:"delete_at,omitempty"`
// }

type User struct {
	Id        uint `json:"id" gorm:"primary_key,AUTO_INCREMENT"`
	Username   string `form:"username" json:"username"`
	Realname   string `form:"realname" json:"realname"`
	Password   string `form:"password" json:"password"`
	Email      string `form:"email" json:"email"`
	Phone      string `form:"phone" json:"phone"`
	Qq         string `form:"qq" json:"qq"`
	Gender     int    `form:"gender" json:"gender"`
	Department string `form:"department" json:"department"`
	Role       string `form:"role" json:"role"`
	Remark     string `form:"remark" json:"remark"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Doc struct {
	Id        int64 `json:"id,omitempty" gorm:"primary_key,AUTO_INCREMENT"`
	Title    string `json:"title,omitempty"`
	Content  string `json:"content,omitempty"`
	Category string `json:"category,omitempty"`
	Uid      int    `json:"uid,omitempty"`
	User     User   `gorm:"ForeignKey:Uid;AssociationForeignKey:id"`
	CreatedAt int64 `json:"create_at,omitempty"`
	UpdatedAt int64 `json:"update_at,omitempty"`
	DeletedAt *time.Time `json:"delete_at,omitempty"`
}

type Task struct {
	Id        uint `json:"id,omitempty" gorm:"primary_key,AUTO_INCREMENT"`
	Title      string    `form:"title" json:"title,omitempty" binding:"required"`
	Content    string    `form:"content" json:"content,omitempty" binding:"required"`
	Uid        string    `form:"uid" json:"uid,omitempty" binding:"required"`
	User       User      `gorm:"ForeignKey:Uid;AssociationForeignKey:id"`
	Status     string    `form:"status" json:"status,omitempty" binding:"required"`
	Progress   string    `form:"progress" json:"progress,omitempty" binding:"required"`
	Project    string    `form:"project" json:"project,omitempty" binding:"required"`
	Type       string    `form:"type" json:"type,omitempty" binding:"required"`
	Accessory  string    `form:"accessory" json:"accessory,omitempty"`
	StartTime  int64 `form:"start_time" json:"start_time,omitempty" binding:"required"`
	EndTime    int64 `form:"end_time" json:"end_time,omitempty" binding:"required"`
	BeginTime  int64 `json:"begin_time,omitempty"`
	FinishTime int64 `json:"finish_time,omitempty"`
	CreatedAt int64 `json:"create_at,omitempty"`
	UpdatedAt int64 `json:"update_at,omitempty"`
	DeletedAt *time.Time `json:"delete_at,omitempty"`
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

	// db.DropTable(&User{}, &Doc{}, &Task{})
	// db.AutoMigrate(&User{}, &Doc{}, &Task{})

	// defer db.Close()

	return db, err
}

//
//@@@@@@@@@@@@@@@@@@@@@@@@@@@user
//

func UserList() []User {
	var val []User
	// todo  page

	db.Debug().Model(&User{}).Order("id desc").Scan(&val)
	fmt.Printf("%#v", val)

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
		var data = Task{Status: "doing", BeginTime: time.Now().Unix()}
		db.Model(&Task{}).Where("id = ?", tid).Updates(data)
	} else if "done" == status {
		var data = Task{Status: "done", FinishTime: time.Now().Unix()}
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

