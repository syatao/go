package main

import "errors"

 var (
	 NoRowsFoundError     = add(1060006, "没有找到任何记录")

 )
//没有记录时需要抛出错误以及相应错误码，与一般SQL错误区分出来；大概逻辑如下
func SelectUserByID()  {
	currentSql := d.MySQL.Table(ctx, copyright.TableName()).Where(query)
	err = currentSql.Find(&res).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.Wrapf(errcode.NoRowsFoundError, "%s", err)
	} else if err != nil {
		return nil, errors.Wrapf(errcode.MysqlError, "%s", err)
	}
}
