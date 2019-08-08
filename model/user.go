package model

import (
	"github.com/phjt-go/logger"
	"jenkins_demo/setting"
	"time"
)

// UserInfo 获取用户信息
func UserInfo(name string) (user setting.User) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("MGetAccountInfo mysql error, ", r)
		}
	}()

	rows, err := DB.Query("SELECT u.name,u.mobile FROM user u WHERE u.name = ?", name)
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	// 数据处理
	for rows.Next() {
		rows.Scan(&user.Name, &user.Mobile)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	return user
}

// UpdateMobileByName 修改手机号
func UpdateMobileByName(mobile, name string) error {
	return Update("Update `user` set mobile = ? where name = ? ", mobile, name)
}

// AddUser 添加账户
func AddUser(name, mobile string) (id int64, err error) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("MAddAccount mysql error, ", r)
		}
	}()
	return Insert("INSERT INTO user(name,mobile,timestamp) values (?,?,?)",
		name, mobile, time.Now())
}
