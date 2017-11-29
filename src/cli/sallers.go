package cli

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/qiniu/log"
)

type Sallers struct {
	Name     string `json:"name" orm:"pk"`
	Password string `json:"password"`
}

func InsertSaller(saller *Sallers) error {
	log.Info("insert saller")
	info, err := orm.NewOrm().Insert(saller)
	fmt.Println(info)
	if err != nil {
		return err
	}
	return nil
}

func QuerySaller(name, password string) error {
	saller := &Sallers{}
	err := orm.NewOrm().QueryTable("sallers").Filter("name", name).Filter("password", password).One(saller)
	if err != nil {
		return err
	}
	fmt.Printf("name=%s,password=%s,room=%s,deadline=%d\n", saller.Name, saller.Password)
	return nil
}
