package main

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	UserId         int64  `gorm:"primary_key;column:user_id"`
	PassWord       string `gorm:"column:pass_word"`
	NickName       string `gorm:"column:nick_name"`
	ProfilePicture string `gorm:"column:profile_picture"`
}

func (User) TableName() string {
	return "user_info_tab"
}

func main() {
	user, err := findUser(10)
	if err != nil {
		fmt.Printf("错误内容: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("堆栈信息: \n%+v\n", err)
		return
	}
	fmt.Println(user)
}

func findUser(id int64) (User, error) {
	var list []User
	db := createGorm()
	res := db.Where("user_id = ?", id).Find(&list)
	if res.Error != nil {
		return User{}, errors.Wrap(res.Error, fmt.Sprintf("数据库错误参数为ID=%d", id))
	}
	if len(list) <= 0 {
		return User{}, errors.New(fmt.Sprintf("用户id为%d用户不存在", id))
	}
	return list[0], nil
}

// CreateGorm 初始化Gorm
func createGorm() (db *gorm.DB) {
	user := "root"
	password := "Ww19981003"
	address := "127.0.0.1"
	port := "3306"
	dbName := "entry_task_cn_db"

	dsn :=
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, address, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
