package cli

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	pili2 "github.com/pili-engineering/pili-sdk-go.v2/pili"
	"github.com/qiniu/log"
)

type Users struct {
	Name     string `json:"name" orm:"pk"`
	Password string `json:"password"`
	Room     string `json:"room"`
	Deadline int64  `json:"deadline"`
}

func InsertUser(user *Users) error {
	log.Info("insert user")
	info, err := orm.NewOrm().Insert(user)
	fmt.Println(info)
	if err != nil {
		return err
	}
	return nil
}

func UserIsExisted(name string) (Users, error) {
	var user Users
	err := orm.NewOrm().QueryTable("users").Filter("name", name).One(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func QueryUser(name, password string) error {
	user := &Users{}
	err := orm.NewOrm().QueryTable("users").Filter("name", name).Filter("password", password).One(user)
	if err != nil {
		return err
	}
	fmt.Printf("name=%s,password=%s,room=%s,deadline=%d\n", user.Name, user.Password, user.Room, user.Deadline)
	return nil
}

func UpdateUser(name, pwd string) error {
	//user
	user := new(Users)
	//query user
	err := orm.NewOrm().QueryTable("users").Filter("name", name).One(user)
	if err != nil {
		return errors.New("user not found")
	}
	if pwd != "" {
		user.Password = pwd
	}
	_, uErr := orm.NewOrm().Update(user, "password")
	if uErr != nil {
		return uErr
	}
	return nil
}

func DeleteUser(mac *pili2.MAC, name string) (int64, error) {
	user := new(Users)
	//查询用户是否存在
	err := orm.NewOrm().QueryTable("users").Filter("name", name).One(user)
	if err != nil {
		return 0, err
	}
	cnt, dErr := orm.NewOrm().Delete(user)
	if dErr != nil {
		return 0, dErr
	}
	fmt.Printf("delete user , room = %s\n", user.Room)
	_, rErr := RoomDelete(mac, user.Room)
	if rErr != nil {
		return 0, rErr
	}
	return cnt, nil
}

func DeleteUserByTimer(mac *pili2.MAC, ttl int64) {
	var users []Users
	_, err := orm.NewOrm().QueryTable("users").Filter("deadline__lte", ttl).All(&users)
	if err != nil {
		fmt.Println("delete temp user failur ....")
		return
	}
	fmt.Println(len(users))
	for i := 0; i < len(users); i++ {
		DeleteUser(mac, users[i].Name)
	}
	return
}
