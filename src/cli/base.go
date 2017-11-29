package cli

import (
	"config"
	"github.com/astaxie/beego/orm"
	"github.com/qiniu/log"
)

func InitOrm(cfg *config.OrmConfig) (err error) {
	log.Infof("Init orm , %s\n", cfg.DataSource)
	regErr := orm.RegisterDataBase("default", cfg.DriverName, cfg.DataSource, cfg.MaxIdleConn, cfg.MaxOpenConn)
	if regErr != nil {
		err = regErr
		log.Errorf("connect database error , %s\n", err.Error())
		return
	}

	orm.Debug = cfg.DebugMode

	orm.RegisterModel(new(Users))
	orm.RegisterModel(new(Sallers))

	name := "default"
	force := false
	verbose := true
	cErr := orm.RunSyncdb(name, force, verbose)
	if cErr != nil {
		err = cErr
		log.Errorf("create table error , %s\n", err.Error())
		return
	}
	log.Info("Init orm success")
	return
}
