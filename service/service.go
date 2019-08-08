package service

import (
	"fmt"
	"github.com/Jenkins/model"
	"github.com/Jenkins/setting"
)

func AddUser(name, mobile string) error {
	id, err := model.AddUser(name, mobile)
	if err != nil {
		return err
	}
	fmt.Printf("添加用户%v成功，ID为%v", name, id)
	return nil
}

func UserInfo(name string) setting.User {
	return model.UserInfo(name)
}

func UpdateUser(name, mobile string) error {
	if err := model.UpdateMobileByName(mobile, name); err != nil {
		return err
	}
	return nil
}
