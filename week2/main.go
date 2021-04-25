package main

import (
	"context"
	"errors"
)

 var (
	 NoRowsFoundError     = add(1060006, "没有找到任何记录")

 )
//没有记录时需要抛出错误以及相应错误码，与一般SQL错误区分出来；大概逻辑如下

func (d *Dao) GetUserInfoByUserID(ctx context.Context, userID string) (row *entity.UserInfoSimple, err error) {
	row = new(entity.UserInfoSimple)
	err = d.MySQL.Table(ctx, "user_info").
		Where(" podcast_id  =  ? ", userID).
		Find(&row).
		Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.Wrapf(errcode.NoRowsFoundError, "%s", err)
	} else if err != nil {
		return nil, errors.Wrapf(errcode.MysqlError, "%s", err)
	}

	return row, nil
}