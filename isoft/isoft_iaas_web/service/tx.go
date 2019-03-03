package service

import "github.com/astaxie/beego/orm"

func ExecuteServiceWithTx(serviceArgs map[string]interface{}, serviceFunc func(args map[string]interface{}) error) error {
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		return err
	}
	serviceArgs["o"] = o
	if err := serviceFunc(serviceArgs); err != nil {
		o.Rollback()
		return err
	}
	o.Commit()
	return nil
}
